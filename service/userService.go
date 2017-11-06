package service

import (
	"fmt"
	"os"

	"github.com/James-Yip/Agenda/entity"
)

var CurUser string

func init() {
	var err error
	CurUser, err = entity.GetCurUser()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func IsUserRegistered(userName string) bool {
	users := entity.GetUsers()
	for _, user := range users {
		if user.UserName == userName {
			return true
		}
	}
	return false
}

func Register(userName string, password string, email string, phone string) {
	if CurUser != "" {
		fmt.Println("You have been registered！Please don't register repeatedly！")
		fmt.Println("Current login user: " + CurUser)
		return
	}
	if userName != "" && password != "" && email != "" && phone != "" {
		index, _, _ := entity.GetUser(userName)
		if index >= 0 {
			fmt.Println("Username repeats")
		} else {
			fmt.Println("Username available，registering......")
			entity.AddUser(userName, password, email, phone)
			entity.UpdateCurUser(userName)
			fmt.Println("Register: success\nCurrent login user: " + userName)
		}
	} else {
		fmt.Println("Please check your input validity.")
	}
}

func Login(userName string, password string) {
	if CurUser != "" {
		fmt.Println("You have been logged-in！")
		fmt.Println("Current login user: " + CurUser)
		return
	}
	if userName != "" && password != "" {
		index, theUser, _ := entity.GetUser(userName)
		if index >= 0 {
			if theUser.Password == password {
				entity.UpdateCurUser(userName)
				fmt.Println("Login: success\nCurrent login user: " + userName)
			} else {
				fmt.Println("Login: fail\nPassword incorrect.")
			}
		} else {
			fmt.Println("Login: fail\nUser unfound.")
		}
	}

}

func Logout() {
	if CurUser == "" {
		fmt.Println("Logout: fail\nYou aren't logged-in.")
		return
	}
	entity.UpdateCurUser("")
	fmt.Println("Logout: success")
}

func ListUsers() {
	if CurUser == "" {
		fmt.Println("Please login first.")
	} else {
		users := entity.GetUsers()
		fmt.Println("All User Information:")
		fmt.Println("index  username  email  phone")

		for index, user := range users {
			fmt.Printf("%d | %s | %s | %s\n", index, user.UserName, user.Email, user.Phone)
		}
	}
}

func DeleteUser() {
	if CurUser == "" {
		fmt.Println("Please login first.")
		return
	}
	var err error = entity.DeleteUser(CurUser)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("DeleteUser(%s): success\n", CurUser)
		fmt.Println("Logout")
		entity.UpdateCurUser("")
	}
}
