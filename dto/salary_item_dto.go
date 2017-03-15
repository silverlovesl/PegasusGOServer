package dto

import "time"

// SalaryItemDTO
type SalaryItemDTO struct {
	ItemCD      string    `json:"itemCD"`
	InOut       string    `json:"inOut"`
	ItemName    string    `json:"itemName"`
	Description string    `json:"description"`
	TaxCalc0    string    `json:"taxCalc0"`
	TaxCalc1    string    `json:"taxCalc1"`
	TaxCalc2    string    `json:"taxCalc2"`
	TaxCalc3    string    `json:"taxCalc3"`
	TaxCalc4    string    `json:"taxCalc4"`
	TaxCalc5    string    `json:"taxCalc5"`
	TaxCalc6    string    `json:"taxCalc6"`
	TaxCalc7    string    `json:"taxCalc7"`
	TaxCalc8    string    `json:"taxCalc8"`
	TaxCalc9    string    `json:"taxCalc9"`
	InsRep      string    `json:"insRep"`
	InsDate     time.Time `json:"insDate"`
	UpdRep      string    `json:"updRep"`
	UpdDate     time.Time `json:"updDate"`
	DelFlg      string    `json:"delFlg"`
	Version     int       `json:"version"`
}
