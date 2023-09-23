package main

import (
	"fmt"
	"os"
	"test/test_map"
)

func main() {
	test_file()
	test_map_fn()
	test_map.TestMapper()
}

func test_file() {
	// fmt.Println("hello, world")
	var arr [24]int32
	var i, length int32
	length = int32(len(arr))
	for i = 0; i < length; i++ {
		arr[i] = i
	}
	var buffer []byte = make([]byte, 50)
	file_ptr, err := os.Open("test file.txt")
	if err != nil {
		fmt.Println("An Error Occured: {}", err)
		return
	}
	defer file_ptr.Close()
	var n int = -1
	// var file_err error
	for n != 0 {
		n, _ = file_ptr.Read(buffer)
		// if file_err != nil {
		// 	fmt.Println(" => An Error Occured: {}", file_err)
		// }
		fmt.Print(string(buffer[:n]))
	}
}

func test_map_fn() {

	var m map[string]int = make(map[string]int)
	m["ubuntu"] = 1
	m["fedora"] = 2
	m["mint"] = 3
	for k, v := range m {
		fmt.Printf("%s:%d\n", k, v)
	}

}
