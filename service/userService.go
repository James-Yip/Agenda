package service

import (
	"agenda/entity"
	"agenda/utils"
	"fmt"
)

username string

func Register(userName string, password string, email string, phone string) {
	if(userName !=""&&password!=""&&email!=""&&phone!="") {
		index, theUser, err := GetUser(userName);
		if(index>=0) {
			fmt.Println("此用户名已被注册")
		} else {
			fmt.Println("用户名可使用，正在注册......")
			AddUser(userName, password, email, phone)
			username = userName
			fmt.Println("注册成功！当前用户是："+username)
		}
	} else {
		fmt.Println("请检查输入信息是否正确")		
	}
}

func Login(userName string, password string) {
	if(userName !=""&&password!=""&&email!=""&&phone!="") {
		index, theUser, err := GetUser(userName);
		if(index>=0) {
			if(theUser.Password == password) {
				username=userName;
				fmt.Println("登录成功！当前用户是："+username)
			} else {
				fmt.Println("密码错误！登录失败！")				
			}
		} else {			
			fmt.Println("用户名错误！登录失败！")
		}
	}

}

func Logout() {
	username = ""
	fmt.Println("退出登录！")

}

func ListUsers() {
	if(username == "") {
		fmt.Println("请先登录！！！")

	} else {
		users :=  GetUsers()
		fmt.Println("所有用户信息：")
		fmt.Println("编号  用户名  电子邮箱  电话号码")

		for index, user := range users {
			fmt.Println(index+"  "+user.UserName+"  "+user.Email+"  "+user.Phone)
		}
	}
}

func DeleteUser() {
	err:=DeleteUser(username)
	if(err) {
		fmt.Println(err)
	} else {
		fmt.Println("删除用户"+username+"成功")
		username = ""
	}
}
