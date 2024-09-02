package user

type User struct {
	ID          string
	Name        string
	Password    string
	PixKey      string
	Balance     float64
	History     []string
	DisplayName string
	Description string
}
