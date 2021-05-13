package main

import (
	"log"
	"net"
	"net/http"
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

type API int

func (a *API) Add(args Args, reply *int) error { //REtrun type error, 2 inputs //Exported function(UpperCAse)//Invert to method/ papi pointer
	*reply = args.Num1 + args.Num2
	return nil
}

func (a *API) GetTime(args Args, reply *time.Time) error {
	*reply = time.Now()
	return nil
}

func (a *API) Subtract(args Args, reply *int) error { //REtrun type error, 2 inputs //Exported function(UpperCAse)//Invert to method/ papi pointer//To group them
	*reply = args.Num1 - args.Num2
	return nil
}

var database []Item

func (a *API) GetDB(empty string, reply *[]Item) error {
	*reply = database
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Title == title {
			getItem = val
		}
	}

	*reply = getItem

	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(item Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.Title == item.Title {
			database[idx] = Item{item.Title, item.Body}
			changed = database[idx]
		}
	}

	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}

	*reply = del
	return nil
}

func main() {
	api := new(API)          //to know what to call
	err := rpc.Register(api) //register for clients to call them
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP() //Receive handler

	listener, err := net.Listen("tcp", "192.168.42.193:4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

	// fmt.Println("initial database: ", database)
	// a := Item{"first", "a test item"}
	// b := Item{"second", "a second item"}
	// c := Item{"third", "a third item"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("second database: ", database)

	// DeleteItem(b)
	// fmt.Println("third database: ", database)

	// EditItem("third", Item{"fourth", "a new item"})
	// fmt.Println("fourth database: ", database)

	// x := GetByName("fourth")
	// y := GetByName("first")
	// fmt.Println(x, y)

}
