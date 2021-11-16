package apimodels

type LoginModelIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
