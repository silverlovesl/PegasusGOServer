package entity

import "time"

type T_M_Account_Subject struct {
	AccountSubjectID string
	PaymentMethod    string
	SubjectName      string
	SortBy           int
	TaxCode          string
	DebitCredit      string
	PlCalc           string
	Category         string
	Remark           *string
	InsRep           string
	InsDate          time.Time
	UpdRep           string
	UpdDate          time.Time
	DelFlg           string
	Version          int
}
