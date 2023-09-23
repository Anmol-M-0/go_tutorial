package test_map

import (
	"fmt"

	mysql "github.com/go-sql-driver/mysql"
)

func TestMapper() {
	var m map[string]int = make(map[string]int)
	m["ubuntu"] = 1
	m["fedora"] = 2
	m["mint"] = 3
	for k, v := range m {
		fmt.Println("{}:{}", k, v)
	}
	fmt.Println(mysql.ErrBusyBuffer)
}
