package env

import (
	"os"
)

func Get(key string) string {
	// 获取指定环境变量的值
	return os.Getenv(key)
}
