package main

import (
	"context"
	"github.com/Waldemarsch/vk_it_diving/gateway"
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/handler"
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/infrastructure"
	"github.com/Waldemarsch/vk_it_diving/gateway/internal/service"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	inf := infrastructure.NewInfrastructure()
	srv := service.NewService(inf)
	hdlr := handler.NewHandler(srv)

	srvr := new(gateway.Server)
	go func() {
		if err := srvr.Run(viper.GetString("port"), hdlr.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Println("Gateway has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Gateway Shutting Down")

	if err := srvr.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
