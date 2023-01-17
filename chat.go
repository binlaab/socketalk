package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var users []*User
var banned_users []User

const MAX_SIZE = 256

var info = User{
	username: "INFO",
}

var admin = User{
	username: "ADMIN",
}

type User struct {
	username string
	address  string
	conn     net.Conn
}

func eval_admin() {
	for {
		fmt.Print(">>> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		cmd_txt := scanner.Text()
		if strings.HasPrefix(cmd_txt, "!") {

			var cmd, val string
			if len(strings.Split(cmd_txt, " ")) >= 2 { // Avoids a panic
				val = strings.Split(cmd_txt, " ")[1]
			}
			cmd = strings.Split(cmd_txt, " ")[0]

			switch cmd {
			case "!list":
				fmt.Printf("Users connected: \n")
				for _, user := range users {
					fmt.Printf("%s\n", user.username)
				}

			case "!help":
				fmt.Print(help_admin)

			case "!ban":
				ban(val)

			case "!ip":
				fmt.Print(show_ip(val))

			case "!stop":
				close_conns()

			case "!tell":
				split_text := strings.Split(cmd_txt, " ")
				tell(strings.Join(split_text[2:], " "), split_text[1], admin)

			default:
				fmt.Printf("Command not found. Use !help for a list of commands.")
			}

		} else {
			sendToEveryone(cmd_txt, admin)
		}
	}
}

func sendToEveryone(msg string, sender User) {
	for _, user := range users {
		if user.address == sender.address {
			continue
		}
		user.conn.Write([]byte("\r[" + sender.username + "] " + msg + "\n"))
		user.conn.Write([]byte("[" + user.username + "] "))
	}
}

func handle_connection(conn net.Conn) {

	sendToEveryone((conn.RemoteAddr().String() + " has joined the chat."), info)
	message := make([]byte, MAX_SIZE)

	u := User{
		username: conn.RemoteAddr().String(),
		address:  conn.RemoteAddr().String(),
		conn:     conn,
	}
	users = append(users, &u)
	send := true
	conn.Write([]byte("[" + u.username + "] "))
	for {
		l, err := conn.Read(message)
		if !send {
			if l+1 < MAX_SIZE {
				conn.Write([]byte("[" + u.username + "] "))
			}
			continue
		}
		conn.Write([]byte("[" + u.username + "] "))
		var msg string

		if err != nil {
			if err == io.EOF || err == err.(net.Error) {
				sendToEveryone(u.username+" exited the chat", info)
				fmt.Printf("\r%s exited the chat.\n>>>", u.username)
				conn.Close()
				break
			}
			fmt.Println("Cannot read: ", err)
			break
		}

		if l+1 > MAX_SIZE {
			send = false
			conn.Write([]byte("\r[WARNING] This message is too long (max. size is 256 characters). It will not be sent.\n"))
			continue
		}
		for _, char := range message[0:l] {
			msg += string(char)
		}
		msg = strings.ReplaceAll(msg, "\r", "")
		msg = strings.ReplaceAll(msg, "\n", "")

		if strings.ReplaceAll(msg, " ", "") == "" {
			continue
		}

		if strings.HasPrefix(msg, "!") {
			switch strings.Split(msg, " ")[0] {

			case "!exit":
				sendToEveryone(u.username+" exited the chat", info)
				fmt.Printf("\r%s exited the chat.\n>>> ", u.username)
				conn.Close()

			case "!user":
				u.username = strings.Split(msg, " ")[1]
				u.conn.Write([]byte("\r[INFO] User changed to " + u.username + "\n"))
				u.conn.Write([]byte("[" + u.username + "] "))
				continue

			case "!tell":
				to := strings.Split(msg, " ")[1]
				msg = msg[3+3+len(to):]
				tell(msg, to, u)
				continue

			case "!list":
				list(u)
				continue

			case "!help":
				conn.Write([]byte(help))
				continue

			default:
				conn.Write([]byte("Command not found. Use !help to see the command list."))
				continue
			}
		}
		sendToEveryone(msg, u)
		fmt.Printf("\r[%s] %s\n>>> ", u.username, msg)
		send = true

	}
}
