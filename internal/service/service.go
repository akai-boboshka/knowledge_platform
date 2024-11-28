package service

import (
	"awesomeProject/internal/repository"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Repository repository.Repository
	Log        *logrus.Logger
}

func NewService(r repository.Repository, log *logrus.Logger) *Service {
	return &Service{Repository: r, Log: log}
}
