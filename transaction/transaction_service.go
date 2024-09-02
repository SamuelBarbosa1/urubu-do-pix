package transaction

import (
	"errors"
	"fmt"

	"urubu-do-pix/user"
)

func Transfer(fromPixKey, toPixKey string, amount float64) (*user.User, error) {
	fromUser, fromExists := user.GetUserByPixKey(fromPixKey)
	toUser, toExists := user.GetUserByPixKey(toPixKey)

	if !fromExists || !toExists {
		return nil, errors.New("uma das chaves Pix não existe")
	}

	if fromUser.Balance < amount {
		return nil, errors.New("saldo insuficiente")
	}

	fromUser.UpdateBalance(-amount, fmt.Sprintf("Transferido R$ %.2f para %s", amount, toUser.Name))
	toUser.UpdateBalance(amount, fmt.Sprintf("Recebido R$ %.2f de %s", amount, fromUser.Name))

	user.UpdateUser(fromUser)
	user.UpdateUser(toUser)

	return toUser, nil
}
