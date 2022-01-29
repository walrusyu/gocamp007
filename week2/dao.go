package week2

import (
	"database/sql"
)

func query() ([]interface{}, error) {
	err := sql.ErrNoRows
	if err == sql.ErrNoRows {
		return nil, myError{code: notFoundCode, msg: "not found"}
	} else {
		return nil, myError{code: sysErrorCode, msg: "not found"}
	}
	var data []interface{}
	data = append(data, 1)
	return data, nil
}
