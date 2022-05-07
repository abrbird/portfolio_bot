package repository

type Repository interface {
	User() UserRepository
}
