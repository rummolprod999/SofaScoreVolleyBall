package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Filelog string

var DirLog = "log_server_sofa_volley"
var FileDB = "sofa.db"
var SofaTable = "sofa"
var FileLog Filelog
var mutex sync.Mutex

func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_journal_mode=OFF&_synchronous=OFF", FileDB))
	return db, err
}

func Logging(args ...interface{}) {
	mutex.Lock()
	file, err := os.OpenFile(string(FileLog), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("Ошибка записи в файл лога", err)
		return
	}
	fmt.Fprintf(file, "%v  ", time.Now())
	for _, v := range args {

		fmt.Fprintf(file, " %v", v)
	}
	//fmt.Fprintf(file, " %s", UrlXml)
	fmt.Fprintln(file, "")
	mutex.Unlock()
}
func CreateLogFile() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dirlog := filepath.FromSlash(fmt.Sprintf("%s/%s", dir, DirLog))
	if _, err := os.Stat(dirlog); os.IsNotExist(err) {
		err := os.MkdirAll(dirlog, 0711)

		if err != nil {
			fmt.Println("Не могу создать папку для лога")
			os.Exit(1)
		}
	}
	t := time.Now()
	ft := t.Format("2006-01-02")
	FileLog = Filelog(filepath.FromSlash(fmt.Sprintf("%s/log_sofa_%v.log", dirlog, ft)))
}

func CreateNewDB() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fileDB := filepath.FromSlash(fmt.Sprintf("%s/%s", dir, FileDB))
	if _, err := os.Stat(fileDB); os.IsNotExist(err) {
		Logging(err)
		f, err := os.Create(fileDB)
		if err != nil {
			Logging(err)
			panic(err)
		}
		err = f.Chmod(0777)
		if err != nil {
			Logging(err)
			panic(err)
		}
		err = f.Close()
		if err != nil {
			Logging(err)
			panic(err)
		}
		db, err := DbConnection()
		if err != nil {
			Logging(err)
			panic(err)
		}
		defer db.Close()
		_, err = db.Exec(`CREATE TABLE "sofa" (
	"id"	INTEGER NOT NULL,
	"id_game"	TEXT,
	"period"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
)`)
		if err != nil {
			Logging(err)
			panic(err)
		}

		_, err = db.Exec(`CREATE INDEX "id_game" ON "sofa" (
	"id_game"
)`)
		if err != nil {
			Logging(err)
			panic(err)
		}
		_, err = db.Exec(`CREATE INDEX "period" ON "sofa" (
	"period"
)`)
		if err != nil {
			Logging(err)
			panic(err)
		}
		_, err = db.Exec(`CREATE UNIQUE INDEX "id" ON "sofa" (
	"id"
)`)
		if err != nil {
			Logging(err)
			panic(err)
		}
	}
}

func CreateEnv() {
	CreateLogFile()
	CreateNewDB()
}
