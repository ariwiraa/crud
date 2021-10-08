package entity

import (
	"github.com/gofrs/uuid"
)

const (
	PlayerTableName = "player"
)

type Player struct {
	PlayerID   uuid.UUID `gorm:"type:uuid;primary_key" json:"player_id"`
	PlayerName string    `gorm:"type:varchar(100);not_null" json:"player_name"`
	Clubs      []*Club   `gorm:"many2many:player_club" json:"clubs"`
	Country    *Country  `json:"country"`
}

func NewPlayer(playerID uuid.UUID, playerName string, clubs []*Club, country Country) *Player {
	return &Player{
		PlayerID:   playerID,
		PlayerName: playerName,
		Clubs:      clubs,
		Country:    &country,
	}
}

func (model *Player) TableName() string {
	return PlayerTableName
}
