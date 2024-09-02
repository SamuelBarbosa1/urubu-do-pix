package loan

import (
	"fmt"
	"math"
	"urubu-do-pix/user"
)

const interestRate = 0.5 // 50% de juros

func RequestLoan(u *user.User, amount float64) {
	loanAmount := amount * (1 + interestRate)
	u.UpdateBalance(amount, fmt.Sprintf("Empréstimo recebido: R$ %.2f, você deve pagar R$ %.2f", amount, loanAmount))
}

func RepayLoan(u *user.User, loanAmount float64) {
	payment := loanAmount * math.Pow(1+interestRate, 1) // Aplicação do juro
	if u.Balance >= payment {
		u.UpdateBalance(-payment, fmt.Sprintf("Empréstimo pago: R$ %.2f", payment))
	} else {
		fmt.Println("Saldo insuficiente para pagar o empréstimo!")
	}
}
