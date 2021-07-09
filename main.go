package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/turugrura/codebkk-banking/handler"
	"github.com/turugrura/codebkk-banking/repository"
	"github.com/turugrura/codebkk-banking/service"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type SqlLogger struct {
	gormLogger.Interface
}

func (s SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n", sql)
}

func main() {
	initTimeZone()
	initConfig()
	db := initDatabaseGorm()

	custRepo := repository.NewCustomerRepositoryGorm(db)
	// custRepoMock := repository.NewCustomerRepositoryMock()
	// _ = custRepoMock
	custService := service.NewCustomerService(custRepo)
	custHandler := handler.NewCustomerHandler(custService)

	accRepo := repository.NewAccountRepositoryGorm(db)
	accService := service.NewAccountService(accRepo)
	accHandler := handler.NewAccountHandler(accService)

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		TimeZone: time.Local.String(),
	}))

	customerGroup := app.Group("/customers")
	customerGroup.Get("/", custHandler.GetCustomers)
	customerGroup.Get("/:customerID", custHandler.GetCustomer)

	customerGroup.Get("/:customerID/accounts", accHandler.GetAccounts)
	customerGroup.Post("/:customerID/accounts", accHandler.NewAccount)

	app.Get("/env", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"BaseURL":     c.BaseURL(),
			"Hostname":    c.Hostname(),
			"IP":          c.IP(),
			"IPs":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocal":    c.Protocol(),
			"Subdomains":  c.Subdomains(),
		})
	})

	addr := fmt.Sprintf(":%v", viper.GetString("app.port"))
	app.Listen(addr)
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

func initDatabaseSqlx() *sqlx.DB {
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

	return db
}

func initDatabaseGorm() *gorm.DB {
	dsn := fmt.Sprintf("server=%v;user id=%v;password=%v;database=%v;",
		viper.GetString("db.server"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
	)

	db, err := gorm.Open(sqlserver.Open(dsn))
	if err != nil {
		panic(err)
	}

	// if err = db.Ping(); err != nil {
	// 	panic(err)
	// }

	// db.SetConnMaxLifetime(3 * time.Minute)
	// db.SetMaxOpenConns(5)
	// db.SetMaxOpenConns(5)

	db.AutoMigrate(repository.Customer{}, repository.Account{})

	return db
}
