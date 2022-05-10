/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sankar228/fsync/cmd"
)

func main() {
	cmd.Execute()
}

func init() {
	dbfile := "./sqllite-dc.db"
	sqlscriptF := "./dbsetup.sql"
	os.Remove(dbfile)
	log.Println("creating and loading the necessary configuration tables")

	sqldbfile, err := os.Create(dbfile)
	if err != nil {
		log.Fatalf("unable to create sqllite db file: %s\n", dbfile)
	}

	sqldbfile.Close()
	log.Printf("sqllite db created: %s\n", dbfile)

	sqldb, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal("unable to open db connection: %s\n", dbfile)
	}

	defer sqldb.Close()

	sqlscript, err := ioutil.ReadFile(sqlscriptF)
	if err != nil {
		log.Fatalf("unable to read sql script: %s\n", sqlscriptF)
	}

	sqlscritpstr := string(sqlscript)
	log.Println(sqlscritpstr)
	re, err := sqldb.Exec(sqlscritpstr)
	if err != nil {
		log.Fatalf("unable to execute the sqlscript: %s , err: %v\n", sqlscriptF, err)
	}

	resp, _ := re.RowsAffected()
	log.Printf("necessary table are created: %d\n", resp)
}
