package ginrouter

// import (
// 	"context"
// 	"example/cleaner/domain"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type MongoTaskRepo struct {
// 	db         *mongo.Database
// 	collection string
// }

// func NewMongoTaskRepository(db *mongo.Database, collection string) *MongoTaskRepo {
// 	return &MongoTaskRepo{
// 		db:         db,
// 		collection: collection,
// 	}
// }

// func (mts *MongoTaskRepo) GetAllTasks() []*domain.Task {
// 	findOption := options.Find()
// 	findOption.SetLimit(100)
// 	Tasks := []*domain.Task{}

// 	cursor, err := mts.db.Collection(mts.collection).Find(context.TODO(), bson.D{}, findOption)
// 	if err != nil {
// 		fmt.Println("could not load all the Tasks 1")
// 		return Tasks
// 	}
// 	for cursor.Next(context.TODO()) {
// 		var Task domain.Task
// 		err := cursor.Decode(&Task)
// 		if err != nil {
// 			fmt.Println(err)
// 		} else {
// 			Tasks = append(Tasks, &Task)

// 		}
// 	}
// 	return Tasks

// }
// func (mts *MongoTaskRepo) GetTaskById(id int) (*domain.Task, error) {
// 	var Task domain.Task
// 	filter := bson.D{{Key: "id", Value: id}}
// 	err := mts.db.Collection(mts.collection).FindOne(context.TODO(), filter).Decode(&Task)
// 	if err != nil {
// 		fmt.Println("could not find a result")
// 		return &Task, err
// 	}
// 	return &Task, nil

// }
// func (mts *MongoTaskRepo) CreateTask(Task domain.Task) (string, error) {
// 	_, err := mts.GetTaskById(Task.Id)
// 	if err != nil {
// 		result, err := mts.db.Collection(mts.collection).InsertOne(context.TODO(), Task)
// 		if err != nil {
// 			return "can't add the Task", err
// 		}
// 		fmt.Println("this is the result id", result.InsertedID)
// 		return "Sucessfully added the Task", nil
// 	}
// 	return "invalid request id is taken", err
// }
// func (mts *MongoTaskRepo) UpdateTask(id int, updateBson bson.M) error {
// 	_, err := mts.GetTaskById(id)
// 	if err != nil {
// 		filter := bson.M{
// 			"id": id,
// 		}
// 		update := bson.M{
// 			"$set": updateBson,
// 		}
// 		fmt.Println("this is the filter: ", filter)
// 		fmt.Println("this is the update: ", update)

// 		result, err := mts.db.Collection(mts.collection).UpdateOne(context.TODO(), filter, update)
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println("update is sucessful")
// 		fmt.Println(result)
// 		return nil
// 	}
// 	return err
// }
// func (mts *MongoTaskRepo) DeleteTask(id int) error {
// 	_, err := mts.GetTaskById(id)
// 	if err != nil {
// 		filter := bson.M{
// 			"id": id,
// 		}
// 		_, err := mts.db.Collection(mts.collection).DeleteOne(context.TODO(), filter)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// 	return err

// }

// func (mts *MongoTaskRepo) FilterTask(filter bson.M) []*domain.Task {
// 	findOptions := options.Find()
// 	findOptions.SetLimit(100)

// 	fmt.Println("this is the filter", filter)
// 	cur, err := mts.db.Collection(mts.collection).Find(context.TODO(), filter, findOptions)
// 	if err != nil {
// 		return []*domain.Task{}
// 	}
// 	result := []*domain.Task{}
// 	for cur.Next(context.TODO()) {
// 		var elem domain.Task
// 		err := cur.Decode(&elem)
// 		if err != nil {
// 			return []*domain.Task{}
// 		}
// 		result = append(result, &elem)
// 	}
// 	if err := cur.Err(); err != nil {
// 		return []*domain.Task{}
// 	}
// 	cur.Close(context.TODO())
// 	return result

// }
