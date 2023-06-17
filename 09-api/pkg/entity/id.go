package entity

import "github.com/google/uuid"

type ID uuid.UUID

func NewId() ID {
	return ID(uuid.New())
}

func ParseID(id string) (ID, error) {
	idParse, err := uuid.Parse(id)
	return ID(idParse), err
}
