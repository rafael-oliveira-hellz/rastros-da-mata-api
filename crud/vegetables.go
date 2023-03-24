package crud

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Vegetable struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name                 string             `bson:"name,omitempty" json:"name,omitempty"`
	Description          string             `bson:"description,omitempty" json:"description,omitempty"`
	DevelopmentEta       string             `bson:"development_eta,omitempty" json:"development_eta,omitempty"`
	IdealDevelopmentTemp string             `bson:"ideal_development_temperature,omitempty" json:"ideal_development_temperature,omitempty"`
	Harvest              string             `bson:"harvest,omitempty" json:"harvest,omitempty"`
	Sunlight             string             `bson:"sunlight,omitempty" json:"sunlight,omitempty"`
	Irrigation           string             `bson:"irrigation,omitempty" json:"irrigation,omitempty"`
	Planting             string             `bson:"planting,omitempty" json:"planting,omitempty"`
	ExtraInfo            string             `bson:"extra_info,omitempty" json:"extra_info,omitempty"`
	Observation          string             `bson:"observation,omitempty" json:"observation,omitempty"`
	ImagePath            string             `bson:"image_path,omitempty" json:"image_path,omitempty"`
}

func (v *Vegetable) Create(ctx context.Context, coll *mongo.Collection) error {
	res, err := coll.InsertOne(ctx, v)
	if err != nil {
		return err
	}
	v.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (v *Vegetable) Read(ctx context.Context, coll *mongo.Collection, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	err := coll.FindOne(ctx, filter).Decode(&v)
	if err != nil {
		return err
	}
	return nil
}

func (v *Vegetable) Update(ctx context.Context, db *mongo.Collection, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"name":                          v.Name,
			"description":                   v.Description,
			"development_eta":               v.DevelopmentEta,
			"ideal_development_temperature": v.IdealDevelopmentTemp,
			"harvest":                       v.Harvest,
			"sunlight":                      v.Sunlight,
			"irrigation":                    v.Irrigation,
			"planting":                      v.Planting,
			"extra_info":                    v.ExtraInfo,
			"observation":                   v.Observation,
			"image_path":                    v.ImagePath,
		},
	}

	_, err := db.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (v *Vegetable) Delete(db *mongo.Collection, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := db.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func GetAllVegetables(db *mongo.Collection, limit, offset int64) ([]Vegetable, error) {
	ctx := context.TODO()

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)

	cur, err := db.Find(ctx, bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cur, ctx)

	var vegetables []Vegetable
	for cur.Next(ctx) {
		var vegetable Vegetable
		if err := cur.Decode(&vegetable); err != nil {
			return nil, err
		}
		vegetables = append(vegetables, vegetable)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return vegetables, nil
}
