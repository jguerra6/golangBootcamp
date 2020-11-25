package repository

import(
	"simpleAPI/entity"
)

type LeagueRepository interface {
	Save(league *entity.League) (*entity.League, error)
	GetAll() ([]entity.League, error)
}

