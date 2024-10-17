package models

import (
	"net/http"
	"github.com/HenkCode/golang-restapi/db"
	"github.com/go-playground/validator/v10"
)

type Siswa struct {
	Id     int    `json:"id"`
	Nama   string `json:"name" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Nohp   string `json:"nohp" validate:"required"`
}

func FetchSiswa() (Response, error) {
	var obj Siswa
	var arrobj []Siswa
	var res Response

	connection := db.CreateConf()

	sqlStatement := "SELECT * FROM siswa"

	rows, err := connection.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Nohp)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj

	return res, nil
}

func StoreSiswa(nama, alamat, nohp string) (Response, error) {
	var res Response

	v := validator.New()
	siswa := Siswa{
		Nama: nama,
		Alamat: alamat,
		Nohp: nohp,
	}

	err := v.Struct(siswa)
	if err != nil {
		return res, err
	}

	connection := db.CreateConf()
	sqlStatement := "INSERT siswa (nama, alamat, nohp) VALUES (?, ?, ?)"

	statement, err := connection.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	
	result, err := statement.Exec(nama, alamat, nohp)
	if err != nil {
		return res, err
	}
	
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]int64{
		"last_insert_id": lastInsertedId,
	}

	return res, nil
}

func UpdateSiswa(id int, nama, alamat , nohp string) (Response, error) {
	var res Response
	connection := db.CreateConf()
	sqlStatement := "UPDATE siswa SET nama = ?, alamat = ?, nohp = ? WHERE id = ?"

	statement, err := connection.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	
	result, err := statement.Exec(nama, alamat, nohp, id)
	if err != nil {
		return res, err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affeted": rowsAffected,
	}

	return res, nil

}

func DeleteSiswa(id int) (Response, error) {
	var res Response
	connection := db.CreateConf()
	sqlStatement := "DELETE FROM siswa 	WHERE id = ?"

	statement, err := connection.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	
	result, err := statement.Exec(id)
	if err != nil {
		return res, err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil


}