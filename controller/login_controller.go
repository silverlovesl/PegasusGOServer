package controller

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pegasus/entity"
	"pegasus/utils"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	"pegasus/model"

	render "gopkg.in/unrolled/render.v1"
)

//LoginController Login控制器
type LoginController struct {
	dba      utils.DataBaseAccessor
	renderer *render.Render
	auth     utils.Authorization
}

type authReuslt struct {
	Token     string
	LoginUser *entity.T_M_Login
	Success   bool
	ErrorCode string
}

// UserCredentials 用户认证
type userCredentials struct {
	LoginID  string `json:"loginID"`
	Password string `json:"password"`
}

// NewLoginController LoginController 构造函数
func NewLoginController(auth utils.Authorization, dba utils.DataBaseAccessor, renderer *render.Render) *LoginController {
	return &LoginController{
		dba:      dba,
		renderer: renderer,
		auth:     auth,
	}
}

// Register 注册Router
func (c LoginController) Register(router *mux.Router) {
	router.HandleFunc("/api/login", c.doLogin).Methods(utils.MethodPOST)
}

func (c LoginController) doLogin(w http.ResponseWriter, r *http.Request) {

	var user userCredentials

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println(err)
		fmt.Fprintf(w, "Error in request")
		return
	}

	// 初始化 LoginModel
	loginModel := model.NewLoginModel(c.dba)
	bPwd := []byte(user.Password)
	cryptoPwd := md5.Sum(bPwd)
	cryptoStr := hex.EncodeToString(cryptoPwd[:])

	fmt.Println(user.LoginID, cryptoStr)

	// 执行登录操作
	loginUser := loginModel.Certificate(user.LoginID, cryptoStr)

	// Login失败
	if loginUser.LoginID == "" {
		fmt.Println("用户名和密码不存在")
		resp := authReuslt{LoginUser: nil, Token: "", Success: false, ErrorCode: "001"}
		c.renderer.JSON(w, http.StatusOK, resp)
		return
	}

	// 用户被锁住
	if loginUser.Islock == "1" {
		fmt.Println("被锁住的用户")
		resp := authReuslt{LoginUser: nil, Token: "", Success: false, ErrorCode: "002"}
		c.renderer.JSON(w, http.StatusOK, resp)
		return
	}

	signer := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["iss"] = loginUser.LoginID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	claims["jti"] = "1"
	claims["CustomUserInfo"] = struct {
		Name string
		Role string
	}{user.LoginID, "Member"}

	signer.Claims = claims

	tokenString, err := signer.SignedString(c.auth.SigningKey)

	fmt.Println(string(tokenString))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
	}

	//create a token instance using the token string
	resp := authReuslt{LoginUser: loginUser, Token: tokenString, Success: true, ErrorCode: ""}
	c.renderer.JSON(w, http.StatusOK, resp)
}
