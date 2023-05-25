package http

import (
	"os"

	"tcgasstation-backend/docs"
	_ "tcgasstation-backend/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (h *httpDelivery) registerRoutes() {
	h.RegisterDocumentRoutes()
	h.RegisterV1Routes()
}

func (h *httpDelivery) RegisterV1Routes() {
	h.Handler.Use(h.MiddleWare.LoggingMiddleware)
	h.Handler.Use(h.MiddleWare.Pagination)

	//api
	api := h.Handler.PathPrefix("/api").Subrouter()
	api.HandleFunc("/generate-address", h.generateDepositAddress).Methods("POST")

	api.HandleFunc("/tokens", h.listToken).Methods("GET")
	api.HandleFunc("/history", h.historyTcGasStation).Methods("GET")

	api.HandleFunc("/hello", h.hello).Methods("GET")

	//AUTH
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/nonce", h.generateMessage).Methods("POST")
	auth.HandleFunc("/nonce/verify", h.verifyMessage).Methods("POST")

	//profile
	profile := api.PathPrefix("/profile").Subrouter()
	profile.HandleFunc("/wallet/{walletAddress}", h.profileByWallet).Methods("GET")
	profile.HandleFunc("/wallet/{walletAddress}/histories", h.currentUerProfileHistories).Methods("GET")

	profileAuth := api.PathPrefix("/profile").Subrouter()
	profileAuth.Use(h.MiddleWare.AuthorizationFunc)
	profileAuth.HandleFunc("/me", h.currentUerProfile).Methods("GET")
	profileAuth.HandleFunc("/histories", h.createProfileHistory).Methods("POST")
	profileAuth.HandleFunc("/histories", h.confirmProfileHistory).Methods("PUT")

	//admin
	admin := api.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/update-enabled-job", h.updateCronJobStatus).Methods("POST")
	admin.HandleFunc("/update-enabled-job-name", h.updateCronJobStatusByFuncName).Methods("POST")
	admin.HandleFunc("/get-job-info", h.getCronJobInfo).Methods("POST")

	admin.HandleFunc("/redis", h.getRedisKeys).Methods("GET")
	admin.HandleFunc("/redis/{key}", h.getRedis).Methods("GET")
	admin.HandleFunc("/redis", h.upsertRedis).Methods("POST")
	admin.HandleFunc("/redis", h.deleteAllRedis).Methods("DELETE")
	admin.HandleFunc("/redis/{key}", h.deleteRedis).Methods("DELETE")

}

func (h *httpDelivery) RegisterDocumentRoutes() {
	documentUrl := `/swagger/`
	domain := os.Getenv("swagger_domain")
	docs.SwaggerInfo.Host = domain
	docs.SwaggerInfo.BasePath = "/"
	swaggerURL := documentUrl + "swagger/doc.json"
	h.Handler.PathPrefix(documentUrl).Handler(httpSwagger.Handler(
		httpSwagger.URL(swaggerURL), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		//httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))
}
