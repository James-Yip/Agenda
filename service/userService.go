package service

import (
	"fmt"
	"github.com/James-Yip/Agenda/entity"
)

var CurUser string

func Register(userName string, password string, email string, phone string) {
	if(userName !=""&&password!=""&&email!=""&&phone!="") {
		index, _, _ := entity.GetUser(userName);
		if(index>=0) {
			fmt.Println("此用户名已被注册")
		} else {
			fmt.Println("用户名可使用，正在注册......")
			entity.AddUser(userName, password, email, phone)
			CurUser= userName
			fmt.Println("注册成功！当前用户是："+CurUser)
		}
	} else {
		fmt.Println("请检查输入信息是否正确")		
	}
}

func Login(userName string, password string) {
	if(userName !=""&&password!="") {
		index, theUser, _ := entity.GetUser(userName);
		if(index>=0) {
			if(theUser.Password == password) {
				CurUser= userName;
				fmt.Println("登录成功！当前用户是："+CurUser)
			} else {
				fmt.Println("密码错误！登录失败！")				
			}
		} else {			
			fmt.Println("用户名错误！登录失败！")
		}
	}

}

func Logout() {
	CurUser= ""
	fmt.Println("退出登录！")

}

func ListUsers() {
	if(CurUser== "") {
		fmt.Println("请先登录！！！")

	} else {
		users :=  entity.GetUsers()
		fmt.Println("所有用户信息：")
		fmt.Println("编号  用户名  电子邮箱  电话号码")

		for index, user := range users {
			fmt.Print(index)
			fmt.Print(" | ")
			fmt.Print(user.UserName)
			fmt.Print(" | ")
			fmt.Print(user.Email)
			fmt.Print(" | ")
			fmt.Print(user.Phone)
		}
	}
}

func DeleteUser() {
	var err error = entity.DeleteUser(CurUser)
	if(err!=nil) {
		fmt.Println(err)
	} else {
		fmt.Println("删除用户"+CurUser+"成功")
		fmt.Println("退出登陆")
		CurUser= ""
	}
}
