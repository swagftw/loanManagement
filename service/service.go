package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
)

type LoanService interface {
	CreateLoan(ctx context.Context, customerName string, phoneNo string, email string, loanAmount int64, creditScore int64) (string, error)
	GetLoans(ctx context.Context, status []string, loanAmountGreaterThan int64, page int64, limit int64) ([]Loan, error)
	GetLoan(ctx context.Context, id string) (Loan, error)
	PatchLoan(ctx context.Context, id string, status string) (string, error)
	CancelLoan(ctx context.Context, id string) (string, error)
}

type loanService struct {
	repo   MongoRepository
	logger log.Logger
}

func NewLoanService(repo MongoRepository, logger log.Logger) LoanService {
	return loanService{
		repo:   repo,
		logger: logger,
	}
}

func (l loanService) CreateLoan(ctx context.Context, customerName string, phoneNo string, email string, loanAmount int64, creditScore int64) (string, error) {
	l.logger.Log("method", "create loan")
	uuid := uuid.New()
	loan := Loan{
		Id:           uuid.String(),
		CustomerName: customerName,
		PhoneNo:      phoneNo,
		Email:        email,
		LoanAmount:   loanAmount,
		Status:       "new",
		CreditScore:  creditScore,
	}
	err := l.repo.CreateLoan(ctx, loan)
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

func (l loanService) GetLoans(ctx context.Context, status []string, loanAmountGreaterThan int64, page int64, limit int64) ([]Loan, error) {
	l.logger.Log("method", "Get loans")
	loans, _ := l.repo.GetLoans(ctx, status, loanAmountGreaterThan, page, limit)
	return loans, nil
}

func (l loanService) GetLoan(ctx context.Context, id string) (Loan, error) {
	l.logger.Log("method", "get loan", "id", id)
	loan, _ := l.repo.GetLoan(ctx, id)
	return loan, nil
}

func (l loanService) PatchLoan(ctx context.Context, id string, status string) (string, error) {
	l.logger.Log("method", "patch loan")
	updatedId, _ := l.repo.PatchLoan(ctx, id, status)
	return updatedId, nil
}

func (l loanService) CancelLoan(ctx context.Context, id string) (string, error) {
	l.logger.Log("method", "cancel loan")
	fmt.Println(id)
	cancelledId, _ := l.repo.CancelLoan(ctx, id)
	return cancelledId, nil
}
