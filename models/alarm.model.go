package models

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mirzaRakha28/BahasaMata/db"
)

type Alarm struct {
	Id         int    `json:"id"`
	Id_difabel string `json:"id_difabel" validate:"required"`
	Id_perawat string `json:"id_perawat" validate:"required"`
	Title      string `json:"title " validate:"required"`
	Jam        string `json:"jam " validate:"required"`
}

func AddAlarm(id_difabel, id_perawat, title, jam string) (Response, error) {
	var res Response
	var obj Alarm
	var arrobj []Alarm
	val := validator.New()

	alarm := Alarm{
		Id_difabel: id_difabel,
		Id_perawat: id_perawat,
		Title:      title,
		Jam:        jam,
	}
	err := val.Struct(alarm)
	if err != nil {
		return res, err
	}
	con := db.CreateCon()
	sqlStatement := "SELECT * FROM alarm WHERE jam = ?"

	stmt, err := con.Query(sqlStatement, jam)
	for stmt.Next() {
		err = stmt.Scan(&obj.Id, &obj.Id_difabel, &obj.Id_perawat, &obj.Title, &obj.Jam)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	if obj.Id != 0 {
		res.Status = http.StatusConflict
		res.Message = "Data Sudah Ada"
		res.Data = arrobj

		return res, err
	}
	sqlStatement = "INSERT alarm (id_difabel, id_perawat,title , jam) VALUES (?, ?, ?,?)"

	stmt_insert, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt_insert.Exec(id_difabel, id_perawat, title, jam)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	sqlStatementData := "SELECT * FROM alarm WHERE id = ?"
	rows, err := con.Query(sqlStatementData, lastInsertedId)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Id_difabel, &obj.Id_perawat, &obj.Title, &obj.Jam)
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

func UpdateAlarm(id, id_difabel, id_perawat, title, jam string) (Response, error) {
	var res Response
	var obj Alarm
	var arrobj []Alarm
	val := validator.New()

	alarm := Alarm{
		Id_difabel: id_difabel,
		Id_perawat: id_perawat,
		Title:      title,
		Jam:        jam,
	}
	err := val.Struct(alarm)
	if err != nil {
		return res, err
	}
	con := db.CreateCon()
	sqlStatement := "SELECT * FROM alarm WHERE id = ?"

	stmt, err := con.Query(sqlStatement, id)
	for stmt.Next() {
		err = stmt.Scan(&obj.Id, &obj.Id_difabel, &obj.Id_perawat, &obj.Title, &obj.Jam)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	if obj.Id == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Data Tidak Ada"
		res.Data = arrobj

		return res, err
	}
	sqlStatement = "Update alarm SET id_difabel = ?, id_perawat = ?,title = ? , jam = ? WHERE Id=?"

	stmt_insert, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt_insert.Exec(id_difabel, id_perawat, title, jam, id)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Update Berhasil"
	res.Data = map[string]int64{
		"rows_affected": lastInsertedId,
	}

	return res, err
}
func FetchAlarm(id string) (Response, error) {
	var res Response
	var obj Alarm
	var arrobj []Alarm

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM alarm WHERE id_perawat = ?"

	stmt, err := con.Query(sqlStatement, id)
	for stmt.Next() {
		err = stmt.Scan(&obj.Id, &obj.Id_difabel, &obj.Id_perawat, &obj.Title, &obj.Jam)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	if obj.Id == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Data Tidak Ada"
		res.Data = arrobj

		return res, err
	} else {
		res.Status = http.StatusOK
		res.Message = "Data Alarm"
		res.Data = arrobj

		return res, err
	}
}
func DeleteAlarm(id string) (Response, error) {
	var res Response
	var obj Alarm
	var arrobj []Alarm

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM alarm WHERE id = ?"

	stmt, err := con.Query(sqlStatement, id)
	for stmt.Next() {
		err = stmt.Scan(&obj.Id, &obj.Id_difabel, &obj.Id_perawat, &obj.Title, &obj.Jam)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	if obj.Id == 0 {
		res.Status = http.StatusNotFound
		res.Message = "Data Tidak Ada"
		res.Data = arrobj

		return res, err
	}
	sqlStatement = "DELETE FROM alarm WHERE id = ?"

	stmt_insert, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt_insert.Exec(id)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Data Terhapus"
	res.Data = map[string]int64{
		"rows_affected": lastInsertedId,
	}

	return res, err
}
