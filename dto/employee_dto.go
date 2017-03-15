package dto

import "time"

// EmployeeDTO 员工DTO
type EmployeeDTO struct {
	EmpID              string     `json:"empID"`
	LastName           string     `json:"lastName"`
	FirstName          string     `json:"firstName"`
	LastNameKana       string     `json:"lastNameKana"`
	FirstNameKana      string     `json:"firstNameKana"`
	Sex                string     `json:"sex"`
	Birthday           time.Time  `json:"birthday"`
	Mobile             *string    `json:"mobile"`
	Zip                *string    `json:"zip"`
	ComEmail           *string    `json:"comEmail"`
	PersonalEmail      *string    `json:"personalEmail"`
	Address1           *string    `json:"address1"`
	Address2           *string    `json:"address2"`
	BankCode           *string    `json:"bankCode"`
	BranchCode         *string    `json:"branchCode"`
	AccountCode        *string    `json:"accoutCode"`
	AccountHolderName  *string    `json:"accountHolderName"`
	AccountHolderSpell *string    `json:"accountHolderSpell"`
	ResidenceNumber    *string    `json:"residenceNumber"`
	PeriodOfStay       *int       `json:"periodOfStay"`
	ResidenceTerm      *time.Time `json:"residenceTerm"`
	PassportNumber     *string    `json:"passportNumber"`
	MyNumber           string     `json:"myNumber"`
	Photo              *string    `json:"photo"`
	HireDate           *time.Time `json:"hireDate"`
	Resigned           *string    `json:"resigned"`
	Status             *string    `json:"status"`
	ResignedDate       *time.Time `json:"resignedDate"`
	EmpContractType    string     `json:"empContractType"`
	InsRep             string     `json:"insRep"`
	InsDate            time.Time  `json:"insDate"`
	UpdRep             string     `json:"updRep"`
	UpdDate            time.Time  `json:"updDate"`
	DelFlg             string     `json:"delFlg"`
	Version            int        `json:"version"`
}
