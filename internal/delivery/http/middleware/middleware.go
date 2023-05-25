package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"tcgasstation-backend/internal/delivery/http/request"
	"tcgasstation-backend/internal/delivery/http/response"
	"tcgasstation-backend/internal/usecase"
	"tcgasstation-backend/utils"
	"tcgasstation-backend/utils/global"
	"tcgasstation-backend/utils/helpers"
	"tcgasstation-backend/utils/logger"
	"tcgasstation-backend/utils/redis"

	"go.uber.org/zap"
)

type IMiddleware interface {
	LoggingMiddleware(next http.Handler) http.Handler
	AuthorizationFunc(next http.Handler) http.Handler
	Pagination(next http.Handler) http.Handler
	ValidateAccessToken(next http.Handler) http.Handler
}

type middleware struct {
	usecase          usecase.Usecase
	response         response.IHttpResponse
	cache            redis.IRedisCache
	cacheAuthService redis.IRedisCache
}

func NewMiddleware(uc usecase.Usecase, g *global.Global) *middleware {
	m := new(middleware)
	m.usecase = uc
	m.response = response.NewHttpResponse()
	m.cache = g.Cache
	m.cacheAuthService = g.CacheAuthService
	return m
}

func (m *middleware) LoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				logger.AtLog.Logger.Error("err", zap.Any("err", err), zap.Any("trace", debug.Stack()))

			}
		}()

		start := time.Now()
		wrapped := wrapResponseWriter(w)
		next.ServeHTTP(wrapped, r)
		logger.AtLog.Info(fmt.Sprintf("Request:[%s] %s - status: %d - duration %s", r.Method, r.URL.EscapedPath(), wrapped.status, time.Since(start)))
	}

	return http.HandlerFunc(fn)
}

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

func (m *middleware) Pagination(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		pageInt := 1
		limitInt := 10

		page := r.URL.Query().Get("page")
		limit := r.URL.Query().Get("limit")
		sortBy := r.URL.Query().Get("sort_by")
		sortStr := r.URL.Query().Get("sort")

		if page != "" {
			tmp, err := strconv.Atoi(page)
			if err == nil {
				pageInt = tmp
			}
		}

		if limit != "" {
			tmp, err := strconv.Atoi(limit)
			if err == nil {
				limitInt = tmp
			}

			if limitInt > 100 {
				limitInt = 100
			}
		}

		offset := limitInt * (pageInt - 1)

		pag := request.PaginationReq{
			Page:   &pageInt,
			Limit:  &limitInt,
			Offset: &offset,
		}

		if sortStr != "" {
			sortInt, err := strconv.Atoi(sortStr)
			if err == nil {
				pag.Sort = &sortInt
			}
		}

		if sortBy != "" {
			pag.SortBy = &sortBy
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, utils.PAGINATION, pag)
		wrapped := wrapResponseWriter(w)
		next.ServeHTTP(wrapped, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// Authenticate
func (m *middleware) ValidateAccessToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		wrapped, ctx, err := m.ValidateToken(w, r)
		if err != nil {
			logger.AtLog.Logger.Error("accessToken", zap.Error(err))
			m.response.RespondWithError(w, http.StatusUnauthorized, response.Error, err)
			return
		}
		next.ServeHTTP(wrapped, r.WithContext(*ctx))
	}
	return http.HandlerFunc(fn)
}

// Just set SIGNED_WALLET_ADDRESS, SIGNED_USER_ID to context if have
// Authorization
func (m *middleware) AuthorizationFunc(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		wrapped, ctx, err := m.ValidateToken(w, r)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(wrapped, r.WithContext(*ctx))
	}
	return http.HandlerFunc(fn)
}

func (m *middleware) ValidateToken(w http.ResponseWriter, r *http.Request) (*responseWriter, *context.Context, error) {
	ctx := r.Context()
	token := helpers.ReplaceToken(r.Header.Get(utils.AUTH_TOKEN))
	if token == "" {
		err := errors.New("Token is empty")
		return nil, nil, err
	}
	p, err := m.usecase.ValidateAccessToken(token)
	if err != nil {
		return nil, nil, err
	}

	ctx = context.WithValue(ctx, utils.AUTH_TOKEN, token)
	ctx = context.WithValue(ctx, utils.SIGNED_WALLET_ADDRESS, p.WalletAddress)
	ctx = context.WithValue(ctx, utils.SIGNED_USER_ID, p.Uid)
	wrapped := wrapResponseWriter(w)
	return wrapped, &ctx, nil
}
