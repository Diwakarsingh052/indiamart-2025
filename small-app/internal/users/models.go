package users

type User struct {
	Id           string
	Email        string `json:"email"`
	Name         string `json:"name"` // giving the name of the field in the json ouptut
	Age          int64  `json:"age"`
	PasswordHash string `json:"password_hash"`
}
