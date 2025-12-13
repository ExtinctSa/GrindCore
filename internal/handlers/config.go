package handler

import (
	"github.com/ExtinctSa/final_project/internal/database"
)

type ApiConfig struct {
	Platform  string
	Sk        string
	DBQueries *database.Queries
}
