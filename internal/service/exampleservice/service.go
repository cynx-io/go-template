package exampleservice

import (
	"github.com/cynx-io/micro-name/internal/repository/database"
)

type Service struct {
	TblExample *database.TblExample
}

func New(tblExample *database.TblExample) *Service {
	return &Service{
		TblExample: tblExample,
	}
}
