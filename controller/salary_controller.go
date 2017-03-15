package controller

import (
	"net/http"
	"pegasus/model"
	"pegasus/utils"
	"strconv"

	"github.com/gorilla/mux"

	render "gopkg.in/unrolled/render.v1"
)

//SalaryController 工资信息控制器
type SalaryController struct {
	dba      utils.DataBaseAccessor
	renderer *render.Render
	auth     utils.Authorization
}

// NewSalaryController SalaryController 构造函数
func NewSalaryController(auth utils.Authorization, dba utils.DataBaseAccessor, renderer *render.Render) *SalaryController {
	return &SalaryController{
		dba:      dba,
		renderer: renderer,
		auth:     auth,
	}
}

// Register 注册路由
func (c SalaryController) Register(router *mux.Router) {
	_oAuth := c.auth
	// router.Handle("/api/salary/currentMonthSalary/{empID}/{year}/{month}", _oAuth.UserActionWrap(c.findEmployeeCurrentMonthSalary)).Methods(utils.MethodGET)
	router.Handle("/api/salary/currentMonthSalary/{loginID}/{year}/{month}", _oAuth.UserActionWrap(c.findEmployeeCurrentMonthSalary)).Methods(utils.MethodGET)
}

func (c SalaryController) findEmployeeCurrentMonthSalary(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	params := mux.Vars(r)

	loginID := params["loginID"]

	year, errY := strconv.Atoi(params["year"])
	if errY != nil {
		panic(errY)
	}

	month, errM := strconv.Atoi(params["month"])
	if errM != nil {
		panic(errM)
	}

	salaryModel := model.NewSalaryModel(c.dba)
	loginModel := model.NewLoginModel(c.dba)

	loginUser := loginModel.FindLoginUserByID(loginID)

	salaryList := salaryModel.FindCurrentMonthSalary(loginUser.EmpID, year, month)

	c.renderer.JSON(w, http.StatusOK, salaryList)
}
