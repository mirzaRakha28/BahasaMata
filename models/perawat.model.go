package models

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mirzaRakha28/BahasaMata/db"
)

type Perawat struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

func RegisterPerawat(name, username, password string) (Response, error) {
	var res Response
	var obj Perawat
	var arrobj []Perawat
	val := validator.New()

	perawat := Perawat{
		Name:     name,
		Username: username,
		Password: password,
	}
	err := val.Struct(perawat)
	if err != nil {
		return res, err
	}
	con := db.CreateCon()
	sqlStatement := "SELECT * FROM perawat WHERE username = ?"

	stmt, err := con.Query(sqlStatement, username)
	for stmt.Next() {
		err = stmt.Scan(&obj.Id, &obj.Username, &obj.Password, &obj.Name)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	if obj.Id != 0 {
		res.Status = http.StatusConflict
		res.Message = "Username Sudah Tersedia"
		res.Data = nil
		return res, err
	}
	sqlStatement = "INSERT perawat ( name, username, password) VALUES (?, ?, ?)"

	stmt_insert, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt_insert.Exec(name, username, password)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	sqlStatementData := "SELECT * FROM perawat WHERE id = ?"
	rows, err := con.Query(sqlStatementData, lastInsertedId)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Username, &obj.Password, &obj.Name)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Registrasi Berhasil"
	res.Data = arrobj

	return res, err
}

func LoginPerawat(password, username string) (Response, error) {
	var res Response
	var obj Perawat
	var arrobj []Perawat

	con := db.CreateCon()
	sqlStatementData := "SELECT * FROM perawat WHERE username = ? "
	rows, err := con.Query(sqlStatementData, username)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Username, &obj.Password, &obj.Name)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	if obj.Id == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Username Tidak Tersedia"
		res.Data = nil
		return res, err
	}
	sqlStatementData = "SELECT * FROM perawat WHERE username = ? && password = ?"
	rows, err = con.Query(sqlStatementData, username, password)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Username, &obj.Password, &obj.Name)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	if obj.Id == 0 {
		res.Status = http.StatusUnauthorized
		res.Message = "Password Salah"
		res.Data = nil
		return res, err
	}
	fmt.Println(arrobj)
	res.Status = http.StatusOK
	res.Message = "Login Berhasil"
	res.Data = arrobj

	return res, err
}
