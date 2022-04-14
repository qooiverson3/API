package service

import "ces-api/pkg/model"

type InstanceService struct {
	Repo model.Repository
}

func NewInstanceService(r model.Repository) *InstanceService {
	return &InstanceService{r}
}

func (s *InstanceService) GetInstanceList(dept, page string) *[]model.Instance {
	return s.Repo.QueryInstance(dept, page)
}
