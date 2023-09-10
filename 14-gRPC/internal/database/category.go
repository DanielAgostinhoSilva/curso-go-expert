package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (props *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := props.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description)
	return Category{ID: id, Name: name, Description: description}, err
}

func (props *Category) FindAll() ([]Category, error) {
	rows, err := props.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []Category
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}
	return categories, err
}

func (props *Category) FindByCategoryId(categoryId string) (Category, error) {
	var id, name, description string
	query := "SELECT id, name, description FROM categories WHERE id = $1"
	err := props.db.QueryRow(query, categoryId).Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}
	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, err
}

func (props *Category) FindByCourseId(courseID string) (Category, error) {
	var id, name, description string
	query := "SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = $1"
	err := props.db.QueryRow(query, courseID).Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}
	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, err
}
