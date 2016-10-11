package entity

import "time"

// T_P_PopulationSupporting 抚养人口
type T_P_PopulationSupporting struct {
	EmpID         string
	Relationship  string
	LastName      *string
	FirstName     *string
	Birthday      *time.Time
	RsDescription *string
	Effective     string
	InsRep        string
	InsDate       time.Time
	UpdRep        string
	UpdDate       time.Time
	DelFlg        string
	Version       int
}
