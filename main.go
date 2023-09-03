package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GabrielAchumba/go-backend/common/config"
	"github.com/GabrielAchumba/go-backend/db"
	userModule "github.com/GabrielAchumba/go-backend/user-module"
	userService "github.com/GabrielAchumba/go-backend/user-module/services"
	"github.com/joho/godotenv"

	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	/* "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref" */)

var (
	server         *gin.Engine
	ctx            context.Context
	configSettings config.Settings
	//mongoClient    *mongo.Client
	sqlClient *sql.DB
)

func Production() string {
	return os.Getenv("APP_ENV")
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	server = gin.Default()

	if Production() != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env files")
		}
	}

	ctx = context.TODO()
	config.Setup()
	configSettings = *config.AppSettings
	//------------------------------------Mongo DB Connection-----------//
	/* mongoConn := options.Client().ApplyURI(config.AppSettings.Database.DatabaseConnection)
	client, err := mongo.Connect(ctx, mongoConn)

	if err != nil {
		log.Fatal(err)
	}

	mongoClient = client
	err = mongoClient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connection established") */

	//-------------------------SQLite DB Connection------------------//
	_, err := os.Stat("movie-db")
	if err != nil {
		file, err := os.Create("movie-db")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		log.Println("movie-db created")
	}

	moviesDB, _ := sql.Open("sqlite3", "./movie-db.db")
	sqlClient = moviesDB

	//defer moviesDB.Close()

	db.CreateTable(sqlClient, db.CREATEUSERTABLE, "users")

}

func main() {
	server.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "*",
		RequestHeaders:  "*",
		ExposedHeaders:  "Content-Length",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	apiBaseName := server.Group("")

	_userService := userService.New(sqlClient)
	userModule.InjectService(_userService).RegisterRoutes(apiBaseName)

	port := config.AppSettings.Server.Port

	networkingServer := &http.Server{
		Addr:         ":" + port,
		Handler:      server,
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 1200 * time.Second,
	}

	fmt.Println("Networking service is running on port: " + port)
	log.Fatal(networkingServer.ListenAndServe())

}
