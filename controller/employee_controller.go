package controller

import (
	"fmt"
	"net/http"
	"pegasus/model"
	"pegasus/utils"

	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
)

// EmployeeController 员工控制器
type EmployeeController struct {
	dba      utils.DataBaseAccessor
	renderer *render.Render
	auth     utils.Authorization
}

// NewEmployeeController EmployeeController 构造函数
func NewEmployeeController(auth utils.Authorization, dba utils.DataBaseAccessor, renderer *render.Render) *EmployeeController {
	return &EmployeeController{
		dba:      dba,
		renderer: renderer,
		auth:     auth,
	}
}

// Register 注册路由
func (c EmployeeController) Register(router *mux.Router) {
	//_oAuth2 := c.auth
	//router.Handle("/api/employee/findEmployeeById/{empID}", _oAuth2.UserActionWrap(c.findEmployeeByID)).Methods(utils.MethodGET)
	router.HandleFunc("/api/employee/findEmployeeById/{empID}", c.findEmployeeByID).Methods(utils.MethodGET)
}

// FindEmployeeByID 根据ID获取制定的员工
//func (c EmployeeController) findEmployeeByID(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
func (c EmployeeController) findEmployeeByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	empID := params["empID"]

	fmt.Println("empID: " + empID)

	empModel := model.NewEmployeeModel(c.dba)

	employee := empModel.FindEmployeeByID(empID)

	// if err != nil {
	// 	c.renderer.JSON(w, http.StatusInternalServerError, employee)
	// }

	c.renderer.JSON(w, http.StatusOK, employee)
}
