package models

import (
	"golang-boilerplate/libs"
	"log"
	"net/http"
)

type Framework struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
	Type string `db:"type"`
}

func GetFrameworkList(frameworks *[]Framework) int {
	query := "SELECT * FROM framework"
	log.Println(query)

	err := libs.Db.Select(frameworks, query)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func GetFramework(framework *Framework, name string) int {
	query := "SELECT * FROM framework WHERE name = '" + name + "'"
	log.Println(query)

	err := libs.Db.Get(framework, query)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound
	}

	return http.StatusOK
}

func CreateFramework(framework *Framework) int {
	query := "INSERT INTO framework VALUES(NULL, ?, ?)"
	log.Println(query)

	tx, err := libs.Db.Begin()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}
	_, err = tx.Exec(query, framework.Name, framework.Type)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func ChangeFramework(framework *Framework, name string) int {
	query := "UPDATE framework SET type = ? WHERE name = ?"
	log.Println(query)

	tx, err := libs.Db.Begin()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}
	_, err = tx.Exec(query, framework.Type, framework.Name)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func DeleteFramework(name string) int {
	query := "DELETE FROM framework WHERE name = ?"
	log.Println(query)

	tx, err := libs.Db.Begin()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}
	_, err = tx.Exec(query, name)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	}

	return http.StatusOK
}
