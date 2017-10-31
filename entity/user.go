package entity

type User struct {
	UserName string
	Password string
	Email    string
	Phone    string
}

type Users []*User

var users Users
var meetings Meetings

func (users *Users) AddUser(userName string, password string, email string, phone string) {

}

func DeleteUser(userName string) {

}

func GetUser(userName string) {

}

func init() {

}
