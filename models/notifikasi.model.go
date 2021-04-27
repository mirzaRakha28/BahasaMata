package models

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mirzaRakha28/BahasaMata/db"
)

type Notifikasi struct {
	Id_difabel string `json:"id_difabel" validate:"required"`
	Id_perawat string `json:"id_perawat" validate:"required"`
	Name       string `json:"name" `
}

func AddNotifikasi(id_difabel, id_perawat string) (Response, error) {
	var res Response
	var obj Notifikasi
	var arrobj []Notifikasi
	val := validator.New()

	notifikasi := Notifikasi{
		Id_difabel: id_difabel,
		Id_perawat: id_perawat,
	}
	err := val.Struct(notifikasi)
	if err != nil {
		return res, err
	}
	con := db.CreateCon()
	sqlStatement := "INSERT notifikasi (id_difabel, id_perawat) VALUES (?, ?)"

	stmt_insert, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	stmt_insert.Exec(id_difabel, id_perawat)

	sqlStatementData := "SELECT * FROM notifikasi WHERE id_difabel = ? AND id_perawat = ?"
	rows, err := con.Query(sqlStatementData, id_difabel, id_perawat)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	var data string
	sqlStatementData = "select name from difabel where id = " + id_difabel
	err = con.QueryRow(sqlStatementData).Scan(&data)
	for rows.Next() {
		obj.Name = data
		err = rows.Scan(&obj.Id_difabel, &obj.Id_perawat)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	if err != nil {
		return res, err
	}
	res.Status = http.StatusContinue
	res.Message = "Notifikasi"
	res.Data = arrobj

	return res, err
}
