package service

import (
	ht "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPServer(endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	r.Path("/loans").Handler(ht.NewServer(
		endpoints.CreateLoanEndpoint,
		DecodeCreateLoanRequest,
		EncodeCreateLoanResponse,
	)).Methods("POST")

	r.Path("/loans").Handler(ht.NewServer(
		endpoints.GetLoansEndpoint,
		DecodeGetLoansRequest,
		EncodeGetLoansResponse,
	)).Methods("GET")

	r.Path("/loan/{id}").Handler(ht.NewServer(
		endpoints.GetLoanEndpoint, DecodeGetLoanRequest, EncodeGetLoanResponse))

	r.Path("/loans/{id}").Handler(ht.NewServer(
		endpoints.PatchLoanEndpoint, DecodePatchLoanRequest, EncodePatchLoanResponse,
	)).Methods("PATCH")

	r.Path("/loans/{id}").Handler(ht.NewServer(
		endpoints.CancelLoanEndpoint,
		DecodeCancelLoanRequest,
		EncodeCancelLoanResponse,
	)).Methods("DELETE")
	return r
}
