package storage

import (
	"ces-api/pkg/model"

	"gorm.io/gorm"
)

type InstanceRepo struct {
	Db *gorm.DB
}

func NewInstanceRepo(r *gorm.DB) model.Repository {
	return &InstanceRepo{
		Db: r,
	}
}

func (r *InstanceRepo) GetInstanceList(dept string) *[]model.Instance {
	var instance []model.Instance
	r.Db.Table("orderList_sz").Where("a_dept = ? ", dept).Find(&instance)

	return &instance
}
