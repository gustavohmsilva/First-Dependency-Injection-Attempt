package datasource

import (
	"database/sql"

	"github.com/gustavohmsilva/test-dependency-injection/model"
	// Driver for SQLite3 Database
	_ "github.com/mattn/go-sqlite3"
)

// UserDataSource ...
type UserDataSource struct {
	Db *sql.DB
}

// NewUserDataSource ...
func NewUserDataSource(dbFile string) (UserDataSource, error) {
	database, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return UserDataSource{}, err
	}
	var uds UserDataSource
	uds.Db = database
	return uds, nil
}

// SelectLatestUser ...
func (uds *UserDataSource) SelectLatestUser() (model.User, error) {
	r := uds.Db.QueryRow(`SELECT id, name FROM user ORDER BY id DESC LIMIT 1;`)
	var u model.User
	err := r.Scan(&u.ID, &u.Name)
	return u, err
}

// InsertUser ...
func (uds *UserDataSource) InsertUser(nu model.User) error {
	_, err := uds.Db.Exec(`INSERT INTO user (name) VALUES (?)`, nu.Name)
	if err != nil {
		return err
	}
	return nil
}
