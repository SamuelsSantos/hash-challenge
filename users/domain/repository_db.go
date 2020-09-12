package domain

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/SamuelsSantos/product-discount-service/users/domain/pb"
	"github.com/golang/protobuf/ptypes"
)

//SQLRepo repository
type SQLRepo struct {
	db *sql.DB
}

//NewRepository create new repository
func NewRepository(db *sql.DB) *SQLRepo {
	return &SQLRepo{db}
}

//Close database connection
func (r *SQLRepo) Close() error {
	return r.db.Close()
}

// GetByID fetch user by ID
func (r *SQLRepo) GetByID(id string) (*pb.User, error) {

	stmt, err := r.db.Prepare(`select id, first_name, last_name, date_of_birth from public.user where id = $1`)
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

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
