package entity

import "time"

// T_M_Employee 员工实体
type T_M_Employee struct {
	EmpID              string
	LastName           string
	FirstName          string
	LastNameKana       string
	FirstNameKana      string
	Sex                string
	Birthday           time.Time
	Mobile             *string
	Zip                *string
	ComEmail           *string
	PersonalEmail      *string
	Address1           *string
	Address2           *string
	BankCode           *string
	BranchCode         *string
	AccountCode        *string
	AccountHolderName  *string
	AccountHolderSpell *string
	ResidenceNumber    *string
	PeriodOfStay       *int
	ResidenceTerm      *time.Time
	PassportNumber     *string
	MyNumber           string
	Photo              *string
	HireDate           *time.Time
	Resigned           *string
	Status             *string
	ResignedDate       *time.Time
	EmpContractType    string
	InsRep             string
	InsDate            time.Time
	UpdRep             string
	UpdDate            time.Time
	DelFlg             string
	Version            int
	PopSupporting      []T_P_PopulationSupporting
}
