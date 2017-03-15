package dto

import "pegasus/entity"

// MappingEmployee 员工Mapping
func MappingEmployee(entity entity.T_M_Employee) EmployeeDTO {

	employeeDTO := EmployeeDTO{
		EmpID:              entity.EmpID,
		LastName:           entity.LastName,
		FirstName:          entity.FirstName,
		LastNameKana:       entity.LastNameKana,
		FirstNameKana:      entity.FirstNameKana,
		Sex:                entity.Sex,
		Birthday:           entity.Birthday,
		Mobile:             entity.Mobile,
		Zip:                entity.Zip,
		ComEmail:           entity.ComEmail,
		PersonalEmail:      entity.PersonalEmail,
		Address1:           entity.Address1,
		Address2:           entity.Address2,
		BankCode:           entity.BankCode,
		BranchCode:         entity.BranchCode,
		AccountCode:        entity.AccountCode,
		AccountHolderName:  entity.AccountHolderName,
		AccountHolderSpell: entity.AccountHolderSpell,
		ResidenceNumber:    entity.ResidenceNumber,
		PeriodOfStay:       entity.PeriodOfStay,
		ResidenceTerm:      entity.ResidenceTerm,
		PassportNumber:     entity.PassportNumber,
		MyNumber:           entity.MyNumber,
		Photo:              entity.Photo,
		HireDate:           entity.HireDate,
		Resigned:           entity.Resigned,
		Status:             entity.Status,
		ResignedDate:       entity.ResignedDate,
		EmpContractType:    entity.EmpContractType,
		InsRep:             entity.InsRep,
		InsDate:            entity.InsDate,
		UpdRep:             entity.UpdRep,
		UpdDate:            entity.UpdDate,
		DelFlg:             entity.DelFlg,
		Version:            entity.Version,
	}

	return employeeDTO
}

// MappingSalaryDetail 工资明细Mapping
func MappingSalaryDetail(entity entity.T_P_SalaryDetails) SalaryDetailDTO {

	salaryDetailDTO := SalaryDetailDTO{
		EmpID:   entity.EmpID,
		Year:    entity.Year,
		Month:   entity.Month,
		Amount:  entity.Amount,
		Remark:  entity.Remark,
		InsRep:  entity.InsRep,
		InsDate: entity.InsDate,
		UpdRep:  entity.UpdRep,
		UpdDate: entity.UpdDate,
		DelFlg:  entity.DelFlg,
		Version: entity.Version,
	}

	return salaryDetailDTO
}

// MappingSalaryItem 工资项Mapping
func MappingSalaryItem(entity entity.T_M_SalaryItem) SalaryItemDTO {

	salaryItemDTO := SalaryItemDTO{
		ItemCD:      entity.ItemCD,
		InOut:       entity.InOut,
		ItemName:    entity.ItemName,
		Description: entity.Description,
		TaxCalc0:    entity.TaxCalc0,
		TaxCalc1:    entity.TaxCalc1,
		TaxCalc2:    entity.TaxCalc2,
		TaxCalc3:    entity.TaxCalc3,
		TaxCalc4:    entity.TaxCalc4,
		TaxCalc5:    entity.TaxCalc5,
		TaxCalc6:    entity.TaxCalc6,
		TaxCalc7:    entity.TaxCalc7,
		TaxCalc8:    entity.TaxCalc8,
		TaxCalc9:    entity.TaxCalc9,
	}

	return salaryItemDTO
}
