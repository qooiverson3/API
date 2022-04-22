package service

import (
	"ces-api/pkg/model"
)

type InstanceService struct {
	Repo model.Repository
}

func NewInstanceService(r model.Repository) *InstanceService {
	return &InstanceService{r}
}

func (s *InstanceService) GetInstanceList(q model.GetInstanceForm) *[]model.Instance {
	return s.Repo.QueryInstance(q)
}

// Actions 機器開關機等操作
func (s *InstanceService) Actions(q model.ActionRequestBody) int64 {
	var state string

	switch q.State {
	case 1:
		state = "Boot"
	case 2:
		state = "Shut down"
	case 3:
		state = "Reboot"
	case 4:
		state = "Turn off"
	case 5:
		state = "Reset"
	}

	result, _ := s.Repo.UpdateInstance(q.UUID, state)
	return result
}
