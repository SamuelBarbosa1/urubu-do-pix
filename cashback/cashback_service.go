package cashback

import (
	"fmt"
	"math/rand"
	"time"
	"urubu-do-pix/user"
)

func ApplyRandomCashback(u *user.User) {
	rand.Seed(time.Now().UnixNano())
	cashback := rand.Float64() * 10 // Cashback aleatório de até 10 reais
	u.UpdateBalance(cashback, fmt.Sprintf("Cashback recebido: R$ %.2f", cashback))
}
