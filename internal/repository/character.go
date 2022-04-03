package repository

import (
	"context"
	"dysn/character/internal/model"
	"dysn/character/internal/service/database"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CharacterRepoInterface interface {
	List(limit int, offset int) ([]*model.Character, error)
	Create(character *model.Character) (*model.Character, error)
	Update(id string, character *model.Character) (error)
	Delete(id string) error
	FindById(id string) (*model.Character, error)
	Exist(id string) (bool, error)
}

var collectionName = "characters"

type CharacterRepo struct {
	db *database.Mongodb
}

func NewCharacterRepo(db *database.Mongodb) *CharacterRepo {
	return &CharacterRepo{db}
}

func (c *CharacterRepo) List(limit int, offset int) ([]*model.Character, error) {
	var list []*model.Character
	filters := bson.D{{}}
	lt := int64(limit)
	skip := int64(offset)
	opts := &options.FindOptions{
		Skip:  &skip,
		Limit: &lt,
	}

	ctx := context.TODO()
	cur, err := c.db.Database.Collection(collectionName).Find(ctx, filters, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		character := model.NewCharacterIngot()
		err := cur.Decode(character)
		if err != nil {
			return nil, err
		}
		list = append(list, character)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return list, nil
}

func (c *CharacterRepo) Create(character *model.Character) (*model.Character, error) {
	character.ID = primitive.NewObjectID().Hex()
	result, err := c.db.Database.Collection(collectionName).InsertOne(context.TODO(), character)
	if err != nil {
		return nil, err
	}
	character.ID = fmt.Sprintf("%s", result.InsertedID)

	return character, nil
}

func (c *CharacterRepo) Update(id string, character *model.Character) error {
	updId := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.M{}
	if character.Name != "" {
		update["name"] = character.Name
	}
	if character.Description != "" {
		update["description"] = character.Description
	}
	if character.Status != 0 {
		update["status"] = character.Status
	}
	updateBson := bson.M{
		"$set": update,
	}

	return c.db.Database.
		Collection(collectionName).
		FindOneAndUpdate(context.TODO(), updId,updateBson).Err()
}

func (c *CharacterRepo) Delete(id string) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	result, err := c.db.Database.
		Collection(collectionName).
		DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no tasks were deleted")
	}

	return nil
}

func (c *CharacterRepo) FindById(id string) (*model.Character, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	character := model.NewCharacterIngot()

	err := c.db.Database.
		Collection(collectionName).
		FindOne(context.TODO(), filter).
		Decode(character)
	if err != nil {
		return nil, err
	}

	return character, err
}

func (c *CharacterRepo) Exist(id string) (bool, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	count, err := c.db.Database.Collection(collectionName).
		CountDocuments(context.TODO(), filter)

	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}
