//go:build integration
// +build integration

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/JesseNicholas00/EniqiloStore/utils/logging"
	testingUtils "github.com/JesseNicholas00/EniqiloStore/utils/testing"
	"github.com/JesseNicholas00/EniqiloStore/utils/validation"
	"github.com/labstack/echo/v4"
)

func TestMain(t *testing.T) {
	db := testingUtils.SetupTestDatabase(t)
	defer db.Close()

	mainInitLogger := logging.GetLogger("test", "main", "init")

	cfg, err := loadConfig()
	if err != nil {
		mainInitLogger.Fatal(err)
	}

	mainInitLogger.Printf("config loaded: %+v", cfg)

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

	go func() {
		server.Logger.Print(
			server.Start(
				fmt.Sprintf(
					"%s:%d",
					cfg.serverHost,
					cfg.serverPort,
				),
			),
		)
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			server.Logger.Print(err)
		}
	}()

	stop()
}
