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
