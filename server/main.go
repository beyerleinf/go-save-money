package server

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "root"
	password = "example"
	dbname = "gosavemoney"
)

var db *sql.DB

func main() {
	http.StatusOK
}