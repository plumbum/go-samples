package data

//go:generate ffjson data.go

type Item struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"pwd"`
}

type Users struct {
	Description string `json:"description"`
	Users []Item `json:"users"`
}

