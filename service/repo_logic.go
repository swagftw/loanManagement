package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (m mongoRepository) CreateLoan(ctx context.Context, loan Loan) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(m.timeout)*time.Second)
	defer cancel()
	collection := m.client.Database(m.database).Collection("loans")
	_, err := collection.InsertOne(ctx, loan)
	if err != nil {
		return err
	}
	SendMail(loan.Email, loan.CustomerName, loan.Status)
	return nil
}

func (m mongoRepository) GetLoans(ctx context.Context, status []string, loanAmountGreaterThan int64, page int64, limit int64) ([]Loan, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(m.timeout)*time.Second)
	defer cancel()
	var statusQuery []bson.M
	if len(status) != 0 {
		for i := range status {
			s := status[i]
			statusQuery = append(statusQuery, bson.M{"status": s})
		}
	} else {
		statusQuery = []bson.M{}
	}

	if page == 0 {
		page = 1
	}

	var loans []Loan
	var err error
	var result *mongo.Cursor

	collection := m.client.Database(m.database).Collection("loans")
	if len(statusQuery) != 0 {
		if limit == 0 {
			result, err = collection.Find(ctx, bson.M{"$and": []bson.M{{"loanAmount": bson.M{"$gt": loanAmountGreaterThan}}, {"$or": statusQuery}}}, options.Find().SetSkip((page-1)*10))
		} else {
			result, err = collection.Find(ctx, bson.M{"$and": []bson.M{{"loanAmount": bson.M{"$gt": loanAmountGreaterThan}}, {"$or": statusQuery}}}, options.Find().SetLimit(limit).SetSkip((page-1)*10))
		}

	} else {
		if limit == 0 {
			result, err = collection.Find(ctx, bson.M{"loanAmount": bson.M{"$gt": loanAmountGreaterThan}}, options.Find().SetSkip((page-1)*10))
		} else {
			result, err = collection.Find(ctx, bson.M{"loanAmount": bson.M{"$gt": loanAmountGreaterThan}}, options.Find().SetLimit(limit).SetSkip((page-1)*10))
		}

	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = result.All(ctx, &loans)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return loans, nil
}

func (m mongoRepository) GetLoan(ctx context.Context, id string) (Loan, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(m.timeout)*time.Second)
	defer cancel()
	collection := m.client.Database(m.database).Collection("loans")
	result := collection.FindOne(ctx, bson.M{"id": id})
	var loan Loan
	result.Decode(&loan)
	return loan, nil
}

func (m mongoRepository) PatchLoan(ctx context.Context, id string, status string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(m.timeout)*time.Second)
	defer cancel()
	var loan Loan
	collection := m.client.Database(m.database).Collection("loans")
	err := collection.FindOneAndUpdate(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"status": status}}).Decode(&loan)
	SendMail(loan.Email, loan.CustomerName, status)
	if err != nil {
		return "", err
	}
	return id, err
}

func (m mongoRepository) CancelLoan(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(m.timeout)*time.Second)
	defer cancel()
	var loan Loan
	collection := m.client.Database(m.database).Collection("loans")
	err := collection.FindOneAndUpdate(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"status": "cancelled"}}).Decode(&loan)
	if err != nil {
		return "", err
	}
	SendMail(loan.Email, loan.CustomerName, "cancelled")
	return id, nil
}
