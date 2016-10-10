package model

import (
	"fmt"
	"pegasus/utils"
	"time"
)

// Employee 员工实体
type Employee struct {
	EmpID              string     `json:"empId"`
	LastName           string     `json:"lastName"`
	FirstName          string     `json:"firstName"`
	LastNameKana       string     `json:"lastNameKana"`
	FirstNameKana      string     `json:"firstNameKana"`
	Sex                string     `json:"sex"`
	Birthday           time.Time  `json:"birthday"`
	Mobile             *string    `json:"Mobile"`
	Zip                *string    `json:"zip"`
	ComEmail           *string    `json:"comEmail"`
	PersonalEmail      *string    `json:"PersonalEmail"`
	Address1           string     `json:"address1"`
	Address2           *string    `json:"address2"`
	BankCode           *string    `json:"bankCode"`
	BranchCode         *string    `json:"branchCode"`
	AccountCode        *string    `json:"accountCode"`
	AccountHolderName  *string    `json:"accountHolderName"`
	AccountHolderSpell *string    `json:"accountHolderSpell"`
	ResidenceNumber    *string    `json:"residenceNumber"`
	PeriodOfStay       *int       `json:"periodOfStay"`
	ResidenceTerm      *time.Time `json:"residenceTerm"`
	PassportNumber     *string    `json:"passportNumber"`
	MyNumber           *string    `json:"myNumber"`
	Photo              *string    `json:"photo"`
	HireDate           *time.Time `json:"hireDate"`
	Resigned           *string    `json:"resigned"`
	Status             *string    `json:"status"`
	ResignedDate       *time.Time `json:"resignedDate"`
	EmpContractType    string     `json:"empContractType"`
	ComField           BaseEntity `json:"comField"`
}

// EmployeeModel 初始化员工
type EmployeeModel struct {
	dba utils.DataBaseAccessor
}

// NewEmployeeModel 构造函数
func NewEmployeeModel(dba utils.DataBaseAccessor) *EmployeeModel {
	return &EmployeeModel{dba: dba}
}

// FindEmployeeByID 根据员工ID查找
func (empM EmployeeModel) FindEmployeeByID(empID string) (*Employee, error) {

	sqlBuf := `SELECT
	     t1.emp_id              
		,t1.last_name           
        ,t1.first_name          
        ,t1.last_name_kana      
        ,t1.first_name_kana    
		,t1.sex                
        ,t1.birthday            
        ,t1.mobile             
        ,t1.zip                 
        ,t1.com_email          
        ,t1.personal_email      
        ,t1.address1           
        ,t1.address2            
        ,t1.bank_code          
        ,t1.branch_code         
        ,t1.account_code
		,t1.account_holder_name 
        ,t1.account_holder_spell
        ,t1.residence_number  
        ,t1.period_of_stay   
        ,t1.residence_term    
        ,t1.passport_number
		,t1.my_number         
        ,t1.photo            
        ,t1.hire_date         
        ,t1.resigned         
        ,t1.status            
        ,t1.resigned_date    
        ,t1.emp_contract_type 
		,t1.ins_rep     
		,t1.ins_date          
        ,t1.upd_rep          
        ,t1.upd_date          
        ,t1.del_flg          
        ,t1.version                
	 FROM t_m_employee AS t1
	WHERE t1.emp_id = $1 `

	stmt, err := empM.dba.Prepare(sqlBuf)

	if err != nil {
		return nil, err
	}

	// 执行SQL
	resultRows, err := stmt.Query(empID)

	result := Employee{ComField: BaseEntity{}}

	for resultRows.Next() {
		resultRows.Scan(&result.EmpID,
			&result.LastName,
			&result.FirstName,
			&result.LastNameKana,
			&result.FirstNameKana,
			&result.Sex,
			&result.Birthday,
			&result.Mobile,
			&result.Zip,
			&result.ComEmail,
			&result.PersonalEmail,
			&result.Address1,
			&result.Address2,
			&result.BankCode,
			&result.BranchCode,
			&result.AccountCode,
			&result.AccountHolderName,
			&result.AccountHolderSpell,
			&result.ResidenceNumber,
			&result.PeriodOfStay,
			&result.ResidenceTerm,
			&result.PassportNumber,
			&result.MyNumber,
			&result.Photo,
			&result.HireDate,
			&result.Resigned,
			&result.Status,
			&result.ResignedDate,
			&result.EmpContractType,
			&result.ComField.InsRep,
			&result.ComField.InsDate,
			&result.ComField.UpdRep,
			&result.ComField.UpdDate,
			&result.ComField.DelFlg,
			&result.ComField.Version)
	}

	fmt.Println(result)

	return &result, nil
}
