package data

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)



type Profile struct {
	UserId 			string   			`json:"user_id"`
	Email   		string   			`json:"email"`

}

type ProfileRepository struct {
	db *pgxpool.Pool
}

func NewProfileRepository(db *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (r *ProfileRepository) CheckProfileExists(ctx context.Context, email string) (bool, error) {

	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM profiles
			WHERE LOWER(email) = LOWER($1)
		)
	`


	err := r.db.QueryRow(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}