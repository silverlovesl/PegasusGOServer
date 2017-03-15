package model

import (
	"pegasus/dto"
	"pegasus/entity"
	"pegasus/utils"
	"strings"
)

// EmployeeModel 初始化员工
type EmployeeModel struct {
	dba utils.DataBaseAccessor
}

// NewEmployeeModel 构造函数
func NewEmployeeModel(dba utils.DataBaseAccessor) *EmployeeModel {
	return &EmployeeModel{dba: dba}
}

// FindEmployeeByID 根据员工ID查找
func (m EmployeeModel) FindEmployeeByID(empID string) dto.EmployeeDTO {

	var employee entity.T_M_Employee
	var result dto.EmployeeDTO

	// 获取员工信息
	m.dba.Table("t_m_employee").Where(&entity.T_M_Employee{EmpID: empID}).Find(&employee)

	// 如果找不到员工
	if strings.TrimSpace(employee.EmpID) == "" {
		return result
	}

	result = dto.MappingEmployee(employee)

	return result
}
