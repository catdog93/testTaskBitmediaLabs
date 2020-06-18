package main

type UserAuth struct {
	Login string `json:"login" binding:"required"`
	Pass  string `json:"pass" binding:"required"`
}

type TokenStruct struct {
	Token string `json:"token" binding:"required"`
}
