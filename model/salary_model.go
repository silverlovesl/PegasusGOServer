package model

import (
	"pegasus/dto"
	"pegasus/entity"
	"pegasus/utils"
)

// SalaryModel 初始化员工
type SalaryModel struct {
	dba utils.DataBaseAccessor
}

// NewSalaryModel 构造函数
func NewSalaryModel(dba utils.DataBaseAccessor) *SalaryModel {
	return &SalaryModel{dba: dba}
}

// FindCurrentMonthSalary 查找当月工资
func (m SalaryModel) FindCurrentMonthSalary(empID string, year int, month int) dto.SalaryDTO {

	var salaryDetails []entity.T_P_SalaryDetails
	var salaryItemDTOs []dto.SalaryItemDTO
	var salaryDetailDTOs []dto.SalaryDetailDTO

	empModel := EmployeeModel{dba: m.dba}

	salaryDTO := dto.SalaryDTO{}

	m.dba.Table("t_p_salary_details").
		Where(&entity.T_P_SalaryDetails{EmpID: empID, Year: year, Month: month}).
		Find(&salaryDetails)

	for i := 0; i < len(salaryDetails); i++ {

		m.dba.Table("t_m_salary_item").
			Where(&entity.T_M_SalaryItem{ItemCD: salaryDetails[i].ItemCD}).
			First(&salaryDetails[i].SalaryItem)

		salaryItemDTOs = append(salaryItemDTOs, dto.MappingSalaryItem(salaryDetails[i].SalaryItem))

		salaryDetailDTOs = append(salaryDetailDTOs, dto.MappingSalaryDetail(salaryDetails[i]))
	}

	salaryDTO.Details = salaryDetailDTOs
	salaryDTO.Employee = empModel.FindEmployeeByID(empID)

	return salaryDTO
}
