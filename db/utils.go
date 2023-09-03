package db

import (
	"database/sql"
	"log"

	"github.com/GabrielAchumba/go-backend/user-module/dtos"
	"github.com/GabrielAchumba/go-backend/user-module/models"
	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(db *sql.DB, querry string, tableName string) {

	statement, err := db.Prepare(querry)

	log.Println("Creating " + tableName + " table")
	if err == nil {
		statement.Exec()
		log.Println(tableName + "table created")
	}
}

/* type User struct {
	CreatedAt time.Time `json:"createdAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
} */

func InsertPersonalDetails(db *sql.DB, querry string, personalDetails dtos.PersonalProfieDTO) {

	statement, err := db.Prepare(querry)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(personalDetails.FirstName, personalDetails.LastName, personalDetails.Password, personalDetails.Username,
		false)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertMovieRow(db *sql.DB, querry string, movie dtos.MoviesDTO) {

	statement, err := db.Prepare(querry)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(movie.Category, movie.Name, movie.Rating)
	if err == nil {
		log.Fatal(err)
	}
}

func GetUsers(db *sql.DB, querry string) []models.User {
	statement, err := db.Prepare(querry)

	if err != nil {
		log.Fatal(err)
	}

	isDelete := false
	rows, err := statement.Query(isDelete)

	if err != nil {
		log.Fatal(err)
	}

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		rows.Scan(&user.FirstName, &user.LastName, &user.Password, &user.Username)
		users = append(users, user)
	}

	return users
}

func GetUser(db *sql.DB, querry string, lastname string) models.User {
	statement, err := db.Prepare(querry)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := statement.Query(lastname)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	for rows.Next() {
		rows.Scan(&user.FirstName, &user.LastName, &user.Password, &user.Username)
	}

	return user
}

func UpdateUser(db *sql.DB, querry string, user models.User) {

	statement, err := db.Prepare(querry)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(user.FirstName, user.LastName)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUser2(db *sql.DB, querry string, user models.User) {

	statement, err := db.Prepare(querry)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(user.IsDelete, user.LastName)
	if err != nil {
		log.Fatal(err)
	}
}

func Deleteser(db *sql.DB, querry string, lastName string) {

	statement, err := db.Prepare(querry)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(lastName)
	if err != nil {
		log.Fatal(err)
	}
}
