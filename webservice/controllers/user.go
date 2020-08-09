package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"webservice/models"
)

type userController struct{
	userIdPattern *regexp.Regexp
}

//Constructor
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request){
	//w.Write([]byte("Hello from the User Controller!"))
	if r.URL.Path=="/users"{
		switch r.Method {
		case http.MethodGet:
			uc.getAllUser(w,r)
		}
	} else if r.URL.Path=="/insert" {
		switch r.Method {
		case http.MethodGet:
			uc.insertUsers(w)
		}
	} else if strings.Contains(r.URL.Path,"/users/"){
		//w.Write([]byte("Hello from the User Controller!"))
		matches := uc.userIdPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) ==0{
			w.WriteHeader(http.StatusNotFound)
		}
		id,err := strconv.Atoi(matches[1])
		if err !=nil{
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			uc.getUserById(w,r,id)
		}
	} else {
		w.Write([]byte("Hello from the User Controller!"))
	}
}

func newUserController() *userController{
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

func (uc *userController) getAllUser(w http.ResponseWriter, r *http.Request){
	encodeResponseAsJSON(models.GetUsers(),w)
}

func (uc *userController) getUserById(w http.ResponseWriter, r *http.Request,id int){
	_, user,_ :=models.GetUserById(id)
	encodeResponseAsJSON(user,w)
}

func (uc *userController) insertUsers(w http.ResponseWriter){
	user := models.User{
		FirstName: "Bora",
		LastName: "Kasmer",
	}
	//models.AddUser(user)

	user2 := models.User{
		FirstName: "Duru",
		LastName: "Kasmer",
	}

	user3 := models.User{
		FirstName: "Arya",
		LastName: "Kasmer",
	}
	models.AddUsers([]models.User{user,user2,user3})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All Users Inserted!"))
}

func(uc *userController) parseRequest(r *http.Request)(models.User,error){
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err !=nil{
		return models.User{},err
	}

	return u,err
}