package lib

import "go.uber.org/fx"

var Module = fx.Module("lib",
	fx.Provide(
		NewEnv,
		NewValidator,
		NewRequestHandler,
	),
)
