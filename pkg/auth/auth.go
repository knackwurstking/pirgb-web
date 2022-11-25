// TODO: use "alice" repo
package auth

import (
	"database/sql"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Authorization struct {
	UserId string `json:"userId"` // telegram is starting with a "t..."
	ApiKey string `json:"apikey"` // uuid string
	Server string `json:"server"` // server name
	Time   int64  `json:"time"`
}

func NewAuthorization(userId string, server, apiKey string) *Authorization {
	return &Authorization{
		UserId: userId,
		ApiKey: apiKey,
		Server: server,
		Time:   time.Now().Unix(),
	}
}

type DB struct {
	Sql *sql.DB
}

func NewDB(dbPath string) *DB {
	// create sqlite3 database
	db, err := NewSqlite(dbPath)
	if err != nil {
		panic(err)
	}

	// prepare sqlite database for usage
	err = CreateSqlDB(db)
	if err != nil {
		panic(err)
	}

	return &DB{
		Sql: db,
	}
}

func (db *DB) Close() error {
	return db.Sql.Close()
}

func (db *DB) AddToRegister(auth *Authorization) error {
	query := `
		INSERT INTO register (UserId, Key, Server, Time) VALUES (?, ?, ?, ?)
	`

	_, err := db.Sql.Exec(query, auth.UserId, auth.ApiKey, auth.Server, auth.Time)
	return err
}

func (db *DB) RemoveFromRegister(userId, server string) error {
	query := `
		DELETE FROM register WHERE UserId=? AND Server=?
	`

	_, err := db.Sql.Exec(query, userId, server)
	return err
}

func (db *DB) GetFromRegister(userId, server, apiKey string) (*Authorization, error) {
	query := `
		SELECT UserId, Key, Server, Time FROM register WHERE UserId=? AND Server=? AND Key=?
	`

	row := db.Sql.QueryRow(query, userId, server, apiKey)
	var auth Authorization
	err := row.Scan(&auth.UserId, &auth.ApiKey, &auth.Server, &auth.Time)
	if err != nil {
		return nil, err
	}
	return &auth, err
}

func (db *DB) ExistsInRegister(userId, server, apiKey string) bool {
	query := `
		SELECT FROM register WHERE UserId=? AND Server=? AND Key=?
	`

	_, err := db.Sql.Exec(query, userId, server, apiKey)
	return err == nil
}

func CreateSqlDB(db *sql.DB) error {
	if db == nil {
		return nil
	}

	// TODO: Add table: "users" && "users_whitelist"
	query := `
		CREATE TABLE IF NOT EXISTS register(
			UserId TEXT,
			Key TEXT, 
			Server TEXT,
			Time INTEGER
		);
	`

	_, err := db.Exec(query)
	return err
}

func NewSqlite(path string) (db *sql.DB, err error) {
	// get directory from path
	err = os.Mkdir(filepath.Dir(path), 0700)
	if err != nil {
		if !os.IsExist(err) {
			return
		}
	}

	return sql.Open("sqlite3", path)
}
