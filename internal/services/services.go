package services

import (
	"github.com/tarkovskynik/Golang-ninja-project-3/internal/repository"
)

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
