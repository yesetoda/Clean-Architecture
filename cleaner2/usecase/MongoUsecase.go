package usecase

import (
	"context"
	"errors"
	"example/cleaner2/domain"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUsecase struct {
}
func NewMongoUsecase() *MongoUsecase{
	return &MongoUsecase{
		
	}
}

// the task usecase comcrete implementation

func GetNewMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(os.Getenv("MongodbUri"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
func NewCollection(dbname ,taskCollectionName string) *mongo.Collection{
	client := GetNewMongoClient()
	return  client.Database(dbname).Collection(taskCollectionName)

}

var mongoTaskCollection = NewCollection(os.Getenv("MongodbName"),os.Getenv("TaskCollectionName"))
var mongoUserCollection = NewCollection(os.Getenv("MongodbName"),os.Getenv("UserCollectionName"))

func (mts *MongoUsecase) GetAllTasks() []*domain.Task {
	fmt.Println("mongoTaskCollection---GetAllTasks")
	fmt.Println("mongoTaskCollection---GetAllTasks")
	findOption := options.Find()
	findOption.SetLimit(100)
	tasks := []*domain.Task{}

	cursor, err := mongoTaskCollection.Find(context.TODO(), bson.D{}, findOption)
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
func (mts *MongoUsecase) GetTaskById(id int) (*domain.Task, error) {
	fmt.Println("mongoTaskCollection---GetTaskById")
	fmt.Println("mongoTaskCollection---GetTaskById")

	var task domain.Task
	filter := bson.M{"id": id}
	err := mongoTaskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		fmt.Println("could not find a result")
		return &task, err
	}
	return &task, nil

}
func (mts *MongoUsecase) CreateTask(task domain.Task) (string, error) {
	fmt.Println("mongoTaskCollection---CreateTask")
	fmt.Println("mongoTaskCollection---CreateTask")

	_, err := mts.GetTaskById(task.Id)
	if err != nil {
		result, err := mongoTaskCollection.InsertOne(context.TODO(), task)
		if err != nil {
			return "can't add the task", err
		}
		fmt.Println("this is the result id", result.InsertedID)
		return "Sucessfully added the task", nil
	}
	return "invalid request id is taken", err
}
func (mts *MongoUsecase) UpdateTask(id int, updateBson interface{}) error {
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

	result, err := mongoTaskCollection.UpdateOne(context.TODO(), filter, updateBson)
	if err != nil {
		return err
	}
	fmt.Println("update is sucessful")
	fmt.Println(result)
	return nil
}
func (mts *MongoUsecase) DeleteTask(id int) error {
	fmt.Println("mongoTaskCollection---DeleteTask")
	fmt.Println("mongoTaskCollection---DeleteTask")

	_, err := mts.GetTaskById(id)
	if err != nil {
		return err
	}
	filter := bson.D{{
		Key: "id", Value: id,
	}}
	_, err = mongoTaskCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil

}

func (mts *MongoUsecase) FilterTask(filter interface{}) []*domain.Task {
	fmt.Println("mongoTaskCollection---FilterTask")
	fmt.Println("mongoTaskCollection---FilterTask")

	findOptions := options.Find()
	findOptions.SetLimit(100)

	fmt.Println("this is the filter", filter)
	cur, err := mongoTaskCollection.Find(context.TODO(), filter, findOptions)
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

// the user usecase comcrete implementation


func (mts *MongoUsecase) GetAllUsers() []*domain.User {
	fmt.Println("mongoUserCollection---GetAllUsers")
	fmt.Println("mongoUserCollection---GetAllUsers")
	findOption := options.Find()
	findOption.SetLimit(100)
	Users := []*domain.User{}

	cursor, err := mongoUserCollection.Find(context.TODO(), bson.D{}, findOption)
	if err != nil {
		fmt.Println("could not load all the Users 1")
		return Users
	}
	for cursor.Next(context.TODO()) {
		var User domain.User
		err := cursor.Decode(&User)
		if err != nil {
			fmt.Println(err)
		} else {
			Users = append(Users, &User)

		}
	}
	return Users

}
func (mts *MongoUsecase) GetUserByUsername(username string) (*domain.User, error) {
	fmt.Println("mongoUserCollection---GetUserByUsername---", username)
	fmt.Println("mongoUserCollection---GetUserByUsername")

	var User domain.User
	filter := bson.M{"username": username}
	err := mongoUserCollection.FindOne(context.TODO(), filter).Decode(&User)
	if err != nil {
		fmt.Println("could not find a result")
		return &User, err
	}
	return &User, nil

}
func (mts *MongoUsecase) CreateUser(User domain.User) (string, error) {
	fmt.Println("mongoUserCollection---CreateUser")
	fmt.Println("mongoUserCollection---CreateUser")

	_, err := mts.GetUserByUsername(User.Username)
	if err != nil {
		User.Role = "user"
		User.Password, err = domain.HashPassword(User.Password)
		if err != nil {
			return "can't add the User", err
		}
		cnt, _ := mongoUserCollection.CountDocuments(context.TODO(), bson.M{})
		if cnt == 0 {
			User.Role = "admin"
		}
		result, err := mongoUserCollection.InsertOne(context.TODO(), User)
		if err != nil {
			return "can't add the User", err
		}
		fmt.Println("this is the result Username", result.InsertedID)
		return "Sucessfully added the User", nil
	}
	return "invalid  request Username is taken", err
}
func (mts *MongoUsecase) PromoteUser(username string, updateBson interface{}) error {
	fmt.Println("mongoUserCollection---UpdateUser", username)
	fmt.Println("mongoUserCollection---UpdateUser")

	r, err := mts.GetUserByUsername(username)
	fmt.Println("this is the result of get user by username", r)
	if err != nil {
		fmt.Println(err)
		return err
	}
	filter := bson.M{
		"username": username,
	}

	fmt.Println("this is the filter: ", filter)
	fmt.Println("this is the update: ", updateBson)

	result, err := mongoUserCollection.UpdateOne(context.TODO(), filter, updateBson)
	if err != nil {
		return err
	}
	fmt.Println("update is sucessful")
	fmt.Println(result)
	return nil
}
func (mts *MongoUsecase) DeleteUser(username string) error {
	fmt.Println("mongoUserCollection---DeleteUser")
	fmt.Println("mongoUserCollection---DeleteUser")

	_, err := mts.GetUserByUsername(username)
	if err != nil {
		return err
	}
	filter := bson.D{{
		Key: "username", Value: username,
	}}
	_, err = mongoUserCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil

}

func (mts *MongoUsecase) FilterUser(filter interface{}) []*domain.User {
	fmt.Println("mongoUserCollection---FilterUser")
	fmt.Println("mongoUserCollection---FilterUser")

	findOptions := options.Find()
	findOptions.SetLimit(100)

	fmt.Println("this is the filter", filter)
	cur, err := mongoUserCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return []*domain.User{}
	}
	result := []*domain.User{}
	for cur.Next(context.TODO()) {
		var elem domain.User
		err := cur.Decode(&elem)
		if err != nil {
			return []*domain.User{}
		}
		result = append(result, &elem)
	}
	if err := cur.Err(); err != nil {
		return []*domain.User{}
	}
	cur.Close(context.TODO())
	return result

}

func (mts *MongoUsecase) Login(username, password string) (string, error) {
	fmt.Println("mongoUserCollection---Login")
	fmt.Println("mongoUserCollection---Login")

	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.M{
		"username": username,
	}
	// 	cursor := userCollection.FindOne(context.TODO(), filter)
	cursor := mongoUserCollection.FindOne(context.TODO(), filter)

	// 	// Check user credentials
	var user domain.User
	err := cursor.Decode(&user)
	// fmt.Println("this is the user", user)
	if err != nil {
		return "", errors.New("no such user")
	}
	if !domain.VerifyPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	tokenString, err := domain.GenerateToken(username,user.Role)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
