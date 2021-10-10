package entity

import "github.com/gofrs/uuid"

const (
	CountryTableName = "country"
)

type Country struct {
	CountryId   uuid.UUID `gorm:"type:uuid;primary_key" json:"country_id"`
	PlayerID    uuid.UUID `gorm:"type:uuid;not_null" json:"player_id"`
	CountryName string    `gorm:"type:varchar(100);not_null" json:"country_name"`
	Players     *Player   `gorm:"foreignKey:PlayerID" json:"players"`
}

func NewCountry(countryID, playerID uuid.UUID, countryName string) *Country {
	return &Country{
		CountryId:   countryID,
		PlayerID:    playerID,
		CountryName: countryName,
	}
}

func (model *Country) TableName() string {
	return CountryTableName
}
