package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	"github.com/swagftw/loan_management_service/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var logger log.Logger

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger, "service", "loan service", "caller", log.DefaultCaller, "time", log.DefaultTimestampUTC)

	err := godotenv.Load(".env")
	if err != nil {
		level.Error(logger).Log("error", "failed to load env file")
		return
	}
	dbPass := getEnvironmentVariable("DB_PASSWORD")
	var dbString = fmt.Sprintf("mongodb+srv://swag:%s@cluster0.ddq1l.mongodb.net/loanDB?retryWrites=true&w=majority",dbPass)
	mongoRepo, err := service.NewMongoRepository(dbString, "loanDB", 10)
	if err != nil {
		panic(err)
	}

	loanService := service.NewLoanService(mongoRepo, logger)
	level.Info(logger).Log("msg", "started service")
	defer level.Info(logger).Log("msg", "ended service")

	endpoints := service.NewLoanEndpoints(loanService)
	errors := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errors <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		handler := service.NewHTTPServer(endpoints)
		errors <- http.ListenAndServe(":8080", handler)
	}()

	level.Error(logger).Log("exit", <-errors)
}

func getEnvironmentVariable(key string) string {
	return os.Getenv(key)
}
