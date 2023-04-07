package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra"
	"github.com/imantung/golang_webform_for_gsheet/internal/app/infra/di"
	"github.com/labstack/echo/v4"
)

var exitSigs = []os.Signal{syscall.SIGTERM, syscall.SIGINT}

func Start() {
	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, exitSigs...)

	go func() {
		defer func() { exitCh <- syscall.SIGTERM }()
		if err := di.Invoke(startApp); err != nil {
			log.Println(err.Error())
		}
	}()
	<-exitCh

	if err := di.Invoke(gracefulShutdown); err != nil {
		log.Println(err.Error())
	}
}

func startApp(
	e *echo.Echo,
	cfg *infra.Config,
) error {
	if err := di.Invoke(setRoute); err != nil {
		return err
	}

	return e.StartServer(&http.Server{
		Addr:         cfg.Address,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})
}

func gracefulShutdown(
	e *echo.Echo,
) {

	log.Print("Shuting down")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Println(err.Error())
	}

}
