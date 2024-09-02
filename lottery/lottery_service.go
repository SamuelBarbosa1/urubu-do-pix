package lottery

import (
	"errors"
	"fmt"
	"urubu-do-pix/user"
)

type Lottery struct {
	Prize   float64
	Winner  string
	Message string
}

func ParticipateInLottery(pixKey string, betAmount float64) (Lottery, error) {
	participant, exists := user.GetUserByPixKey(pixKey)
	if !exists {
		return Lottery{}, errors.New("usuário não encontrado")
	}

	if participant.Balance < betAmount {
		return Lottery{}, errors.New("saldo insuficiente para participar da loteria")
	}

	participant.UpdateBalance(-betAmount, fmt.Sprintf("Participou da loteria com aposta de R$ %.2f", betAmount))
	user.UpdateUser(participant)

	lotteryResult := Lottery{
		Prize:   0,
		Winner:  "",
		Message: "O urubu comeu o seu prêmio! Você perdeu a aposta.",
	}

	return lotteryResult, nil
}
