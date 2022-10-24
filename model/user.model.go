package model

type UserModel struct {
	Username string
	Password string
	Email    string
}

type UserLoginModel struct {
	Username string
	Password string
}

type UserDetailModel struct {
	Uuid     uint
	Username string
	Password string
	Email    string
}
