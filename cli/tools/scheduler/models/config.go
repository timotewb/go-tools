package models

type Config struct {
	Email EmailConfig `json:"email"`
}

type EmailConfig struct {
	To string `json:"to"`
	From string `json:"from"`
	Host string `json:"host"`
	Port int `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}