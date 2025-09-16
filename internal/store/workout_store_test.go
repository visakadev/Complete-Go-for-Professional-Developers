package store

// We never touch this testing db out of this file

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib" // Use 'stdlib' for database/sql compatibility
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("Opening test db: %v", err)

	}

	// run migrations for put test db
	err = Migrate(db, "../../migrations/")
	if err != nil {
		t.Fatalf("Migration test db: %v", err)
	}
	_, err = db.Exec(`TRUNCATE workouts, workout_entry CASCADE`)
	if err != nil {
		t.Fatalf("truncate test db error: %v", err)
	}
	return db

}

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewPostgresWorkoutStore(db)
	// Get as closed as the real deal as possible

	// Anonymous struct
	tests := []struct {
		name    string
		workout *Workout
		wantErr bool
	}{
		{
			name: "Valid Workout",
			workout: &Workout{
				Title:           "push day",
				Description:     "upper body day",
				DurationMinutes: 60,
				CaloriesBurned:  200,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Bench Press",
						Sets:         3,
						Reps:         IntPtr(10),
						Weight:       FloatPtr(135.5),
						Notes:        "warm up properly",
						OrderIndex:   1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Workout with invalid entries",
			workout: &Workout{
				Title:           "full body",
				Description:     "upper body day",
				DurationMinutes: 60,
				CaloriesBurned:  200,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Plank",
						Sets:         3,
						Reps:         IntPtr(60),
						Notes:        "warm up properly",
						OrderIndex:   1,
					},
					{
						ExerciseName:    "Squats",
						Sets:            4,
						Reps:            IntPtr(12),
						Notes:           "warm up properly",
						DurationSeconds: IntPtr(185.0),
						OrderIndex:      2,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		https: //frontendmasters.com/courses/complete-go/testing-createworkout-errors/
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := store.Create(tt.workout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func IntPtr(i int) *int {
	// extract the pointer of the provided value
	return &i
}
func FloatPtr(i float64) *float64 {
	return &i
}
