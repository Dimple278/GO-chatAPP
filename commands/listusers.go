package commands

var ActiveUsers = make(map[string]string) // key: socketID, value: username

func ListUsersCommand() string {
	if len(ActiveUsers) == 0 {
		return "No active users."
	}

	message := "Active users:\n"
	for _, username := range ActiveUsers {
		message += "- " + username + "\n"
	}

	return message
}
