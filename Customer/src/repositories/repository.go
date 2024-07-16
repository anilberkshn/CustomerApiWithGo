package repositories

import (
	upRequestModel "CustomerApi/Customer/src/entities/requestModels"
	"CustomerApi/Customer/src/entities/responseModels"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Repo struct {
	Collection *mongo.Collection
}

func NewRepository() *Repo {

	databaseURL := "mongodb://127.0.0.1:27017"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	if err != nil {
		log.Fatal("Error : " + err.Error())
	}

	collection := client.Database("goCustomersApi").Collection("ApiCollection")

	return &Repo{collection}
}

func (r *Repo) GetAll(limit, offset int64) (response []*responseModels.CustomerResponseModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1000)
	defer cancel()

	filter := bson.M{}
	opts := options.Find().SetLimit(limit).SetSkip(offset)
	var customers []*responseModels.CustomerResponseModel

	cursor, err := r.Collection.Find(ctx, filter, opts)
	if err != nil {
		fmt.Println()
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) { // todo : cursor all çalıştırılmaya çalışılacak.
		var customer responseModels.CustomerResponseModel
		if err := cursor.Decode(&customer); err != nil {
			fmt.Println(err, "- Getall repo")
			return nil, err
		}
		customers = append(customers, &customer)
	}
	/*  ??
	err = cursor.All(ctx, customers)
	if err != nil {
		fmt.Println(err, "- Getall repo")
		return nil, err
	}
	*/
	return customers, nil
}

func (r *Repo) Create(m bson.M) (insertedID *string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := r.Collection.InsertOne(ctx, m)

	if err != nil {
		fmt.Println(err, "- Create repo")
		return nil, err
	}
	insertedId := fmt.Sprintf("%v", res.InsertedID)
	return &insertedId, err

}

func (r Repo) GetById(id string) (*responseModels.CustomerResponseModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var customer = responseModels.CustomerResponseModel{}

	filter := bson.M{"_id": id}
	opts := options.FindOne()
	err := r.Collection.FindOne(ctx, filter, opts).Decode(&customer)
	if err != nil {
		fmt.Println(err, "- GetById repo")
		return nil, err
	}

	return &customer, nil
}

func (r Repo) Delete(id string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"_id": id}
	opts := options.Delete()
	_, err := r.Collection.DeleteOne(ctx, filter, opts)
	if err != nil {
		fmt.Println(err, "- Delete repo")
		return false, err
	}

	return true, nil
}

func (r *Repo) Update(id string, updateData *upRequestModel.UpdateRequestModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"_id": id}

	update := bson.D{
		{"$set", bson.D{
			{"first_name", updateData.FirstName},
			{"last_name", updateData.LastName},
			{"email", updateData.Email},
			{"phone", updateData.Phone},
			{"address", updateData.Address},
		}},
	}

	_, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err, "- Update repo")
		return err
	}
	return nil
}
