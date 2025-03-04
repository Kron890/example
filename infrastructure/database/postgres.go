package database

// подключения к бд
import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DBConn struct {
	*sql.DB
}

func NewDbConnection(connectionString string) (*DBConn, error) {
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &DBConn{conn}, nil
}
