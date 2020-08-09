package main

import (
	"errors"
	"fmt"
	"net/http"
	"webservice/controllers"
)

const (
	first = iota + 6
	second  =  2 << iota
    third
)
func main() {
	/*
	var fourth = second
	fifth := &fourth
	fourth +=1

	fmt.Println(first,second,third)
	fmt.Println(*fifth)
	//fmt.Println(first,second)
	*/

	/*
	var i int = 42
	fmt.Println(i)
	//fmt.Println("Hello from a module, Goopers")

	var f float32=3.14
	fmt.Println(f)

	firstName := "Bora"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3,4)
	fmt.Println(c)

	r,im := real(c), imag(c)
	fmt.Println(r,im)
	*/
 /*var firstName *string =new(string)
 *firstName = "bora"
 fmt.Println(*firstName)*/

/*firstName :="Bora"
fmt.Println(firstName)

ptr := &firstName
 fmt.Println(ptr,*ptr)

firstName="Duru"
fmt.Println(ptr,*ptr)

fmt.Println(reflect.TypeOf(ptr))*/

/*	const pi int = 3
	fmt.Println(float32(pi) + 1.2)
*/

	/*arr2 := [3]int{1,2,3}
	var arr [3]int
	arr[0]=1
	arr[1]=2
	arr[2]=3
	fmt.Println(arr)
	fmt.Println(arr2)

	slice := arr[:]

	arr[1]=34
	slice[2]=16
	fmt.Println(slice)*/

	/*
	slice := []int{1,2,3}
	fmt.Println(slice)
	slice = append(slice,4,34,16)
	fmt.Println(slice)

	s2 :=slice[1:4]
	fmt.Println(s2)
	 */

	/* m :=map[string]int{"foo":42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	*/

	/*type user struct {
		ID int
		FirstName string
		LastName string
	}

	var u user
	u.ID=1
	u.FirstName="Bora"
	u.LastName="Kasmer"
	fmt.Println(u)

	u2 := user{
		ID:1,
		FirstName: "Duru",
		LastName: "Kasmer",
	}
	fmt.Println(u2)
	*/
	/*u := models.User{
		ID:2,
		FirstName: "Bora",
		LastName: "Kasmer",
	}
	fmt.Println(u)
	*/
/*
	var firstName *string =new(string)
	*firstName = "bora"
	fmt.Println(firstName)
	fmt.Println(*firstName)

	*firstName = "Kasmer"
	fmt.Println(firstName)
	fmt.Println(*firstName)
 */
	//port :=3000
	/*isStarted := startWebServer(port,2)*/
	/*
	_,error := startWebServer(port,2)
	p,error := startWebServer(port,2)
	fmt.Println(p)
	fmt.Println(error)
	 */

/*
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
	//models.AddUser(user2)

	models.AddUsers([]models.User{user,user2,user3})

	user4 := models.User{
		ID:3,
		FirstName: "Secil",
		LastName: "Kasmer",
	}

	models.UpdateUser(user4)
	Users := models.GetUsers()
	fmt.Println(Users[2].FirstName+ " "+Users[2].LastName)
	fmt.Println("——————————————————")
	models.RemoveUserById(2)
	RemovedUsers :=models.GetUsers()
	for _,u := range RemovedUsers{
		println(u.FirstName +" "+u.LastName)
	}
	models.WorkSwitch("POST")
 */
	 controllers.RegisterController()
	 http.ListenAndServe(":3000", nil)

	//var i int
	/*for i := 0 ;i <5; i++ {
		//i++
		fmt.Println("Bora",i)
		if i==3{
			break
		}
	}*/
/*
	var i int
	for{
		if i >= 5 {
			break
		}
		fmt.Println(i)
		i++
	}
 */

	//slice := []int{1,2,3}
	//for i := 0; i< len(slice); i++{
	//for i,v := range slice{
	//	fmt.Println(i,v)
	//}
	/*
	wellKnownPorts:=map[string]int{"http":80,"https":443}
	for _,v := range wellKnownPorts {
		fmt.Println(v)
	}
	 */
	//panic("There is an error on DB operations")

}

/*func startWebServer(port, numberOfRetries int) bool {*/
func startWebServer(port, numberOfRetries int) (int,error) {
	fmt.Println("Starting Server")
	fmt.Println("Server started on port",port)
	fmt.Println("Number of retries",numberOfRetries)

	return port,errors.New("Someting went wrong")
}
