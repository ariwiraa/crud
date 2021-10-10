package entity

import "github.com/gofrs/uuid"

const (
	ClubTableName = "club"
)

type Club struct {
	ClubID    uuid.UUID `gorm:"type:uuid;primary_key" json:"club_id"`
	ManagerID uuid.UUID `gorm:"type:uuid;not_null" json:"manager_id"`
	ClubName  string    `gorm:"type:varchar(100);not_null" json:"club_name"`
	Players   *Player   `gorm:"many2many:player_club" json:"players"`
	Managers  *Manager  `gorm:"foreignKey:ManagerID" json:"managers"`
}

func NewClub(clubID, managerID uuid.UUID, clubName string, players *Player) *Club {
	return &Club{
		ClubID:    clubID,
		ManagerID: managerID,
		ClubName:  clubName,
		Players:   players,
	}
}

func (model *Club) TableName() string {
	return ClubTableName
}
