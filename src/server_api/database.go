package server_api

import (
  "fmt"

  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)


var db *sql.DB
func GetDB() (*sql.DB) {
  fmt.Println(db)
  return db
}

func OpenDB() {
  _db, err := sql.Open("sqlite3", "file:dbfile.db")
  checkErr(err)
  db = _db // important!
}



/*
  sqlStmt := `
  CREATE TABLE user (first_name text, last_name text, study_year int, study_faculty text, age int, insta_user_id int not null primary key, insta_user_name text);
  delete from user;
  `
  _, err = db.Exec(sqlStmt)
  if err != nil {
    log.Printf("%q: %s\n", err, sqlStmt)
    return
  }
*/

