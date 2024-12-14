package main

import "fmt"

type Contact struct {
	name         string
	email        string
	age          int
	isSubscribed bool
}

func main() {
	allContacts := map[string]Contact{
		"allen":     {name: "Allen Robin Hubert", email: "allenallen@gmail.com", age: 29, isSubscribed: false},
		"nibi":      {name: "Nibimon", email: "nibimon@gmail.com", age: 23, isSubscribed: true},
		"sherinmol": {name: "Sherinmol", email: "sherin@gmail.com", age: 26, isSubscribed: false},
	}

	fmt.Println()
	fmt.Println()

	for k, v := range allContacts {

		var message string

		message = fmt.Sprintf("GET OUT %s. YOU PIECE OF SHIITT!!!", k)

		if v.isSubscribed {
			message = fmt.Sprintf("Welcome back, %s. You are subscribed.", k)
		}

		fmt.Println(message)
		fmt.Println(fmt.Sprintf("name: %s, email:%s", v.name, v.email))
		fmt.Println()
	}
}
