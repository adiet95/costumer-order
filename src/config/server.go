package config

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"

	"github.com/adiet95/costumer-order/src/routers"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start apllication",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	e := echo.New()
	if mainRoute, err := routers.New(e); err == nil {
		var addrs = os.Getenv("PORT")
		serve := &http.Server{
			Addr:    addrs,
			Handler: mainRoute,
		}
		err = mainRoute.StartServer(serve)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	} else {
		return err
	}
}
