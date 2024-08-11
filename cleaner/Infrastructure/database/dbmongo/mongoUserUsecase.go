package dbmongo

import (
	"context"
	"errors"
	"example/cleaner/Infrastructure/hashing"
	"example/cleaner/domain"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mts *MongoRepo) GetAllUsers() []*domain.User {
	fmt.Println("mongoUserCollection---GetAllUsers")
	fmt.Println("mongoUserCollection---GetAllUsers")
	findOption := options.Find()
	findOption.SetLimit(100)
	Users := []*domain.User{}

	cursor, err := mts.db.Collection(mts.collection).Find(context.TODO(), bson.D{}, findOption)
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
func (mts *MongoRepo) GetUserByUsername(username string) (*domain.User, error) {
	fmt.Println("mongoUserCollection---GetUserByUsername---", username)
	fmt.Println("mongoUserCollection---GetUserByUsername")

	var User domain.User
	filter := bson.M{"username": username}
	err := mts.db.Collection(mts.collection).FindOne(context.TODO(), filter).Decode(&User)
	if err != nil {
		fmt.Println("could not find a result")
		return &User, err
	}
	return &User, nil

}
func (mts *MongoRepo) CreateUser(User domain.User) (string, error) {
	fmt.Println("mongoUserCollection---CreateUser")
	fmt.Println("mongoUserCollection---CreateUser")

	_, err := mts.GetUserByUsername(User.Username)
	if err != nil {
		User.Role = "user"
		User.Password, err = hashing.HashPassword(User.Password)
		if err != nil {
			return "can't add the User", err
		}
		cnt, _ := mts.db.Collection(mts.collection).CountDocuments(context.TODO(), bson.M{})
		if cnt == 0 {
			User.Role = "admin"
		}
		result, err := mts.db.Collection(mts.collection).InsertOne(context.TODO(), User)
		if err != nil {
			return "can't add the User", err
		}
		fmt.Println("this is the result Username", result.InsertedID)
		return "Sucessfully added the User", nil
	}
	return "invalid  request Username is taken", err
}
func (mts *MongoRepo) PromoteUser(username string, updateBson bson.M) error {
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

	result, err := mts.db.Collection(mts.collection).UpdateOne(context.TODO(), filter, updateBson)
	if err != nil {
		return err
	}
	fmt.Println("update is sucessful")
	fmt.Println(result)
	return nil
}
func (mts *MongoRepo) DeleteUser(username string) error {
	fmt.Println("mongoUserCollection---DeleteUser")
	fmt.Println("mongoUserCollection---DeleteUser")

	_, err := mts.GetUserByUsername(username)
	if err != nil {
		return err
	}
	filter := bson.D{{
		Key: "username", Value: username,
	}}
	_, err = mts.db.Collection(mts.collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil

}

func (mts *MongoRepo) FilterUser(filter bson.M) []*domain.User {
	fmt.Println("mongoUserCollection---FilterUser")
	fmt.Println("mongoUserCollection---FilterUser")

	findOptions := options.Find()
	findOptions.SetLimit(100)

	fmt.Println("this is the filter", filter)
	cur, err := mts.db.Collection(mts.collection).Find(context.TODO(), filter, findOptions)
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

func (mts *MongoRepo) Login(username, password string) (string, error) {
	fmt.Println("mongoUserCollection---Login")
	fmt.Println("mongoUserCollection---Login")

	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.M{
		"username": username,
	}
	// 	cursor := userCollection.FindOne(context.TODO(), filter)
	cursor := mts.db.Collection(mts.collection).FindOne(context.TODO(), filter)

	// 	// Check user credentials
	var user domain.User
	err := cursor.Decode(&user)
	// fmt.Println("this is the user", user)
	if err != nil {
		return "", errors.New("no such user")
	}
	if !hashing.VerifyPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
	})

	// 	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
