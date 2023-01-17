package main

import (
	"fmt"
	"os"
	"strings"
)

func close_conns() { // Closes every connection made and stops the server.
	fmt.Println("\rStopping server...")
	sendToEveryone("Server will be stopped now.", info)
	for _, user := range users {
		user.conn.Close()
	}
	os.Exit(0)
}

func find_user(user string, users []*User) int { // Finds an user inside the users slice
	for cnt, i := range users {
		if i.username == user {
			return cnt
		}
	}
	return -1
}

func check_banned(ip string) bool { // Checks if an IP address is inside the banned slice
	for _, i := range banned_users {
		if strings.Split(i.address, ":")[0] == ip {
			return true
		}
	}
	return false
}

func remove_user(items []*User, item *User) []*User { // Removes a user from the users slice
	newitems := []*User{}
	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}

	return newitems
}
