package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    _"github.com/go-sql-driver/mysql"
   )
func main() {
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")

    dns := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)
    db,err := sql.Open("mysql",dns)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    		var name string

    		err := db.QueryRow("SELECT name FROM users LIMIT 1").Scan(&name)
    		if err != nil {
    			fmt.Fprintf(w, "DB接続成功（データなし）")
    			return
    		}
    		fmt.Fprintf(w, "Hello, %s!", name)
    	})

    	fmt.Println("サーバーを8080番ポートで起動しています...")
    	http.ListenAndServe(":8080", nil)
    }