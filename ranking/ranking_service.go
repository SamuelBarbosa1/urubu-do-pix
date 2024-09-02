package ranking

import (
	"fmt"
	"sort"
	"urubu-do-pix/user"
)

func GetTopUsers(users map[string]*user.User) {
	type UserRank struct {
		Name    string
		Balance float64
	}
	var rankings []UserRank

	for _, u := range users {
		rankings = append(rankings, UserRank{Name: u.Name, Balance: u.Balance})
	}

	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].Balance > rankings[j].Balance
	})

	fmt.Println("Ranking de Usuários:")
	for i, rank := range rankings {
		fmt.Printf("%d. %s - R$ %.2f\n", i+1, rank.Name, rank.Balance)
	}
}
