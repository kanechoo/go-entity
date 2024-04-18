package env

import (
	"log"
	"os"
)

func Get(key string) string {
	// 获取指定环境变量的值
	value := os.Getenv(key)
	// 如果环境变量存在，则打印其值
	if value != "" {
		log.Printf("环境变量 %s 的值为: %s", key, value)
	} else {
		// 如果环境变量不存在，则打印提示信息
		log.Printf("环境变量 %s 不存在", key)
	}
	return value
}
