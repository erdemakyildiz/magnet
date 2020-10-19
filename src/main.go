package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
)

func main() {
	query := url.Values{}
	query.Add("app name", "MyAppName")

	u := &url.URL{
		Scheme:   "Bundle",
		User:     url.UserPassword("vanilya", "kB7sgdzbhXodTA"),
		Host:     fmt.Sprintf("%s:%d", "52.212.115.125", 1433),
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())

	if err != nil {
		log.Fatal(err)
	}


	result, _ := db.Prepare("select top 10 * from NewsChannels")

	row, _ := result.Query(db)

	for row.Next() {
		var name, logoFileName string

		_ = row.Scan(&name, &logoFileName)

		println(name + "," + logoFileName)
	}
}