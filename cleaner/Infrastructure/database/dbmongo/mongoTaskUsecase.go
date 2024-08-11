package dbmongo

import (
	"context"
	"example/cleaner/domain"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mts *MongoRepo) GetAllTasks() []*domain.Task {
	fmt.Println("mongoTaskCollection---GetAllTasks")
	fmt.Println("mongoTaskCollection---GetAllTasks")
	findOption := options.Find()
	findOption.SetLimit(100)
	tasks := []*domain.Task{}

	cursor, err := mts.db.Collection(mts.collection).Find(context.TODO(), bson.D{}, findOption)
	if err != nil {
		fmt.Println("could not load all the tasks 1")
		return tasks
	}
	for cursor.Next(context.TODO()) {
		var task domain.Task
		err := cursor.Decode(&task)
		if err != nil {
			fmt.Println(err)
		} else {
			tasks = append(tasks, &task)

		}
	}
	return tasks

}
func (mts *MongoRepo) GetTaskById(id int) (*domain.Task, error) {
	fmt.Println("mongoTaskCollection---GetTaskById")
	fmt.Println("mongoTaskCollection---GetTaskById")

	var task domain.Task
	filter := bson.M{"id": id}
	err := mts.db.Collection(mts.collection).FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		fmt.Println("could not find a result")
		return &task, err
	}
	return &task, nil

}
func (mts *MongoRepo) CreateTask(task domain.Task) (string, error) {
	fmt.Println("mongoTaskCollection---CreateTask")
	fmt.Println("mongoTaskCollection---CreateTask")

	_, err := mts.GetTaskById(task.Id)
	if err != nil {
		result, err := mts.db.Collection(mts.collection).InsertOne(context.TODO(), task)
		if err != nil {
			return "can't add the task", err
		}
		fmt.Println("this is the result id", result.InsertedID)
		return "Sucessfully added the task", nil
	}
	return "invalid request id is taken", err
}
func (mts *MongoRepo) UpdateTask(id int, updateBson bson.M) error {
	fmt.Println("mongoTaskCollection---UpdateTask")
	fmt.Println("mongoTaskCollection---UpdateTask")

	_, err := mts.GetTaskById(id)
	if err != nil {

		return err
	}
	filter := bson.M{
		"id": id,
	}

	fmt.Println("this is the filter: ", filter)
	fmt.Println("this is the update: ", updateBson)

	result, err := mts.db.Collection(mts.collection).UpdateOne(context.TODO(), filter, updateBson)
	if err != nil {
		return err
	}
	fmt.Println("update is sucessful")
	fmt.Println(result)
	return nil
}
func (mts *MongoRepo) DeleteTask(id int) error {
	fmt.Println("mongoTaskCollection---DeleteTask")
	fmt.Println("mongoTaskCollection---DeleteTask")

	_, err := mts.GetTaskById(id)
	if err != nil {
		return err
	}
	filter := bson.D{{
		Key: "id", Value: id,
	}}
	_, err = mts.db.Collection(mts.collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil

}

func (mts *MongoRepo) FilterTask(filter bson.M) []*domain.Task {
	fmt.Println("mongoTaskCollection---FilterTask")
	fmt.Println("mongoTaskCollection---FilterTask")

	findOptions := options.Find()
	findOptions.SetLimit(100)

	fmt.Println("this is the filter", filter)
	cur, err := mts.db.Collection(mts.collection).Find(context.TODO(), filter, findOptions)
	if err != nil {
		return []*domain.Task{}
	}
	result := []*domain.Task{}
	for cur.Next(context.TODO()) {
		var elem domain.Task
		err := cur.Decode(&elem)
		if err != nil {
			return []*domain.Task{}
		}
		result = append(result, &elem)
	}
	if err := cur.Err(); err != nil {
		return []*domain.Task{}
	}
	cur.Close(context.TODO())
	return result

}
