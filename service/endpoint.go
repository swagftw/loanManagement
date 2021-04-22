package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateLoanEndpoint endpoint.Endpoint
	GetLoansEndpoint   endpoint.Endpoint
	GetLoanEndpoint    endpoint.Endpoint
	PatchLoanEndpoint  endpoint.Endpoint
	CancelLoanEndpoint endpoint.Endpoint
}

func NewLoanEndpoints(s LoanService) Endpoints {
	return Endpoints{CreateLoanEndpoint: makeCreateLoanEndpoint(s), GetLoansEndpoint: makeGetLoansEndpoint(s), GetLoanEndpoint: makeGetLoanEndpoint(s), PatchLoanEndpoint: makePatchLoanEndpoint(s), CancelLoanEndpoint: makeCancelLoanEndpoint(s)}
}

func makeCreateLoanEndpoint(l LoanService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateLoanRequest)
		id, err := l.CreateLoan(ctx, req.CustomerName, req.PhoneNo, req.Email, req.LoanAmount, req.CreditScore)
		if err != nil {
			return CreateLoanResponse{
				Id:      "",
				Message: fmt.Sprintf("%v", err),
			}, err
		}
		return CreateLoanResponse{
			Id:      id,
			Message: "Loan created successfully",
		}, nil
	}
}

func makeGetLoansEndpoint(s LoanService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetLoansRequest)
		loans, _ := s.GetLoans(ctx, req.Status, req.LoanAmountGreaterThan, req.Page,req.Limit)
		return GetLoansResponse{Loans: loans}, nil
	}
}

func makeGetLoanEndpoint(s LoanService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetLoanRequest)
		loan, _ := s.GetLoan(ctx, req.Id)
		return GetLoanResponse{Loan: loan}, nil
	}
}

func makePatchLoanEndpoint(s LoanService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PatchLoanRequest)
		id, _ := s.PatchLoan(ctx, req.Id, req.Status)
		return PatchLoanResponse{UpdatedId: id}, nil
	}
}

func makeCancelLoanEndpoint(s LoanService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var req CancelLoanRequest
		req = request.(CancelLoanRequest)
		id, _ := s.CancelLoan(ctx, req.Id)
		return CancelLoanResponse{
			CancelledId: id,
		}, nil
	}
}
