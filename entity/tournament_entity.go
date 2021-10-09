package entity

import "github.com/gofrs/uuid"

const (
	TournamentTableName = "tournament"
)

type Tournament struct {
	TournamentId   uuid.UUID `gorm:"type:uuid;primary_key" json:"tournament_id"`
	TournamentName string    `gorm:"type:varchar(100);not_null" json:"tournament_name"`
	ClubID         uuid.UUID `gorm:"type:uuid;not_null" json:"club_id"`
	Clubs          *Club     `gorm:"foreignKey:ClubID" json:"clubs"`
}

func NewTournament(tournamentID uuid.UUID, tournamentName string, clubs *Club) *Tournament {
	return &Tournament{
		TournamentId:   tournamentID,
		TournamentName: tournamentName,
		Clubs:          clubs,
	}
}

func (model *Tournament) TableName() string {
	return TournamentTableName
}
