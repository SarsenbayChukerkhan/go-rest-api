package user

import (
	"context"
	"database/sql"
)

type PgRepository struct {
	db *sql.DB
}

func NewPgRepository(db *sql.DB) *PgRepository {
	return &PgRepository{db: db}
}

func (r *PgRepository) GetAll(ctx context.Context) ([]User, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name, email, age FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age); err != nil {
			return nil, err
		}
		list = append(list, u)
	}
	return list, nil
}

func (r *PgRepository) GetByID(ctx context.Context, id int64) (*User, error) {
	var u User
	err := r.db.QueryRowContext(ctx,
		`SELECT id, name, email, age FROM users WHERE id=$1`, id,
	).Scan(&u.ID, &u.Name, &u.Email, &u.Age)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &u, err
}

func (r *PgRepository) Create(ctx context.Context, u *User) error {
	return r.db.QueryRowContext(ctx,
		`INSERT INTO users (name, email, age) VALUES ($1,$2,$3) RETURNING id`,
		u.Name, u.Email, u.Age,
	).Scan(&u.ID)
}

func (r *PgRepository) Update(ctx context.Context, u *User) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE users SET name=$1,email=$2,age=$3 WHERE id=$4`,
		u.Name, u.Email, u.Age, u.ID,
	)
	return err
}

func (r *PgRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM users WHERE id=$1`, id)
	return err
}
