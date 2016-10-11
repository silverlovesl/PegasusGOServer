package model

import (
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
func (empM EmployeeModel) FindEmployeeByID(empID string) *entity.T_M_Employee {

	var pops []entity.T_P_PopulationSupporting
	var employee entity.T_M_Employee

	// 获取员工信息
	empM.dba.Table("t_m_employee").Where(&entity.T_M_Employee{EmpID: empID}).Find(&employee)

	// 如果找不到员工
	if strings.TrimSpace(employee.EmpID) == "" {
		return nil
	}

	// 获取抚养人口信息
	empM.dba.Where(&entity.T_P_PopulationSupporting{EmpID: empID}).Find(&pops)

	employee.PopSupporting = pops

	return &employee
}
