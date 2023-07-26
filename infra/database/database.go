package database

import (
	entites "github.com/byhammert/lj-go/entities/categories"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Conn               *mongo.Client
	CategoryRepository entites.CategoryRepository
}

func NewDatabase(conn *mongo.Client, cr entites.CategoryRepository) *Database {
	return &Database{
		Conn:               conn,
		CategoryRepository: cr,
	}
}
