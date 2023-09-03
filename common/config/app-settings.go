package config

import (
	"fmt"
	"os"
	"time"
)

// git push heroku main
type Settings struct {
	Token struct {
		AccessTokenDuration   time.Duration
		RefreshTokenDuration  time.Duration
		TokenSecretKey        string
		RefreshTokenSecretKey string
	}
	Server struct {
		Port     string
		UserName string
		Password string
	}
	Database struct {
		DatabaseConnection string
		Databasename       string
	}
	Tables struct {
		User string
	}
}

var AppSettings = &Settings{}

func Setup() {
	AppSettings.Server.Port = os.Getenv("PORT")

	AppSettings.Database.DatabaseConnection = os.Getenv("DATABASECONNECTION")
	AppSettings.Database.Databasename = os.Getenv("DATABASENAME")

	AppSettings.Tables.User = os.Getenv("USER")

	fmt.Println("App settings was successfully loaded.")
}
