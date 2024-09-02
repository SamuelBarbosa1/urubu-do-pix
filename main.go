package main

import (
	"fmt"
	"urubu-do-pix/achievements"
	"urubu-do-pix/cashback"
	"urubu-do-pix/charity"
	"urubu-do-pix/loan"
	"urubu-do-pix/lottery"
	"urubu-do-pix/ranking"
	"urubu-do-pix/transaction"
	"urubu-do-pix/troll"
	"urubu-do-pix/user"
)

func main() {
	for {
		fmt.Println("\nBem-vindo ao Urubu do Pix!")
		fmt.Println("1. Criar Conta")
		fmt.Println("2. Login")
		fmt.Println("3. Sair")
		fmt.Print("Escolha uma opção: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createAccount()
		case 2:
			login()
		case 3:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}

func createAccount() {
	var name, password string
	fmt.Print("Nome: ")
	fmt.Scanln(&name)
	fmt.Print("Senha: ")
	fmt.Scanln(&password)

	newUser, err := user.CreateUser(name, password)
	if err != nil {
		fmt.Println("Erro ao criar conta:", err)
		return
	}

	fmt.Printf("Conta criada com sucesso!\nSua chave Pix é: %s\n", newUser.PixKey)
}

func login() {
	var pixKey, password string
	fmt.Print("Chave Pix: ")
	fmt.Scanln(&pixKey)
	fmt.Print("Senha: ")
	fmt.Scanln(&password)

	loggedInUser, err := user.LoginUser(pixKey, password)
	if err != nil {
		fmt.Println("Erro ao fazer login:", err)
		return
	}

	fmt.Printf("\nBem-vindo, %s!\n", loggedInUser.Name)
	userMenu(loggedInUser)
}

func userMenu(loggedInUser *user.User) {
	for {
		fmt.Println("\n1. Ver Saldo")
		fmt.Println("2. Depositar Dinheiro")
		fmt.Println("3. Ver Histórico de Transações")
		fmt.Println("4. Transferir Dinheiro")
		fmt.Println("5. Participar da Loteria")
		fmt.Println("6. Receber Cashback Aleatório")
		fmt.Println("7. Ver Ranking de Usuários")
		fmt.Println("8. Ver Conquistas")
		fmt.Println("9. Modo Troll - Transferências Surpresa")
		fmt.Println("10. Solicitar Empréstimo")
		fmt.Println("11. Pagar Empréstimo")
		fmt.Println("12. Doar para Caridade")
		fmt.Println("13. Logout")
		fmt.Print("Escolha uma opção: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Printf("Seu saldo é: R$ %.2f\n", loggedInUser.Balance)
		case 2:
			var depositAmount float64
			fmt.Print("Digite o valor que deseja depositar: R$ ")
			fmt.Scanln(&depositAmount)
			loggedInUser = deposit(loggedInUser, depositAmount)
		case 3:
			fmt.Println("Histórico de Transações:")
			for _, transaction := range loggedInUser.History {
				fmt.Println(transaction)
			}
		case 4:
			transferMoney(loggedInUser)
		case 5:
			participateInLottery(loggedInUser)
		case 6:
			cashback.ApplyRandomCashback(loggedInUser)
		case 7:
			ranking.GetTopUsers(user.GetAllUsersMap())
		case 8:
			achievements.CheckAndUnlockAchievements(loggedInUser)
		case 9:
			troll.SurpriseTransfer(loggedInUser, user.GetAllUsers())
		case 10:
			var loanAmount float64
			fmt.Print("Digite o valor do empréstimo: R$ ")
			fmt.Scanln(&loanAmount)
			loan.RequestLoan(loggedInUser, loanAmount)
		case 11:
			var repayAmount float64
			fmt.Print("Digite o valor a ser pago: R$ ")
			fmt.Scanln(&repayAmount)
			loan.RepayLoan(loggedInUser, repayAmount)
		case 12:
			var donationAmount float64
			fmt.Print("Digite o valor da doação: R$ ")
			fmt.Scanln(&donationAmount)
			charity.DonateToCharity(loggedInUser, donationAmount)
		case 13:
			fmt.Println("Deslogando...")
			user.UpdateUser(loggedInUser)
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
func deposit(loggedInUser *user.User, amount float64) *user.User {
	loggedInUser.UpdateBalance(amount, fmt.Sprintf("Depósito de R$ %.2f", amount))
	fmt.Printf("Parabéns, você depositou R$ %.2f!\n", amount)
	return loggedInUser
}

func transferMoney(loggedInUser *user.User) {
	var toPixKey string
	var amount float64
	fmt.Print("Digite a chave Pix do destinatário: ")
	fmt.Scanln(&toPixKey)
	fmt.Print("Digite o valor que deseja transferir: R$ ")
	fmt.Scanln(&amount)

	_, err := transaction.Transfer(loggedInUser.PixKey, toPixKey, amount)
	if err != nil {
		fmt.Println("Erro na transferência:", err)
	} else {
		fmt.Println("Transferência realizada com sucesso!")
	}
}

func participateInLottery(loggedInUser *user.User) {
	var betAmount float64
	fmt.Print("Digite o valor da aposta: R$ ")
	fmt.Scanln(&betAmount)

	lotteryResult, err := lottery.ParticipateInLottery(loggedInUser.PixKey, betAmount)
	if err != nil {
		fmt.Println("Erro ao participar da loteria:", err)
		return
	}

	fmt.Println(lotteryResult.Message)
}
