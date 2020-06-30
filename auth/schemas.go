package auth

type LoginSchema struct {
	Identity string `json:"identity" form:"identity"`
	Password string `json:"password" form:"password"`
}
