package entity

import "github.com/gofrs/uuid"

const (
	CountryTableName = "country"
)

type Country struct {
	CountryId   uuid.UUID `gorm:"type:uuid;primary_key" json:"country_id"`
	CountryName string    `gorm:"type:varchar(100);not_null" json:"country_name"`
	PlayerID    uuid.UUID `gorm:"type:uuid;not_null" json:"players_id"`
	Players     *Player   `gorm:"foreignKey:PlayerID" json:"players"`
}

func NewCountry(countryID uuid.UUID, countryName string) *Country {
	return &Country{
		CountryId:   countryID,
		CountryName: countryName,
	}
}

func (model *Country) TableName() string {
	return CountryTableName
}
