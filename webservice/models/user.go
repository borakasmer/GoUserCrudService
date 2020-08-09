package models

import (
	"errors"
	"fmt"
)

type HTTPRequest struct {
	Method string
}

type User struct{
	ID int
	FirstName string
	LastName string
}

var(
	users []*User
	nextID= 1
)

func GetUsers() []*User{
	return users
}

func AddUser(u User) (User,error){
	if u.ID !=0 {
		return User{}, errors.New("This Is Not New User!")
	}
	u.ID=nextID
	nextID++
	users=append(users,&u)
	return u,nil
}

func AddUsers(u []User) error{
	for _,v :=range u {
		if v.ID !=0 || HasUserByName(v.FirstName+v.LastName){
			return errors.New("This Is Not New User!")
		}
		v.ID = nextID
		nextID++
		user := v
		users = append(users, &user)
	}
	return nil
}

func GetUserById(id int) (int,User,error) {
	for i, v := range users {
		if v.ID == id {
			return i,*v,nil
		}
	}
	return 0,User{},fmt.Errorf("No User Found With This ID")
}

func HasUserByName(fullName string) bool {
	for _, v := range users {
		if v.FirstName + v.LastName == fullName {
			return true
		}
	}
	return false
}

func UpdateUser(u User) (User,error){
	index, user,_ := GetUserById(u.ID)
	users[index]=&u
	return user, nil
}

func RemoveUserById(Id int) error {
	for i,u := range users{
		if u.ID==Id {
			users= append(users[:i],users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found",Id)
}

func WorkSwitch(htype string){
	req := HTTPRequest{Method: htype}
	switch req.Method {
	case "GET":
		println("Get Request")
	case "PUT":
		println("Put Request")
		fallthrough
	case "POST":
		println("post Request")
	default:
		println("Unhandel Method")
	}
}