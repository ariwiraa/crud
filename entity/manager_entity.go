package entity

import "github.com/gofrs/uuid"

const (
	ManagerTableName = "Manager"
)

type Manager struct {
	ManagerId   uuid.UUID `gorm:"type:uuid;primary_key" json:"manager_id"`
	ManagerName string    `gorm:"type:varchar(100);not_null" json:"manager_name"`
}

func NewManager(managerID uuid.UUID, managerName string) *Manager {
	return &Manager{
		ManagerId:   managerID,
		ManagerName: managerName,
	}
}

func (model *Manager) TableName() string {
	return ManagerTableName
}
