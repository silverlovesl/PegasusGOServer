package utils

import "os"

const (
	// PORT 端口
	PORT string = "8088"
	// MethodGET  Http Method GET
	MethodGET string = "GET"
	// MethodPOST  Http Method POST
	MethodPOST string = "POST"
)

// SysConfig 系统环境变量
type SysConfig struct {
	IsDevelopment bool
}

// GetSystemConfig 系统配置
func GetSystemConfig() *SysConfig {
	return &SysConfig{
		IsDevelopment: os.Getenv("ENVIRONMENT") == "DEVELOPMENT",
	}
}
