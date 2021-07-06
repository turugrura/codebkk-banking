package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/turugrura/codebkk-banking/handler"
	"github.com/turugrura/codebkk-banking/logs"
	"github.com/turugrura/codebkk-banking/repository"
	"github.com/turugrura/codebkk-banking/service"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()

	custRepo := repository.NewCustomerRepositoryDB(db)
	// custRepoMock := repository.NewCustomerRepositoryMock()
	// _ = custRepoMock
	custService := service.NewCustomerService(custRepo)
	custHandler := handler.NewCustomerHandler(custService)

	accRepo := repository.NewAccountRepositoryDB(db)
	accService := service.NewAccountService(accRepo)
	accHandler := handler.NewAccountHandler(accService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", custHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", custHandler.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accHandler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accHandler.NewAccount).Methods(http.MethodPost)

	addr := fmt.Sprintf(":%v", viper.GetString("app.port"))
	logs.Info("Starting at port " + addr)
	http.ListenAndServe(addr, router)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("server=%v;user id=%v;password=%v;database=%v;",
		viper.GetString("db.server"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(5)
	db.SetMaxOpenConns(5)

	return db
}
