package commands

import socketio "github.com/googollee/go-socket.io"

func LogoutCommand(s socketio.Conn) string {
	username, exists := ActiveUsers[s.ID()]
	if exists {
		delete(ActiveUsers, s.ID())
		return username
	}
	return "User not found."
}
