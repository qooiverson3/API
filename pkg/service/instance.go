package service

import "ces-api/pkg/model"

type Service struct {
	Repo model.Repository
}

func NewService(r model.Repository) *Service {
	return &Service{r}
}

func (s *Service) GetInstanceList(dept string) *[]model.Instance {
	return s.Repo.GetInstanceList(dept)
}
