package postgre

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
)

const (
	SelectLoginByToken = `
		SELECT login, password 
		FROM whatwasit.credentials
		WHERE access_hash=$1
		LIMIT 1
	;`

	InsertLogin = `
		INSERT INTO whatwasit.credentials
		(login, password, access_hash)
		VALUES ($1, $2, $3)  
		ON CONFLICT (access_hash)
			DO UPDATE SET 
				login = EXCLUDED.login,
				password = EXCLUDED.password
	;`

	DeleteLoginByToken = `
		DELETE FROM whatwasit.credentials
		WHERE access_hash=$1
	;`
)

type LoginObject struct {
	Login    string `field:"login"`
	Password string `field:"password"`
}

var ErrorLoginNotFound = errors.New("Login data was not found")

func (wr Wrapper) GetLogin(accessToken string) (LoginObject, error) {
	rows, err := wr.db.Query(SelectLoginByToken, accessToken)
	blankFormat := LoginObject{}
	if err != nil {
		return blankFormat, err
	}
	resArr := parseSQLRows(rows, blankFormat)
	if len(resArr) < 1 {
		return blankFormat, ErrorLoginNotFound
	}
	return *resArr[0], nil
}

func (wr Wrapper) SetLogin(login, password, accessToken string) error {
	_, err := wr.db.Exec(InsertLogin, login, password, accessToken)
	return err
}

func (wr Wrapper) DelLogin(accessToken string) error {
	_, err := wr.db.Exec(DeleteLoginByToken, accessToken)
	return err
}

func parseSQLRows[T any](rows *sql.Rows, outputFormat T) []*T {
	defer rows.Close()

	results := make([]*T, 0)
	i := 0
	for rows.Next() {
		results = append(results, new(T))
		fieldMap := ExtractFieldPointersIntoNamedMap(results[i])
		sqlColumns, err := rows.Columns()
		if err != nil {
			log.Panicln(err)
		}

		orderedPointersArr := make([]any, len(fieldMap))
		for i, column := range sqlColumns {
			orderedPointersArr[i] = fieldMap[column]
		}
		err = rows.Scan(orderedPointersArr...)
		if err != nil {
			log.Panicln(err)
		}
		i++
	}

	if err := rows.Err(); err != nil {
		log.Panicln(err)
	}
	return results
}

func ExtractFieldPointersIntoNamedMap[T any](in *T) map[string]any {
	fieldMap := make(map[string]any)
	iter := reflect.ValueOf(in).Elem()
	for i := 0; i < iter.NumField(); i++ {
		currPtr := iter.Field(i).Addr().Interface()

		columnName := iter.Type().Field(i).Tag.Get("field") // sql field tag
		if columnName == "" {
			log.Panicln(fmt.Errorf("Struct type %T doesn't provide the necessary field tags for successful sql parsing", *in))
		}

		fieldMap[columnName] = currPtr
	}
	return fieldMap
}
