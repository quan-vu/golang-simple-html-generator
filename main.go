package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	type User struct {
		Name   string
		Coupon string
		Amount int64
	}
	user := User{
		Name:   "Quan Vu",
		Coupon: "AIX1000",
		Amount: 5000,
	}
	err = tpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

	// START - Write file
	f, err := os.Create("./index_gen.html")
	if err != nil {
			log.Println("create file: ", err)
			return
	}
	err = tpl.Execute(f, user)
	if err != nil {
			log.Print("execute: ", err)
			return
	}
	f.Close()
	// END - Write file
	
}
