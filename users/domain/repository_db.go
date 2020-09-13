package domain

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/SamuelsSantos/product-discount-service/users/config"
	"github.com/SamuelsSantos/product-discount-service/users/domain/pb"
	"github.com/golang/protobuf/ptypes"
)

//SQLRepo repository
type SQLRepo struct {
	Cfg *config.Config
}

//NewSQLRepository create new repository
func NewSQLRepository(cfg *config.Config) *SQLRepo {
	return &SQLRepo{cfg}
}

func newDBConnection(cfg *config.Config) (*sql.DB, error) {
	return sql.Open(cfg.Db.Driver, cfg.Db.ToURL())
}

// GetDB new db connection
func (r *SQLRepo) GetDB() (*sql.DB, error) {
	return newDBConnection(r.Cfg)
}

// GetByID fetch user by ID
func (r *SQLRepo) GetByID(id string) (*pb.User, error) {

	db, err := r.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`select id, first_name, last_name, date_of_birth from public.user where id = $1`)
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		pbUser, err := transform(rows)
		if err != nil {
			return nil, err
		}

		return pbUser, nil
	}

	return nil, errors.New("Not found")
}

func transform(r *sql.Rows) (*pb.User, error) {

	var id string
	var firstName string
	var lastName string
	var dateOfBirth time.Time

	err := r.Scan(&id, &firstName, &lastName, &dateOfBirth)
	if err != nil {
		return nil, err
	}

	dtProto, err := ptypes.TimestampProto(dateOfBirth)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:          id,
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: dtProto,
	}, nil
}
