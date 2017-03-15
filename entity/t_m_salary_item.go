package entity

import "time"

// T_M_SalaryItem 工资项
type T_M_SalaryItem struct {
	ItemCD      string
	InOut       string
	ItemName    string
	Description string
	TaxCalc0    string
	TaxCalc1    string
	TaxCalc2    string
	TaxCalc3    string
	TaxCalc4    string
	TaxCalc5    string
	TaxCalc6    string
	TaxCalc7    string
	TaxCalc8    string
	TaxCalc9    string
	InsRep      string
	InsDate     time.Time
	UpdRep      string
	UpdDate     time.Time
	DelFlg      string
	Version     int
}
