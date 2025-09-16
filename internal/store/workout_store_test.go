package store

// We never touch this testing db out of this file

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib" // Use 'stdlib' for database/sql compatibility
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	_, err = db.Exec(`TRUNCATE workouts, workout_entries CASCADE`)
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
						Notes:        "random",
						OrderIndex:   1,
					},
					{
						ExerciseName:    "Squats",
						Sets:            4,
						Reps:            IntPtr(12),
						Notes:           "notes",
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
			createWorkout, err := store.CreateWorkout(tt.workout)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.workout.Title, createWorkout.Title)
			assert.Equal(t, tt.workout.Description, createWorkout.Description)
			assert.Equal(t, len(tt.workout.Entries), len(createWorkout.Entries))

			retrieve, err := store.GetWorkoutByID(int64(createWorkout.ID))
			require.NoError(t, err)
			assert.Equal(t, createWorkout.Title, retrieve.Title)
			assert.Equal(t, createWorkout.Description, retrieve.Description)
			assert.Equal(t, len(createWorkout.Entries), len(retrieve.Entries))

			// for i, entry := range createWorkout.Entries {
			// 	assert.Equal(t, entry.ExerciseName, retrieve.Entries[i].ExerciseName)
			// 	assert.Equal(t, entry.Sets, retrieve.Entries[i].Sets)
			// 	assert.Equal(t, entry.Reps, retrieve.Entries[i].Reps)
			// 	assert.Equal(t, entry.Weight, retrieve.Entries[i].Weight)
			// 	assert.Equal(t, entry.Notes, retrieve.Entries[i].Notes)
			// 	assert.Equal(t, entry.OrderIndex, retrieve.Entries[i].OrderIndex)
			// }
			for i := range createWorkout.Entries {
				assert.Equal(t, tt.workout.Entries[i].ExerciseName, retrieve.Entries[i].ExerciseName)
				assert.Equal(t, tt.workout.Entries[i].ExerciseName, retrieve.Entries[i].ExerciseName)
				assert.Equal(t, tt.workout.Entries[i].Sets, retrieve.Entries[i].Sets)
				assert.Equal(t, tt.workout.Entries[i].Reps, retrieve.Entries[i].Reps)
				assert.Equal(t, tt.workout.Entries[i].Weight, retrieve.Entries[i].Weight)
				assert.Equal(t, tt.workout.Entries[i].Notes, retrieve.Entries[i].Notes)
				assert.Equal(t, tt.workout.Entries[i].OrderIndex, retrieve.Entries[i].OrderIndex)
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
