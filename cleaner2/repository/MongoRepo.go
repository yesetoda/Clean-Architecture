package repo

import (
	"context"
	"example/cleaner2/domain"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	db         *mongo.Database
	collection string
}

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
func NewMongoRepository(db *mongo.Database, collection string) *MongoRepo {
	return &MongoRepo{
		db:         db,
		collection: collection,
	}
}
func NewCollection(dbname ,taskCollectionName string) *MongoRepo{
	client := GetNewMongoClient()
	return  NewMongoRepository(client.Database(dbname), taskCollectionName)

}
func NewMongoRepo() *MongoRepo{
	return &MongoRepo{}
}


// users
func (*MongoRepo)	GetAllUsers() []*domain.User{
	return []*domain.User{}
}
func (*MongoRepo)	GetUserByUsername(username string) (*domain.User, error){
	return &domain.User{},nil
}
func (*MongoRepo)	CreateUser(User domain.User) (string, error){
	return "",nil
}
func (*MongoRepo)	PromoteUser(username string, updateBson interface{}) error{
	return nil
}
func (*MongoRepo)	DeleteUser(username string) error{
	return nil
}
func (*MongoRepo)	FilterUser(filter interface{}) []*domain.User{
	return []*domain.User{}
}
func (*MongoRepo)	Login(username,password string) (string,error){
	return "",nil
}



// tasks
func (*MongoRepo) GetAllTasks() []*domain.Task{
	return []*domain.Task{}
}
func (*MongoRepo) GetTaskById(id int) (*domain.Task, error){
	return &domain.Task{},nil
}
func (*MongoRepo) CreateTask(task domain.Task) (string, error){
	return  "",nil
}
func (*MongoRepo) UpdateTask(id int, updateBson interface{}) error{
	return nil
}
func (*MongoRepo) DeleteTask(id int) error{
	return  nil
}
func (*MongoRepo) FilterTask(filter interface{}) []*domain.Task{
	return nil
}