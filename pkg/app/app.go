package app

import (
	"go.uber.org/fx"

	"github.com/eliasdn/Elda-Bank/pkg/app/appconfig"
	"github.com/eliasdn/Elda-Bank/pkg/app/appcontext"
	"github.com/eliasdn/Elda-Bank/pkg/controller"
	"github.com/eliasdn/Elda-Bank/pkg/infra"
	"github.com/eliasdn/Elda-Bank/pkg/repo"
	"github.com/eliasdn/Elda-Bank/pkg/server"
	"github.com/eliasdn/Elda-Bank/pkg/service"
	"github.com/eliasdn/Elda-Bank/pkg/x/logger"
	"github.com/eliasdn/Elda-Bank/pkg/x/logger/fxlogger"
)

func New(ctx appcontext.Ctx, additionalOpts ...fx.Option) *fx.App {
	conf, err := appconfig.Parse(ctx)
	if err != nil {
		panic(err)
	}

	// logger and configuration are the only two things that are not in the fx graph
	// because some other packages need them to be initialized before fx starts
	logger.Configure(conf)

	baseOpts := []fx.Option{
		fx.WithLogger(fxlogger.Logger),
		fx.Supply(conf),
		controller.Module(),
		infra.Module(),
		repo.Module(),
		service.Module(),
		server.Module(),
	}

	return fx.New(append(baseOpts, additionalOpts...)...)
}
