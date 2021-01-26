package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers requestmoney-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/requestmoney/request", createRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/requestmoney/request", listRequestHandler(cliCtx, "requestmoney")).Methods("GET")
	r.HandleFunc("/requestmoney/request/{key}", getRequestHandler(cliCtx, "requestmoney")).Methods("GET")
	//r.HandleFunc("/requestmoney/request", setRequestHandler(cliCtx)).Methods("PUT")
	//r.HandleFunc("/requestmoney/request", deleteRequestHandler(cliCtx)).Methods("DELETE")

}
