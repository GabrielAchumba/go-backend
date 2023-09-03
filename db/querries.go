package db

const CREATEUSERTABLE = `CREATE TABLE users (
	"firstname" TEXT,
	"lastname" TEXT,
	"password" TEXT,
	"username" TEXT,
	"isdelete" INTEGER
);`

const CREATEUSERROW = "insert into users(firstname, lastname, password, username, isdelete) values (?, ?, ?, ?, ?)"

const GETUSERS = "select * from users where isdelete = ?"

const GETUSER = "select * from users where lastname = ?"

const UPDATEUSER = "update users set firstname = ? where lastname = ?"

const DELETUSER = "delete from users where lastname = ?"

const DELETUSER2 = "update users set isdelete = ? where lastname = ?"
