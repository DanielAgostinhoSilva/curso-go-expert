package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (props Course) Create(name, description, categoryId string) (Course, error) {
	id := uuid.New().String()
	_, err := props.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)",
		id, name, description, categoryId,
	)
	if err != nil {
		return nil, err
	}
	return Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryId,
	}, err
}

func (props Course) FindAll() ([]Course, error) {
	rows, err := props.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []Course
	for rows.Next() {
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}
		courses = append(courses, Course{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryID:  categoryId,
		})
	}
	return courses, err
}
