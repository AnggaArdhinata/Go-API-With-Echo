package models

import (
	"net/http"

	"github.com/AnggaArdhinata/k-stylehub/src/db"
)

type Member struct {
	Id_Member int    `json:"id"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	Skintype  string `json:"skin_type"`
	Skincolor string `json:"skin_color"`
}

func GetAllMember() (Response, error) {
	var obj Member
	var arrobj []Member
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `echo-db`.`member`"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_Member, &obj.Username, &obj.Gender, &obj.Skintype, &obj.Skincolor)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreMember(username string, gender string, skintype string, skincolor string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO `echo-db`.`member` (username, gender, skintype, skincolor) VALUES(?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, gender, skintype, skincolor)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}
	
	return res, nil
}

func UpdateMember(id_member int, username string, gender string, skintype string, skincolor string) (Response, error)  {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE `echo-db`.`member` SET username= ?, gender= ?, skintype= ?, skincolor= ? WHERE id_member= ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, gender, skintype, skincolor, id_member)
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

func DeleteMember(id_member int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM `echo-db`.`member` WHERE id_member=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id_member)
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
