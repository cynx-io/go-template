package app

import (
	pbhermes "github.com/cynx-io/micro-name/api/proto/gen/hermes"
	pbplutus "github.com/cynx-io/micro-name/api/proto/gen/plutus"
	"github.com/cynx-io/micro-name/internal/repository/database"
	"github.com/cynx-io/micro-name/internal/repository/micro"
)

type Repos struct {
	TblExample *database.TblExample

	HermesUserClient pbhermes.HermesUserServiceClient
	PlutusUserClient pbplutus.PaymentServiceClient
}

func NewRepos(dependencies *Dependencies) *Repos {
	return &Repos{
		TblExample:       database.NewTblExample(dependencies.DatabaseClient.Db),
		HermesUserClient: micro.NewHermesUserClient(),
		PlutusUserClient: micro.NewPlutusPaymentClient(),
	}
}
