package user

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id int64) error
}
