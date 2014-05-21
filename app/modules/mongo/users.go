package mongo

import (
	"errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"revor/app/models"
)

func GetUser(username, password string, mongoDb *mgo.Database) (*models.User, error) {
	user := models.User{}
	var collection = mongoDb.C("users")
	var err = collection.Find(bson.M{"username": username, "password": password}).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func ListUsers(mongoDb *mgo.Database) []models.User {
	collection := mongoDb.C("users")
	results := []models.User{}
	collection.Find(nil).All(&results)
	return results
}

func CreateUser(username, password, firstname, lastname string, admin bool, mongoDb *mgo.Database) error {
	user, err := GetUser(username, password, mongoDb)
	if user != nil {
		return err
	}
	var collection = mongoDb.C("users")
	//
	newUser := models.User{}
	newUser.Username = username
	newUser.FirstName = firstname
	newUser.LastName = lastname
	newUser.Password = password
	newUser.IsAdmin = admin
	err = collection.Insert(newUser)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(username, password, firstname, lastname string, admin bool, mongoDb *mgo.Database) error {
	user := models.User{}
	var collection = mongoDb.C("users")
	change := mgo.Change{
		Update: bson.M{"username": username, "lastname": lastname, "firstname": firstname,
			"isadmin": admin, "password": password},
		ReturnNew: true,
	}
	_, err := collection.Find(bson.M{"username": username}).Apply(change, &user)
	if err != nil {
		return err
	}
	if user.Username != username {
		return errors.New("Can't udpate the given user")
	}
	return nil
}

func DeleteUser(username, password string, mongoDb *mgo.Database) error {

	user, err := GetUser(username, password, mongoDb)
	if user == nil {
		return err
	}
	var collection = mongoDb.C("users")
	//
	var selector = bson.M{"username": username, "password": password}
	err = collection.Remove(selector)
	if err != nil {
		return err
	}

	return nil
}
