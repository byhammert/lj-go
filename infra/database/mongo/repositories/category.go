package repositories

import (
	"context"

	entities "github.com/byhammert/lj-go/entities/categories"
	mongo_infra "github.com/byhammert/lj-go/infra/database/mongo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mongo_drive "go.mongodb.org/mongo-driver/mongo"
)

const (
	CategoryCollection = "categories"
)

type CategoryRepository struct {
	Ctx        context.Context
	Collection *mongo_drive.Collection
}

func NewCategoryRepository(ctx context.Context, client *mongo_drive.Client) *CategoryRepository {
	return &CategoryRepository{
		Ctx:        ctx,
		Collection: mongo_infra.GetCollection(ctx, client, CategoryCollection),
	}
}

func (sr CategoryRepository) Create(Category *entities.Category) error {
	_, err := sr.Collection.InsertOne(sr.Ctx, Category)

	return err
}

func (sr CategoryRepository) List() (Categorys []entities.Category, err error) {
	cur, err := sr.Collection.Find(sr.Ctx, bson.M{})
	if err != nil {
		return Categorys, err
	}

	for cur.Next(sr.Ctx) {
		var Category entities.Category

		if err = cur.Decode(&Category); err != nil {
			return Categorys, err
		}

		Categorys = append(Categorys, Category)
	}

	return Categorys, nil
}

func (sr CategoryRepository) FindByID(id uuid.UUID) (Category entities.Category, err error) {
	err = sr.Collection.FindOne(sr.Ctx, bson.M{"_id": id}).Decode(&Category)

	return Category, err
}

func (sr CategoryRepository) Update(Category *entities.Category) error {
	_, err := sr.Collection.UpdateOne(sr.Ctx, bson.M{"_id": Category.ID}, bson.M{"$set": Category})

	return err
}

func (sr CategoryRepository) Delete(id uuid.UUID) error {
	_, err := sr.Collection.DeleteOne(sr.Ctx, bson.M{"_id": id})

	return err
}
