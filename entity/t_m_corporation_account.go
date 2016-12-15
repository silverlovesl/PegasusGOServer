package entity

import "time"

type T_M_Corporation_Account struct {
	CompanyID          string
	BankCode           string
	BranchCode         string
	AccountCode        string
	AccountHolderName  string
	AccountHolderSpell *string
	AccountType        string
	DefaultUse         string
	InsRep             string
	InsDate            time.Time
	UpdRep             string
	UpdDate            time.Time
	DelFlg             string
	Version            int
}
