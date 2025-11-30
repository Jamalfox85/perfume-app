package api

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jamalfox85/perfume-app/backend/data"
)

type Application struct {
	DB 	  		*pgxpool.Pool
	Perfumes 	*data.PerfumeRepository
	Profiles    *data.ProfileRepository
	Cabinets    *data.CabinetRepository
}