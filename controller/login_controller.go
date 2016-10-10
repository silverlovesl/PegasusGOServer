package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pegasus/utils"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	render "gopkg.in/unrolled/render.v1"
)

//LoginController Login控制器
type LoginController struct {
	dba      utils.DataBaseAccessor
	renderer *render.Render
	auth     utils.Authorization
}

type token struct {
	Data string `json:"token"`
}

// UserCredentials 用户认证
type userCredentials struct {
	Username string `json:"username"`
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
		fmt.Fprintf(w, "Error in request")
		return
	}

	fmt.Println(user.Username, user.Password)

	if user.Username != "admin" || user.Password != "123" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Failed to login")
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	signer := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	claims["jti"] = "1"
	claims["CustomUserInfo"] = struct {
		Name string
		Role string
	}{user.Username, "Member"}

	signer.Claims = claims

	tokenString, err := signer.SignedString(c.auth.SigningKey)

	//fmt.Println(string(signingKey))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
	}

	//create a token instance using the token string
	resp := token{tokenString}

	c.renderer.JSON(w, http.StatusOK, resp)
}
