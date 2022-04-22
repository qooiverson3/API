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

func (r *InstanceRepo) QueryInstance(q model.GetInstanceForm) *[]model.Instance {
	var instance []model.Instance
	r.Db.Table("orderList_sz").Where("a_dept = ? ", q.Dept).Offset((q.Page - 1) * 10).Limit(10).Find(&instance)

	return &instance
}

func (r *InstanceRepo) UpdateInstance(uuid, state string) (int64, error) {
	result := r.Db.Table("orderList_sz").Where("a_uuid = ?", uuid).Update("a_state", state)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}
	return result.RowsAffected, nil
}
