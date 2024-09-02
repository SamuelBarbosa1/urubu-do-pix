package troll

import (
	"fmt"
	"math/rand"
	"time"
	"urubu-do-pix/user"
)

func SurpriseTransfer(u *user.User, allUsers []*user.User) {
	rand.Seed(time.Now().UnixNano())
	amount := rand.Float64() * 20
	var randomUser *user.User

	for _, otherUser := range allUsers {
		if otherUser.PixKey != u.PixKey {
			randomUser = otherUser
			break
		}
	}

	if randomUser != nil {
		u.UpdateBalance(-amount, fmt.Sprintf("Surpresa! Transferiu R$ %.2f para %s", amount, randomUser.Name))
		randomUser.UpdateBalance(amount, fmt.Sprintf("Surpresa! Recebeu R$ %.2f de %s", amount, u.Name))
	}
}
