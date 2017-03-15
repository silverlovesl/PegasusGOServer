package dto

import "time"

// SalaryDetailDTO 工资明细
type SalaryDetailDTO struct {
	EmpID      string        `json:"empID"`
	Year       int           `json:"year"`
	Month      int           `json:"month"`
	SalaryItem SalaryItemDTO `json:"salaryItem"`
	Amount     float64       `json:"amount"`
	Remark     string        `json:"remark"`
	InsRep     string        `json:"insRep"`
	InsDate    time.Time     `json:"insDate"`
	UpdRep     string        `json:"updRep"`
	UpdDate    time.Time     `json:"updDate"`
	DelFlg     string        `json:"delFlg"`
	Version    int           `json:"version"`
}
