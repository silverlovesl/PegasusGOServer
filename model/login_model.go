package model

import (
	"fmt"
	"pegasus/entity"
	"pegasus/utils"
)

// LoginModel 初始化员工
type LoginModel struct {
	dba utils.DataBaseAccessor
}

// NewLoginModel 构造函数
func NewLoginModel(dba utils.DataBaseAccessor) *LoginModel {
	return &LoginModel{dba: dba}
}

// Certificate Do login if success return true else false
func (m LoginModel) Certificate(loginID string, password string) *entity.T_M_Login {

	loginUser := entity.T_M_Login{}

	fmt.Println("★★★")
	fmt.Println(loginUser)

	m.dba.Table("t_m_login").Where(" login_id = ? and password =? ", loginID, password).First(&loginUser)

	fmt.Println(loginUser)

	return &loginUser
}
