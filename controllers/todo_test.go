package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kk3939/golang-http-todo-api/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mock_DB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(mysql.Dialector{Config: &mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}}, &gorm.Config{})
	return gormDb, mock, err
}

func Test_handler(t *testing.T) {
	t.Run("GET", func(t *testing.T) {
		if _, _, err := mock_DB(); err != nil {
			t.Error("mock is failed")
		}
		db, mock, _ := mock_DB()
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `todos`")).WithArgs().WillReturnRows(
			sqlmock.NewRows([]string{
				"Id",
				"Name",
				"Content",
			}))
		database.Db = db

		r, err := http.NewRequest(http.MethodGet, "/todos", nil)
		if err != nil {
			t.Errorf("HTTP call failed, %v", err)
		}
		w := httptest.NewRecorder()
		TodoHandler(w, r)
		if w.Code != 200 {
			byteArray, _ := io.ReadAll(w.Body)
			t.Errorf("request is not valid, %v", string(byteArray))
		}
	})
}
