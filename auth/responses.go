package auth

type AccountResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
