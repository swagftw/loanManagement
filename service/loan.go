package service

import (
	"context"
)

type Loan struct {
	Id           string `json:"id" bson:"id"`
	CustomerName string `json:"customerName" bson:"customerName"`
	PhoneNo      string `json:"phoneNo" bson:"phoneNo"`
	Email        string `json:"email" bson:"email"`
	LoanAmount   int64  `json:"loanAmount" bson:"loanAmount"`
	Status       string `json:"status" bson:"status"`
	CreditScore  int64  `json:"creditScore" bson:"creditScore"`
}

type MongoRepository interface {
	CreateLoan(ctx context.Context, loan Loan) error
	GetLoans(ctx context.Context, status []string, loanAmountGreaterThan int64,page int64,limit int64) ([]Loan, error)
	GetLoan(ctx context.Context, id string) (Loan, error)
	PatchLoan(ctx context.Context, id string, status string) (string, error)
	CancelLoan(ctx context.Context,id string) (string,error)
}

