package entity

import "github.com/gofrs/uuid"

const (
	ClubTableName = "club"
)

type Club struct {
	ClubId   uuid.UUID `gorm:"type:uuid;primary_key" json:"club_id"`
	ClubName string    `gorm:"type:varchar(100);not_null" json:"club_name"`
	Players  []*Player `gorm:"many2many:player_club" json:"players"`
	Managers *Manager  `gorm:"foreignKey:ManagerID" json:"managers"`
}

func NewClub(clubId uuid.UUID, clubName string, players []*Player, managers *Manager) *Club {
	return &Club{
		ClubId:   clubId,
		ClubName: clubName,
		Players:  players,
		Managers: managers,
	}
}

func (model *Club) TableName() string {
	return ClubTableName
}
