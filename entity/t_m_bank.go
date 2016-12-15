package entity

import "time"

type T_M_Bank struct {
	BankCode     string
	BankName     string
	BankNamekana string
	InsRep       string
	InsDate      time.Time
	UpdRep       string
	UpdDate      time.Time
	DelFlg       string
	Version      int
}
