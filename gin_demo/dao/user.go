package dao

import "fmt"

var database = map[string]string{
	"admin": "admin",
	"test":  "test",
}

func AddUser(username, password string) bool {
	if _, ok := database[username]; ok {
		return false
	}
	database[username] = password
	fmt.Printf("username: %s, password: %s\n", username, password)
	fmt.Printf("database: %v\n", database)
	return true
}
func FindUser(username, password string) bool {
	if _, ok := database[username]; !ok {
		return false
	}
	if database[username] != password {
		return false
	}
	return true
}
