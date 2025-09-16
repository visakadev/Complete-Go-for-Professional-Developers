package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/visakadev/go/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	// Workout
	r.Get("/workout/{id}", app.WorkoutHandler.HandleGetWorkoutByID)
	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	r.Put("/workout/{id}", app.WorkoutHandler.HandleUpdateWorkoutByID)

	return r

}
