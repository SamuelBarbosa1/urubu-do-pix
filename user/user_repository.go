package user

var users = make(map[string]*User)

func GetUserByPixKey(pixKey string) (*User, bool) {
	user, exists := users[pixKey]
	return user, exists
}

func UpdateUser(user *User) {
	users[user.PixKey] = user
}
