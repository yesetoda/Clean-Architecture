package ginrouter

// import (
// 	"context"
// 	"errors"
// 	"example/cleaner/Infrastructure/hashing"
// 	"example/cleaner/domain"
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type MongoUserRepo struct {
// 	db         *mongo.Database
// 	collection string
// }

// func NewMongoUserRepository(db *mongo.Database, collection string) *MongoUserRepo {
// 	return &MongoUserRepo{
// 		db:         db,
// 		collection: collection,
// 	}
// }

// func (mts *MongoUserRepo) GetAllUsers() []*domain.User {
// 	findOption := options.Find()
// 	findOption.SetLimit(100)
// 	Users := []*domain.User{}

// 	cursor, err := mts.db.Collection(mts.collection).Find(context.TODO(), bson.D{}, findOption)
// 	if err != nil {
// 		fmt.Println("could not load all the Users 1")
// 		return Users
// 	}
// 	for cursor.Next(context.TODO()) {
// 		var User domain.User
// 		err := cursor.Decode(&User)
// 		if err != nil {
// 			fmt.Println(err)
// 		} else {
// 			Users = append(Users, &User)

// 		}
// 	}
// 	return Users

// }
// func (mts *MongoUserRepo) GetUserByUsername(Username string) (*domain.User, error) {
// 	var User domain.User
// 	filter := bson.M{"username": Username}
// 	err := mts.db.Collection(mts.collection).FindOne(context.TODO(), filter).Decode(&User)
// 	fmt.Println(filter, User)
// 	if err != nil {
// 		fmt.Println("could not find a result")
// 		return &User, err
// 	}

// 	return &User, nil
// }
// func (mts *MongoUserRepo) CreateUser(User domain.User) (string, error) {
// 	_, err := mts.GetUserByUsername(User.Username)
// 	if err != nil {
// 		result, err := mts.db.Collection(mts.collection).InsertOne(context.TODO(), User)
// 		if err != nil {
// 			return "can't add the User", err
// 		}
// 		fmt.Println("this is the result Username", result.InsertedID)
// 		return "Sucessfully added the User", nil
// 	}
// 	return "invalUsername request Username is taken", err
// }
// func (mts *MongoUserRepo) PromoteUser(Username string, updateBson bson.M) error {
// 	_, err := mts.GetUserByUsername(Username)
// 	if err != nil {
// 		filter := bson.M{
// 			"username": Username,
// 		}
// 		fmt.Println("this is the filter: ", filter)
// 		fmt.Println("this is the update: ", updateBson)

// 		result, err := mts.db.Collection(mts.collection).UpdateOne(context.TODO(), filter, updateBson)
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println("update is sucessful")
// 		fmt.Println(result)
// 		return nil
// 	}
// 	return err
// }
// func (mts *MongoUserRepo) DeleteUser(Username string) error {
// 	_, err := mts.GetUserByUsername(Username)
// 	if err != nil {
// 		filter := bson.D{{
// 			Key: "username", Value: Username,
// 		}}
// 		_, err := mts.db.Collection(mts.collection).DeleteOne(context.TODO(), filter)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// 	return err

// }

// func (mts *MongoUserRepo) FilterUser(filter bson.M) []*domain.User {
// 	findOptions := options.Find()
// 	findOptions.SetLimit(100)

// 	fmt.Println("this is the filter", filter)
// 	cur, err := mts.db.Collection(mts.collection).Find(context.TODO(), filter, findOptions)
// 	if err != nil {
// 		return []*domain.User{}
// 	}
// 	result := []*domain.User{}
// 	for cur.Next(context.TODO()) {
// 		var elem domain.User
// 		err := cur.Decode(&elem)
// 		if err != nil {
// 			return []*domain.User{}
// 		}
// 		result = append(result, &elem)
// 	}
// 	if err := cur.Err(); err != nil {
// 		return []*domain.User{}
// 	}
// 	cur.Close(context.TODO())
// 	return result

// }

// func (mts *MongoUserRepo) Login(username, password string) (string, error) {

// 	filter := bson.M{
// 		"username": username,
// 	}
// 	// 	cursor := userCollection.FindOne(context.TODO(), filter)
// 	cursor := mts.db.Collection(mts.collection).FindOne(context.TODO(), filter)

// 	// 	// Check user credentials
// 	var user domain.User
// 	err := cursor.Decode(&user)
// 	if err != nil {
// 		return "", err
// 	}
// 	if !hashing.VerifyPassword(password, user.Password) {
// 		return "", errors.New("invalid credentials")
// 	}

// 	// Create a new token object, specifying signing method and the claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"username": username,
// 		"role":     user.Role,
// 		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
// 	})

// 	// 	// Sign and get the complete encoded token as a string
// 	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
