package cli

import (
	"context"

	"go.uber.org/fx"

	"github.com/eliasdn/Elda-Bank/pkg/app"
	"github.com/eliasdn/Elda-Bank/pkg/app/appcontext"
)

func Start(module fx.Option) {
	app.New(appcontext.Declare(appcontext.EnvCLI), module).Start(context.Background())
}
