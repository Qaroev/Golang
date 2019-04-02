

package main

import (
"database/sql"
"fmt"

_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
func main() {
db, err = sql.Open("mysql", "root:root@/nodejs")
fmt.Println("Database connected.",db)
}