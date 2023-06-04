package routes

import "go.uber.org/fx"

var Module = fx.Module("routes",
	fx.Provide(
		NewUserRoutes,
		NewForumRoutes,
		NewThreadRoutes,
		NewRoutes,
	),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// Setup all routes
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}

// NewRoutes creates a new Routes instance with the provided routes
func NewRoutes(
	userRoutes UserRoutes,
	forumRoutes ForumRoutes,
	threadRoutes ThreadRoutes,
) Routes {
	return Routes{
		userRoutes,
		forumRoutes,
		threadRoutes,
	}
}
