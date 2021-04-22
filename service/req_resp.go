package service

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

type (
	CreateLoanRequest struct {
		CustomerName string `json:"customerName"`
		PhoneNo      string `json:"phoneNo"`
		Email        string `json:"email"`
		LoanAmount   int64  `json:"loanAmount"`
		CreditScore  int64  `json:"creditScore"`
	}

	CreateLoanResponse struct {
		Id      string `json:"id"`
		Message string `json:"message"`
	}

	GetLoansRequest struct {
		Status                []string `json:"status"`
		LoanAmountGreaterThan int64    `json:"loanAmountGreaterThan"`
		Page                  int64    `json:"page"`
		Limit                 int64    `json:"limit"`
	}

	GetLoansResponse struct {
		Loans []Loan `json:"loans"`
	}

	GetLoanRequest struct {
		Id string `json:"id"`
	}

	GetLoanResponse struct {
		Loan Loan `json:"loan"`
	}

	PatchLoanRequest struct {
		Status string `json:"status"`
		Id     string `json:"id"`
	}

	PatchLoanResponse struct {
		UpdatedId string `json:"updatedId"`
	}

	CancelLoanRequest struct {
		Id string `json:"id"`
	}

	CancelLoanResponse struct {
		CancelledId string `json:"cancelledId"`
	}
)

func EncodeCreateLoanResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var res CreateLoanResponse
	res = response.(CreateLoanResponse)
	if len(res.Id) != 0 {
		w.WriteHeader(http.StatusCreated)
	}
	return json.NewEncoder(w).Encode(response)
}

func DecodeCreateLoanRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	var req CreateLoanRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func EncodeGetLoansResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var res GetLoansResponse
	res = response.(GetLoansResponse)
	loans := res.Loans
	if len(loans) == 0 {
		w.WriteHeader(http.StatusNotFound)
		res.Loans = []Loan{}
		err := json.NewEncoder(w).Encode(res)
		return err
	}
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetLoansRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetLoansRequest
	query := r.URL.Query()

	req.Status = query["status"]
	if query["loanAmountGreater"] != nil {
		lgt, err := strconv.ParseInt(query["loanAmountGreater"][0], 10, 64)
		if err != nil {
			return nil, err
		}
		req.LoanAmountGreaterThan = lgt
	}

	if query["page"] != nil {
		page, err := strconv.ParseInt(query["page"][0], 10, 64)
		if err != nil {
			return nil, err
		}
		req.Page = page
	}

	if query["limit"] != nil {
		limit, err := strconv.ParseInt(query["limit"][0], 10, 64)
		if err != nil {
			return nil, err
		}
		req.Limit = limit
	}

	var status []string
	if len(req.Status) != 0 {
		for i := range req.Status {
			if strings.Contains(req.Status[i], ",") {
				list := strings.Split(req.Status[i], ",")
				for i2 := range list {
					status = append(status, list[i2])
				}
			} else {
				status = append(status, req.Status[i])
			}
		}
		req.Status = status
	}
	return req, nil
}

func DecodeGetLoanRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	id := mux.Vars(request)["id"]
	return GetLoanRequest{
		Id: id,
	}, nil
}

func EncodeGetLoanResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var res GetLoanResponse
	res = response.(GetLoanResponse)
	if len(res.Loan.Id) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func EncodePatchLoanResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var res PatchLoanResponse
	res = response.(PatchLoanResponse)
	if len(res.UpdatedId) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func DecodePatchLoanRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	id := mux.Vars(request)["id"]
	var req PatchLoanRequest
	_ = json.NewDecoder(request.Body).Decode(&req)
	req.Id = id
	return PatchLoanRequest{
		Status: req.Status,
		Id:     req.Id,
	}, nil
}

func EncodeCancelLoanResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var res CancelLoanResponse
	res = response.(CancelLoanResponse)
	if len(res.CancelledId) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func DecodeCancelLoanRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	id := mux.Vars(request)["id"]
	return CancelLoanRequest{
		Id: id,
	}, nil
}
