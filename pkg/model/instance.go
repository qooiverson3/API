package model

import (
	"time"
)

type Instance struct {
	A_id          int64     `json:"nid"`
	A_hostname    string    `json:"hostname"`
	A_cpu         int8      `json:"vcpu"`
	A_ram         int64     `json:"ram"`
	Networks      Networks  `gorm:"embedded" json:"ip_list"`
	Volumes       Volumes   `gorm:"embedded" json:"volume_space"`
	Monitor       Monitor   `gorm:"embedded" json:"monitor"`
	A_user        string    `json:"createBt"`
	A_time        time.Time `json:"createAt"`
	A_accept      string    `json:"confirmBy"`
	A_accept_time time.Time `json:"confirmAt"`
	A_drs         string    `json:"drs"`
	A_dept        string    `json:"dept"`
	A_uuid        string    `json:"UUID"`
	A_status      string    `json:"status"`
}

type Networks struct {
	A_ip1        string `json:"ip1"`
	A_ip2        string `json:"ip2"`
	A_netapp_ip  string `json:"storage"`
	A_loginsight string `json:"log"`
}

type Volumes struct {
	A_home_size    int64 `json:"home"`
	A_var_log_size int64 `json:"varLog"`
}

type Monitor struct {
	A_monitor_affect string `json:"effect"`
	A_monitor_use    string `json:"purpose"`
	A_monitor_port   string `json:"port"`
}

type XHeader struct {
	Token string `header:"token" binding:"required"`
}

type GetInstanceForm struct {
	Dept string `form:"dept" validate:"required"`
	Page int    `form:"page" validate:"required"`
}

type ActionRequestBody struct {
	UUID  string `json:"uuid" validate:"required"`
	State uint   `json:"state" validate:"required"`
}

type Repository interface {
	QueryInstance(q GetInstanceForm) *[]Instance
	UpdateInstance(uuid, state string) (int64, error)
}

type Service interface {
	GetInstanceList(s GetInstanceForm) *[]Instance
}
