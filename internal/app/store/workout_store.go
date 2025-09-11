package store

import "database/sql"

// Login to communicate with the database

type WorkoutEntry struct {
	ID              int     `json:"id"`
	ExerciseName    string  `json:"exercise_name"`
	Sets            int     `json:"sets"`
	Reps            *int    `json:"reps"` // Pointer to and int because it can be null
	DurationSeconds *int    `json:"duration_seconds"`
	Weight          float64 `json:"weight"` // in pounds
	Notes           string  `json:"notes"`
	OrderIndex      int     `json:"order_index"`
}

type Workout struct {
	ID              int            `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	DurationMinutes int            `json:"duration_minutes"`
	CaloriesBurned  int            `json:"calories_burned"`
	Entries         []WorkoutEntry `json:"entries,omitempty"`
}

type PostgresWorkoutStore struct {
	db *sql.DB
}

func NewPostgresWorkoutStore(db *sql.DB) *PostgresWorkoutStore {
	return &PostgresWorkoutStore{db: db}
}

type WorkoutStore interface {
	CreateWorkout(*Workout) (*Workout, error)
	GetWorkoutByID(id int64) (*Workout, error)
}

func (pg *PostgresWorkoutStore) CreateWorkout(workout *Workout) (*Workout, error) {
	// sql error
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query :=
		`INSERT INTO workouts (title, description, duration_minutes, calories_burned)
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`
	err = tx.QueryRow(query, workout.Title, workout.Description, workout.DurationMinutes, workout.CaloriesBurned).Scan()

	if err != nil {
		return nil, err
	}
	return nil, err
}
