package entity

import "time"

// T_P_SalaryDetails 工资明细
type T_P_SalaryDetails struct {
	EmpID      string
	Year       int
	Month      int
	ItemCD     string
	SalaryItem T_M_SalaryItem
	Amount     float64
	Remark     string
	InsRep     string
	InsDate    time.Time
	UpdRep     string
	UpdDate    time.Time
	DelFlg     string
	Version    int
}
