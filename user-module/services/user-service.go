package services

import (
	"time"

	"database/sql"

	"github.com/GabrielAchumba/go-backend/common/conversion"
	"github.com/GabrielAchumba/go-backend/db"
	"github.com/GabrielAchumba/go-backend/user-module/dtos"
	"github.com/GabrielAchumba/go-backend/user-module/models"
	/* "github.com/GabrielAchumba/school-mgt-backend/authentication/token"
	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo" */)

type UserService interface {
	LoginUser(requestModel dtos.LoginDTO) (interface{}, error)
	CreateUser(personalProfile dtos.PersonalProfieDTO) (interface{}, error)
	GetUsers() (interface{}, error)
	GetUser(lastname string) (interface{}, error)
	UpdateUser(personalProfile dtos.PersonalProfieDTO) (interface{}, error)
	DeleteUser(lastName string) (interface{}, error)
	DeletUser2(lastName string) (interface{}, error)
}

type serviceImpl struct {
	sqlClient *sql.DB
}

func New(sqlClient *sql.DB) UserService {

	return &serviceImpl{
		sqlClient: sqlClient,
	}
}

func (impl serviceImpl) LoginUser(loginDTO dtos.LoginDTO) (interface{}, error) {

	return loginDTO, nil
}

func (impl serviceImpl) CreateUser(personalProfile dtos.PersonalProfieDTO) (interface{}, error) {
	var userModel models.User
	conversion.Convert(personalProfile, &userModel)
	userModel.CreatedAt = time.Now()
	userModel.IsDelete = false
	db.InsertPersonalDetails(impl.sqlClient, db.CREATEUSERROW, personalProfile)

	return userModel, nil

}

func (impl serviceImpl) GetUsers() (interface{}, error) {

	users := db.GetUsers(impl.sqlClient, db.GETUSERS)

	return users, nil

}

func (impl serviceImpl) GetUser(lastname string) (interface{}, error) {

	user := db.GetUser(impl.sqlClient, db.GETUSER, lastname)

	return user, nil

}

func (impl serviceImpl) UpdateUser(personalProfile dtos.PersonalProfieDTO) (interface{}, error) {

	var userModel models.User
	conversion.Convert(personalProfile, &userModel)

	user := db.GetUser(impl.sqlClient, db.GETUSER, personalProfile.LastName)

	user.FirstName = personalProfile.FirstName
	db.UpdateUser(impl.sqlClient, db.UPDATEUSER, user)

	return user, nil

}

func (impl serviceImpl) DeleteUser(lastName string) (interface{}, error) {

	db.Deleteser(impl.sqlClient, db.DELETUSER, lastName)

	return 0, nil

}

func (impl serviceImpl) DeletUser2(lastName string) (interface{}, error) {

	user := db.GetUser(impl.sqlClient, db.GETUSER, lastName)

	user.IsDelete = true
	db.UpdateUser2(impl.sqlClient, db.DELETUSER2, user)

	return user, nil

}
