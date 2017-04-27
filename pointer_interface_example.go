package main

import "fmt"

type CacheItem interface {
	GetId() string
}

//Application
type Application struct {
	id string
}

func (a *Application) GetId() string {
	return a.id
}

type User struct {
	name string
}

func (a User) GetId() string {
	return a.name
}

func main() {

	var var1 CacheItem = &Application{"Application"}
	//var var2 CacheItem = Application{"Application"}
	app := Application{"Application"}
	app.GetId()
	fmt.Println(var1.GetId()) //, var2.GetId())

	// var var4 CacheItem = &User{"neeraj"}
	// var var3 CacheItem = User{"neeraj"}
	// fmt.Println(var3.GetId(), var4.GetId())
}
