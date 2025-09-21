package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/visakadev/go/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)
		// Workout
		r.Get("/workout/{id}", app.Middleware.RequiredUser(app.WorkoutHandler.HandleGetWorkoutByID))
		r.Post("/workouts", app.Middleware.RequiredUser(app.WorkoutHandler.HandleCreateWorkout))
		r.Put("/workout/{id}", app.Middleware.RequiredUser(app.WorkoutHandler.HandleUpdateWorkoutByID))
		r.Delete("/workout/{id}", app.Middleware.RequiredUser(app.WorkoutHandler.HandleDeleteWorkoutByID))
	})

	r.Get("/health", app.HealthCheck)
	// Users
	r.Post("/users", app.UserHandler.HandleRegisterUser)
	// Token
	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)

	return r

}
