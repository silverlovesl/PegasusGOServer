package dto

// SalaryDTO 工资DTO
type SalaryDTO struct {
	Employee EmployeeDTO       `json:"employee"`
	Details  []SalaryDetailDTO `json:"details"`
}
