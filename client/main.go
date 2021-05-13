package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Item struct {
	Title string
	Body  string
}
type Args struct {
	Num1 int
	Num2 int
}
type Tiiime struct {
}

func main() {
	var reply Item
	var db []Item

	var aaa int
	var bbb time.Time

	client, err := rpc.DialHTTP("tcp", "192.168.42.193:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"First", "A first item"}
	b := Item{"Second", "A second item"}
	c := Item{"Third", "A third item"}

	aab := Args{4, 55}
	bb := Tiiime{}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	client.Call("API.Add", aab, &aaa)
	client.Call("API.GetTime", bb, &bbb)

	fmt.Println("The sum is : ", aaa)
	fmt.Println("I expected 9 ")

	fmt.Println("The Time is : ", bbb)

	fmt.Println("Database: ", db)

	client.Call("API.EditItem", Item{"Second", "A new second item"}, &reply)

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database: ", db)

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("first item: ", reply)

}
