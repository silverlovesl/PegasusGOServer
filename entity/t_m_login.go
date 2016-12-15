package entity

import "time"

type T_M_Login struct {
	LoginID        string
	EmpID          string
	Password       string
	Islock         string
	Authority      string
	FirstTimeLogin string
	InsRep         string
	InsDate        time.Time
	UpdRep         string
	UpdDate        time.Time
	DelFlg         string
	Version        int
}
