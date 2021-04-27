package models

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mirzaRakha28/BahasaMata/db"
)

type Pasien struct {
	Id_difabel string `json:"id_difabel" validate:"required"`
	Id_perawat string `json:"id_perawat" validate:"required"`
	Name       string `json:"name" `
}

func DeletePasien(id_difabel, id_perawat string) (Response, error) {
	var res Response

	val := validator.New()

	pasien := Pasien{
		Id_difabel: id_difabel,
		Id_perawat: id_perawat,
	}
	err := val.Struct(pasien)
	if err != nil {
		return res, err
	}
	con := db.CreateCon()
	sqlStatement := "DELETE FROM pasien WHERE id_difabel= ? AND id_perawat = ?"

	stmt_insert, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt_insert.Exec(id_difabel, id_perawat)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusContinue
	res.Message = "pasien"
	res.Data = map[string]int64{
		"rows_affected": lastInsertedId,
	}

	return res, err
}
func AddPasien(id_difabel, id_perawat string) (Response, error) {
	var res Response
	var obj Pasien
	var arrobj []Pasien
	val := validator.New()

	pasien := Pasien{
		Id_difabel: id_difabel,
		Id_perawat: id_perawat,
	}
	err := val.Struct(pasien)
	if err != nil {
		return res, err
	}
	con := db.CreateCon()
	sqlStatement := "INSERT pasien (id_difabel, id_perawat) VALUES (?, ?)"

	stmt_insert, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	stmt_insert.Exec(id_difabel, id_perawat)

	sqlStatementData := "SELECT * FROM pasien WHERE id_difabel = ? AND id_perawat = ?"
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
	res.Message = "pasien"
	res.Data = arrobj

	return res, err
}
