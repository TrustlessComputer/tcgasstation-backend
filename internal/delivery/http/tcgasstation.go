package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"tcgasstation-backend/internal/delivery/http/request"
	"tcgasstation-backend/internal/delivery/http/response"
	"tcgasstation-backend/utils/logger"
	"time"

	"go.uber.org/zap"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
)

// generateDepositAddress godoc
//
//	@Summary		Generate deposit address
//	@Description	Generate deposit address
//	@Tags			Bridge
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		request.GenerateDepositAddressReq	true	"tc address info"
//	@Success		200		{object}	response.GenerateDepositAddressResp{}
//	@Router			/api/generate-deposit-address [POST]
func (h *httpDelivery) generateDepositAddress(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			reqBody := &request.GenerateDepositAddressReq{}

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(reqBody)
			if err != nil {
				return nil, err
			}

			if reqBody.RecaptchaResponse == "" {

				err = errors.New("the recaptcha is required.")
				return nil, err
			}

			if len(h.Usecase.Config.CaptcharSecret) > 0 {
				captcha, _ := recaptcha.NewReCAPTCHA(h.Usecase.Config.CaptcharSecret, recaptcha.V3, 10*time.Second) // for v2 API get your secret from https://www.google.com/recaptcha/admin

				err = captcha.Verify(reqBody.RecaptchaResponse)
				if err != nil {
					logger.AtLog.Logger.Error("h.generateDepositAddress.recaptcha.Verify", zap.String("err", err.Error()))
					return nil, err
				}
			}

			resp, err := h.Usecase.GenerateDepositAddress(reqBody)
			if err != nil {
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}

func (h *httpDelivery) hello(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {

			fee, _ := h.Usecase.EstFeeDepositBtc(0)

			// h.Usecase.JobBridge_ProcessWithdrawEthTxs()
			// h.Usecase.RunPullAllEthTxs(3431530, 3431532)

			return fee, nil
		},
	).ServeHTTP(w, r)
}

// listToken godoc
// @Summary list bridge tokens
// @Description list bridge tokens
// @Tags Bridge
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.TcToken
// @Router /api/tokens [GET]
func (h *httpDelivery) listToken(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			// resp, err := h.Usecase.ListToken()
			// if err != nil {
			// 	return nil, err
			// }
			// return resp, nil
			return nil, nil
		},
	).ServeHTTP(w, r)
}

// estimateWithdrawFee godoc
// @Summary list list user deposit withdraw
// @Description list user deposit withdraw
// @Tags Bridge
// @Accept  json
// @Produce  json
// @Param none
// @Success 200 {array} entity.DepositWithdraw
// @Router /api/deposit-withdraw-history [GET]
func (h *httpDelivery) historyTcGasStation(w http.ResponseWriter, r *http.Request) {
	response.NewRESTHandlerTemplate(
		func(ctx context.Context, r *http.Request, vars map[string]string) (interface{}, error) {
			address := r.URL.Query().Get("address")
			resp, err := h.Usecase.HistoryTcGasStation(address)
			if err != nil {
				return nil, err
			}

			return resp, nil
		},
	).ServeHTTP(w, r)
}
