package models

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

const dbuser = "user"
const dbpass = "password"
const dbname = "hotel_booking"

func GetAllUsers() []User {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	// if there is an error opening the connection, handle it
	if err != nil {

		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM users")

	if err != nil {

		fmt.Println("Err", err.Error())

		return nil

	}

	users := []User{}
	for results.Next() {

		var user User

		err = results.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users

}

func GetUser(id string) *User {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	user := &User{}

	if err != nil {

		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM users where id=?", id)

	if err != nil {

		fmt.Println("Err", err.Error())

		return nil
	}

	if results.Next() {

		err = results.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return user

}

func AddUser(user *User) Response {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	sqlstm := fmt.Sprintf("INSERT INTO users (id, name, email, password, phone_number, created_at)"+
		" VALUES ('%s','%s','%s','%s','%s', now())",
		user.Id, user.Name, user.Email, user.Password, user.PhoneNumber)
	insert, err := db.Query(sqlstm)

	if err != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}
	defer insert.Close()

	return Response{
		StatusCode: http.StatusOK,
		Message:    "CREATED_SUCCESS",
	}

}

func UpdateUser(user *User) Response {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	sqlstm := fmt.Sprintf("UPDATE users SET name = '%s', email = '%s', password = '%s', phone_number = '%s', updated_at = now() WHERE id = '%s'",
		user.Name, user.Email, user.Password, user.PhoneNumber, user.Id)
	update, err := db.Query(sqlstm)

	if err != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}
	defer update.Close()

	return Response{
		StatusCode: http.StatusOK,
		Message:    "UPDATE_SUCCESS",
	}

}

func DeleteUser(id string) Response {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	sqlstm := fmt.Sprintf("DELETE FROM users WHERE id = '%s'", id)
	del, err := db.Query(sqlstm)

	if err != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}
	defer del.Close()

	return Response{
		StatusCode: http.StatusOK,
		Message:    "DELETE_SUCCESS",
	}

}
