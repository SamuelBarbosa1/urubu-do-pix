package user

var users = make(map[string]*User)

// Função para obter um usuário pelo PixKey
func GetUserByPixKey(pixKey string) (*User, bool) {
	user, exists := users[pixKey]
	return user, exists
}

// Função para atualizar um usuário
func UpdateUser(user *User) {
	users[user.PixKey] = user
}

// Função para obter todos os usuários
func GetAllUsers() []*User {
	allUsers := []*User{}
	for _, user := range users {
		allUsers = append(allUsers, user)
	}
	return allUsers
}

// Função para obter todos os usuários como um mapa
func GetAllUsersMap() map[string]*User {
	return users
}
