package server

import (
	"fmt"
	"net/http"
	"pegasus/controller"
	"pegasus/utils"

	render "gopkg.in/unrolled/render.v1"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// Server 服务器
type Server struct {
	*negroni.Negroni
}

// InitServer 初始化服务器
func InitServer() *Server {

	fmt.Println("initServer...")

	// 创建服务对象
	s := Server{negroni.Classic()}

	// 创建路由
	router := mux.NewRouter()

	// Json render
	renderer := render.New()

	// 数据库访问对象
	dba := utils.NewDataBaseAccessor()

	// WebService 认证
	oAuth2 := utils.NewAuthorization()

	// 初始化Key
	oAuth2.InitKeys()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pegasus Web Service is working fine !!!")
	})

	controller.NewLoginController(*oAuth2, *dba, renderer).Register(router)
	// 初始化员工控制器,路由注册
	controller.NewEmployeeController(*oAuth2, *dba, renderer).Register(router)

	//添加路由
	s.UseHandler(router)

	return &s
}
