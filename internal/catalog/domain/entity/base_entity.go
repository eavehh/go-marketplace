package entity

import "github.com/google/uuid"

type Base_entity struct {
	Id    uuid.UUID
	Title string
}
