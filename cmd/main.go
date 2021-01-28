package main

import (
	"os"

	userrestapi "github.com/fautherdtd/user-restapi"

	"github.com/fautherdtd/user-restapi/pkg/handler"
	"github.com/fautherdtd/user-restapi/pkg/repository"
	"github.com/fautherdtd/user-restapi/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Main ...
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config.")
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf(err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(userrestapi.Server)

	if err := srv.Run(viper.GetString("api.port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf(err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
