package entity

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	UserName string
	Password string
	Email    string
	Phone    string
}

type Users []*User

var users Users

func UpdateCurUser(curUser string) error {
	// write current login user into file
	fout, err := os.Create("data/curUser")
	if err != nil {
		return errors.New("update current user fail: \n->" + err.Error())
	}
	defer fout.Close()
	if err != nil {
		return errors.New("update current user fail: \n->" + err.Error())
	}
	fmt.Fprintf(fout, "%s", curUser)
	return nil
}

func GetCurUser() (string, error) {
	// read current login user from file and return it
	fin, err := os.Open("data/curUser")
	if err != nil {
		return "", errors.New("get current user fail: \n->" + err.Error())
	}
	defer fin.Close()
	curUser, err := ioutil.ReadAll(fin)
	if err != nil {
		return "", errors.New("get current user fail: \n->" + err.Error())
	}
	return string(curUser), nil
}

func AddUser(userName string, password string, email string, phone string) {
	user := &User{userName, password, email, phone}
	users = append(users, user)
	userWriteToFile()
}

func DeleteUser(userName string) error {
	index, _, err := GetUser(userName)
	if err != nil {
		return errors.New("delete user fail: " + err.Error())
	} else {
		// delete user with given index
		users = append(users[:index], users[index+1:]...)
		userWriteToFile()
	}
	return nil
}

func GetUser(userName string) (int, *User, error) {
	for index, user := range users {
		if user.UserName == userName {
			return index, user, nil
		}
	}
	return -1, nil, errors.New("can not get user by " + userName)
}

func GetUsers() Users {
	return users
}

func userEncode(user *User) ([]byte, error) {
	encodedUser, err := json.Marshal(user)
	if err != nil {
		return nil, errors.New("user encode fail: " + err.Error())
	}
	return encodedUser, nil
}

func userDecode(encodedUser []byte) (*User, error) {
	var user User
	err := json.Unmarshal(encodedUser, &user)
	if err != nil {
		return nil, errors.New("user decode fail: " + err.Error())
	}
	return &user, nil
}

func userWriteToFile() error {
	fout, err := os.Create("data/userInfo")
	if err != nil {
		return errors.New("write to file fail: \n->" + err.Error())
	}
	defer fout.Close()

	for _, user := range users {
		encodedUser, err := userEncode(user)
		if err != nil {
			return errors.New("write to file fail: \n->" + err.Error())
		}
		fmt.Fprintf(fout, "%s\n", encodedUser)
	}

	return nil
}

func userReadFromFile() error {
	fin, err := os.Open("data/userInfo")
	if err != nil {
		return errors.New("user read from file fail: \n->" + err.Error())
	}
	defer fin.Close()

	scanner := bufio.NewScanner(fin)

	for scanner.Scan() {
		token := scanner.Text()
		user, err := userDecode([]byte(token))
		if err != nil {
			return errors.New("user read from file fail: \n->" + err.Error())
		}
		users = append(users, user)
	}

	return nil
}

func init() {
	err := userReadFromFile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
