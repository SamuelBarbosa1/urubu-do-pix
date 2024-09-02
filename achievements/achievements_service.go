package achievements

import (
	"fmt"
	"urubu-do-pix/user"
)

type Achievement struct {
	Title       string
	Description string
}

func CheckAndUnlockAchievements(u *user.User) []Achievement {
	var unlocked []Achievement

	// Exemplo de conquistas
	if u.Balance >= 1000 {
		unlocked = append(unlocked, Achievement{
			Title:       "Milionário",
			Description: "Você alcançou um saldo de R$ 1.000,00!",
		})
	}

	for _, ach := range unlocked {
		fmt.Printf("Conquista desbloqueada: %s - %s\n", ach.Title, ach.Description)
	}

	return unlocked
}
