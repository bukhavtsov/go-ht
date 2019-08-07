package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
func Perform(args Arguments, writer io.Writer) (err error) {
	id := args["id"]
	operation := args["operation"]
	item := args["item"]
	fileName := args["fileName"]
	if operation == "" {
		return errors.New("-operation flag has to be specified")
	}
	if fileName == "" {
		return errors.New("-fileName flag has to be specified")
	}
	switch operation {
	case "getUsersList":
		list(args["fileName"], writer, &err)
	case "add":
		if item == "" {
			return errors.New("-item flag has to be specified")
		}
		err = add(item, fileName)
	case "remove":
		if id == "" {
			return errors.New("-id flag has to be specified")
		}
		var user User
		user, err = findUserById(id, fileName)
		if !isValidUser(user) {
			return errors.New("User with id " + id + " not found")
		}
		err = remove(id, fileName)
	case "findById":
		if id == "" {
			return errors.New("-id flag has to be specified")
		}
		var user User
		user, err = findUserById(id, fileName)
		var bytes []byte
		if isValidUser(user) {
			bytes, _ = json.Marshal(user)
		} else {
			bytes = []byte("")
		}
		_, err = writer.Write(bytes)
	default:
		return errors.New("Operation " + args["operation"] + " not allowed!")
	}
	return err
}

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
type Arguments map[string]string

func parseArgs() (args Arguments) {
	id := flag.String("id", "", "id")
	operation := flag.String("operation", "", "operation")
	item := flag.String("item", "", "item")
	fileName := flag.String("fileName", "", "fileName")
	flag.Parse()
	args = make(map[string]string)
	args["id"] = *id
	args["operation"] = *operation
	args["item"] = *item
	args["fileName"] = *fileName
	return
}

func add(item string, fileName string) error {
	user := User{}
	itemByte := []byte(item)
	if err := json.Unmarshal(itemByte, &user); err != nil {
		return err
	}
	users, err := getUsersList(fileName)
	if err != nil {
		return err
	}
	if !isUniqueUser(user.Id, users) {
		return errors.New("Item with id " + user.Id + " already exists")
	}
	users = append(users, user)
	err = saveUsers(users, fileName)
	return err
}
func getUsersList(fileName string) (users []User, err error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return users, err
	}
	userBytes, err := ioutil.ReadAll(file)
	defer file.Close()
	if err != nil {
		return users, err
	}
	err = json.Unmarshal(userBytes, &users)
	return
}

func list(fileName string, writer io.Writer, err *error) {
	var users []User
	users, *err = getUsersList(fileName)
	bytes, _ := json.Marshal(users)
	_, *err = writer.Write(bytes)
}
func remove(id string, fileName string) error {
	users, err := getUsersList(fileName)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
		}
	}
	err = saveUsers(users, fileName)
	return err
}
func findUserById(id string, fileName string) (User, error) {
	user := User{}
	users, err := getUsersList(fileName)
	if err != nil {
		return user, err
	}
	if err != nil {
		return user, err
	}
	for _, currentUser := range users {
		if currentUser.Id == id {
			return currentUser, err
		}
	}
	return user, err
}
func isValidUser(user User) bool {
	if user.Id != "" && user.Email != "" && user.Age != 0 {
		return true
	}
	return false
}
func isUniqueUser(id string, users []User) bool {
	for _, user := range users {
		if user.Id == id {
			return false
		}
	}
	return true
}

func saveUsers(users []User, fileName string) error {
	usersJson, err := json.Marshal(users)
	err = ioutil.WriteFile(fileName, usersJson, 0644)
	if err != nil {
		return err
	}
	return err
}
