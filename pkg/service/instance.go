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

// Actions 機器開關機等操作
func (s *InstanceService) Actions(uuid, state string) int64 {
	result, _ := s.Repo.UpdateInstance(uuid, state)
	return result
}
