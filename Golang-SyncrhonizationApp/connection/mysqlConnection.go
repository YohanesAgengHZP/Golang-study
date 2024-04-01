package mysqlConnection
import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
func ConnectionMysql() {
    db, err := sql.Open("mysql", "test_user:test_password@tcp(localhost:3306)/test_db")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    fmt.Println("Success connected!")
}