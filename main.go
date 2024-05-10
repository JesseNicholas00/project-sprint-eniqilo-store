package main

import (
	"fmt"
	"net/url"

	"github.com/JesseNicholas00/EniqiloStore/controllers"
	authCtrl "github.com/JesseNicholas00/EniqiloStore/controllers/auth"
	customerCtrl "github.com/JesseNicholas00/EniqiloStore/controllers/customer"
	productCtrl "github.com/JesseNicholas00/EniqiloStore/controllers/product"
	"github.com/JesseNicholas00/EniqiloStore/middlewares"
	authRepo "github.com/JesseNicholas00/EniqiloStore/repos/auth"
	customerRepo "github.com/JesseNicholas00/EniqiloStore/repos/customer"
	productRepo "github.com/JesseNicholas00/EniqiloStore/repos/product"
	authSvc "github.com/JesseNicholas00/EniqiloStore/services/auth"
	customerSvc "github.com/JesseNicholas00/EniqiloStore/services/customer"
	productSvc "github.com/JesseNicholas00/EniqiloStore/services/product"
	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	"github.com/JesseNicholas00/EniqiloStore/utils/migration"
	"github.com/JesseNicholas00/EniqiloStore/utils/validation"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	serverHost string
	serverPort int

	dbString           string
	migrateDownOnStart bool
	migrateUpOnStart   bool

	jwtSecretKey   string
	bcryptSaltCost int
}

func loadConfig() (cfg ServerConfig, err error) {
	conf := viper.New()
	conf.SetConfigFile(".env")
	conf.SetConfigType("env")
	conf.AutomaticEnv()

	err = conf.ReadInConfig()
	if err != nil {
		return
	}

	conf.SetDefault("HOST", "0.0.0.0")
	conf.SetDefault("PORT", 8080)

	cfg.serverHost = conf.GetString("HOST")
	cfg.serverPort = conf.GetInt("PORT")

	cfg.dbString = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?%s",
		conf.GetString("DB_USERNAME"),
		// some passwords contain non-safe characters
		url.QueryEscape(conf.GetString("DB_PASSWORD")),
		conf.GetString("DB_HOST"),
		conf.GetInt("DB_PORT"),
		conf.GetString("DB_NAME"),
		conf.GetString("DB_PARAMS"),
	)
	cfg.migrateUpOnStart = conf.GetBool("DB_MIGRATE_UP_ON_START")
	cfg.migrateDownOnStart = conf.GetBool("DB_MIGRATE_DOWN_ON_START")

	cfg.jwtSecretKey = conf.GetString("JWT_SECRET")
	cfg.bcryptSaltCost = conf.GetInt("BCRYPT_SALT")

	return
}

func initControllers(
	cfg ServerConfig,
	db *sqlx.DB,
) (ctrls []controllers.Controller) {
	ctrlInitLogger := logging.GetLogger("main", "init", "controllers")
	defer func() {
		if r := recover(); r != nil {
			// add extra context to help debug potential panic
			ctrlInitLogger.Panicf("panic while initializing controllers: %s", r)
		}
	}()

	authRepo := authRepo.NewAuthRepository(db)
	authSvc := authSvc.NewAuthService(
		authRepo,
		cfg.jwtSecretKey,
		cfg.bcryptSaltCost,
	)
	authCtrl := authCtrl.NewAuthController(authSvc)
	authMw := middlewares.NewAuthMiddleware(authSvc)
	ctrls = append(ctrls, authCtrl)

	productRepo := productRepo.NewProductRepository(db)
	productSvc := productSvc.NewProductService(productRepo)
	productCtrl := productCtrl.NewProductController(productSvc, authMw)
	ctrls = append(ctrls, productCtrl)

	customerRepo := customerRepo.NewCustomerRepository(db)
	customerSvc := customerSvc.NewCustomerService(customerRepo)
	customerCtrl := customerCtrl.NewCustomerController(customerSvc, authMw)
	ctrls = append(ctrls, customerCtrl)

	return
}

func main() {
	mainInitLogger := logging.GetLogger("main", "init")

	cfg, err := loadConfig()
	if err != nil {
		mainInitLogger.Fatal(err)
	}

	mainInitLogger.Printf("config loaded: %+v", cfg)

	if cfg.migrateDownOnStart {
		if err := migration.MigrateDown(cfg.dbString, "migrations"); err != nil {
			mainInitLogger.Fatalf("failed to migrate down db: %s", err)
		}
	}
	if cfg.migrateUpOnStart {
		if err := migration.MigrateUp(cfg.dbString, "migrations"); err != nil {
			mainInitLogger.Fatalf("failed to migrate up db: %s", err)
		}
	}

	db, err := sqlx.Connect("postgres", cfg.dbString)
	if err != nil {
		mainInitLogger.Fatal(err)
	}
	defer db.Close()

	controllers := initControllers(cfg, db)

	server := echo.New()
	for idx, controller := range controllers {
		if err := controller.Register(server); err != nil {
			msg := fmt.Sprintf(
				"failed during controller registration (%d/%d): %s",
				idx+1,
				len(controllers),
				err,
			)
			mainInitLogger.Fatalf(msg)
		}
	}

	server.Validator = validation.NewEchoValidator()
	server.HideBanner = true

	server.Logger.Fatal(
		server.Start(
			fmt.Sprintf(
				"%s:%d",
				cfg.serverHost,
				cfg.serverPort,
			),
		),
	)
}
