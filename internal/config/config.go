package config

import (
	"encoding/json"
	"os"
)

// 配置结构体，用于保存从配置文件加载的配置信息
type config struct {
	Static              string    `json:"static"`              // 静态文件目录路径
	Template            string    `json:"template"`            // 模板文件目录路径
	Address             string    `json:"address"`             // 服务器监听地址
	Port                int       `json:"port"`                // 服务器监听端口
	HandleTimeoutSecond int       `json:"handleTimeoutSecond"` // 处理超时时间（秒）
	Log                 logConfig `json:"log"`                 // 日志配置
}

// 日志配置结构体
type logConfig struct {
	Trace   string `json:"trace"`   // 跟踪日志文件路径
	Info    string `json:"info"`    // 信息日志文件路径
	Warning string `json:"warning"` // 警告日志文件路径
	Error   string `json:"error"`   // 错误日志文件路径
}

// 全局变量，保存加载的配置
var Config config

// 加载配置文件的函数
func LoadConfig(configFile string) error {
	configData, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(configData, &Config)
}

// 测试用方法，用于导出默认配置到文件
func exportDefaultConfig(configFile string) error {
	defaultConfig := config{}
	configData, err := json.MarshalIndent(&defaultConfig, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile(configFile, configData, 0666)
	if err != nil {
		return err
	}
	return nil
}
