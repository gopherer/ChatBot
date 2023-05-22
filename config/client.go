package config

type Client struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Mode     string `json:"mode"`
}
