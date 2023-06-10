package infra

import (
	"go.uber.org/fx"

	"github.com/eliasdn/Elda-Bank/pkg/infra/db"
)

func Module() fx.Option {
	return fx.Module("infra", db.Module())
}
