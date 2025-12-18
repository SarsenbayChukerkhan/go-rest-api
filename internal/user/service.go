package user

import "context"

type Service interface {
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id int64) error
}

type service struct {
	repo      Repository
	validator *Validator
}

func NewService(r Repository, v *Validator) Service {
	return &service{repo: r, validator: v}
}

func (s *service) GetAll(ctx context.Context) ([]User, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) GetByID(ctx context.Context, id int64) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) Create(ctx context.Context, u *User) error {
	if err := s.validator.Validate(u); err != nil {
		return err
	}
	return s.repo.Create(ctx, u)
}

func (s *service) Update(ctx context.Context, u *User) error {
	if err := s.validator.Validate(u); err != nil {
		return err
	}
	return s.repo.Update(ctx, u)
}

func (s *service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
