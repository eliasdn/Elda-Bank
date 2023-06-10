package server

import (
	"go.uber.org/fx"

	"github.com/eliasdn/Elda-Bank/pkg/server/http"
)

func Module() fx.Option {
	return fx.Module("server", http.Module())
}
