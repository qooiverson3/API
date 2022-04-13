package storage

import (
	"ces-api/pkg/model"
	"strconv"

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

func (r *InstanceRepo) GetInstanceList(dept, page string) *[]model.Instance {
	var instance []model.Instance
	offset, _ := strconv.Atoi(page)
	r.Db.Table("orderList_sz").Where("a_dept = ? ", dept).Offset((offset - 1) * 10).Limit(10).Find(&instance)

	return &instance
}
