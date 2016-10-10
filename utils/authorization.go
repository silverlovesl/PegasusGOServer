package utils

import (
	"bytes"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
)

// Authorization 安全认证
type Authorization struct {
	SigningKey      []byte
	verificationKey []byte
}

// ActionHandler 简化请求内容
type ActionHandler func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

// NewAuthorization 构造函数
func NewAuthorization() *Authorization {
	return &Authorization{}
}

// InitKeys 初始化私有-公开密钥
func (auth Authorization) InitKeys() {
	var (
		err         error
		privKey     *rsa.PrivateKey
		pubKey      *rsa.PublicKey
		pubKeyBytes []byte
	)

	privKey, err = rsa.GenerateKey(cryptorand.Reader, 2048)
	if err != nil {
		log.Fatal("Error generating private key")
	}
	pubKey = &privKey.PublicKey

	// Create signingKey from privKey
	// prepare PEM block
	var privPEMBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey), // serialize private key bytes
	}
	// serialize pem
	privKeyPEMBuffer := new(bytes.Buffer)
	pem.Encode(privKeyPEMBuffer, privPEMBlock)
	//done
	auth.SigningKey = privKeyPEMBuffer.Bytes()

	//fmt.Println("private key: " + string(auth.SigningKey))

	// create verificationKey from pubKey. Also in PEM-format
	pubKeyBytes, err = x509.MarshalPKIXPublicKey(pubKey) //serialize key bytes
	if err != nil {
		// heh, fatality
		log.Fatal("Error marshalling public key")
	}

	var pubPEMBlock = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	// serialize pem
	pubKeyPEMBuffer := new(bytes.Buffer)
	pem.Encode(pubKeyPEMBuffer, pubPEMBlock)
	// done
	auth.verificationKey = pubKeyPEMBuffer.Bytes()

	//fmt.Println("public key: " + string(auth.verificationKey))
}

// ValidateTokenMiddleware 验证Token
func (auth Authorization) ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	tokenString := r.Header.Get("Authorization")

	fmt.Println("tokenString" + tokenString)

	// Return a Token using the cookie
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return auth.verificationKey, nil
	})

	if err == nil && token.Valid {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "未认证的请求")
	}
}

// UserActionWrap 用户请求Router包装方法,简化设置操作
func (auth Authorization) UserActionWrap(action ActionHandler) http.Handler {
	return negroni.New(negroni.HandlerFunc(auth.ValidateTokenMiddleware),
		negroni.HandlerFunc(action))
}
