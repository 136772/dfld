package model

type AdminUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	LoginDate int    `json:"logindate"`
	Admin     string `json:"admin"`
}
