package charity

import (
	"fmt"
	"urubu-do-pix/user"
)

func DonateToCharity(u *user.User, amount float64) {
	if u.Balance >= amount {
		u.UpdateBalance(-amount, fmt.Sprintf("Doou R$ %.2f para a caridade. Que generoso!", amount))
	} else {
		fmt.Println("Saldo insuficiente para fazer a doação!")
	}
}
