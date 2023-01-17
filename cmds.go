package main

import "fmt"

// help messages
const help = `Available commands:
!list - List all users connected.
!help - Show this message.
!tell <user> <msg> - Tell another user something.
!user <new_username> - Change your username.
`

const help_admin = `Available commands:
!list - List all users connected
!help - Show this message
!tell <user> <msg> - Tell another user something
!ban <user> - Bans an user's IP
!ip <user> - Shows an user's IP
!stop - Stops the server
`

// everyone commands

func list(sender User) {
	sender.conn.Write([]byte("Users connected: \n"))
	for _, user := range users {
		sender.conn.Write([]byte(user.username + "\n"))
	}
}

func tell(msg, to string, sender User) { // Sends a private message to another user
	sender.conn.Write([]byte("[to " + to + "]" + msg + "\n"))
	user_idx := find_user(to, users)
	if user_idx != -1 {
		u := users[user_idx]
		u.conn.Write([]byte("\n" + "[from " + sender.username + "]" + msg + "\n"))
		u.conn.Write([]byte("[" + u.username + "] "))
		return
	} else {
		sender.conn.Write([]byte("Could not find user."))
	}

}

// admin commands
func show_ip(user string) string { // Shows an user's IP
	user_idx := find_user(user, users)
	if user_idx != -1 {
		return users[user_idx].address
	} else {
		return "Could not find user."
	}
}

func ban(banned string) { // Bans an user's IP
	user_idx := find_user(banned, users)
	if user_idx != -1 {
		users[user_idx].conn.Write([]byte("\nYou have been banned."))
		banned_users = append(banned_users, *users[user_idx])
		users = remove_user(users, users[user_idx])
		users[user_idx].conn.Close()
	} else {
		fmt.Printf("Could not find user.")

	}

	sendToEveryone(banned+"has been banned.", info)
}
