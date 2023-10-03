package loop

import (
	"fmt"
	// "net"
	// "net/http"
	"strings"
)

func Loops() string {
	for {
		fmt.Println("Type QUIT to stop!")
		input := ""
		fmt.Printf("Please enter a string: ")
		fmt.Scanln(&input)

		// host:= net.JoinHostPort("https://google.com", "8080")
		// fmt.Println(host)

		// https, err := http.Get("https://github.com/lowryel")
		// if err != nil{
		// 	return err.Error()
		// }
		// body := &https.Body
		// fmt.Println("Response .body\n", body)

		if strings.ToUpper(input) == "QUIT" {
			break
		}
	}
	var array [10] int
	array[5] = 50
	array[2] = 120
	// num := make(map[int]int)
	for k, v := range array {
		fmt.Println(k, v)
		if array[k] >= 40{
			fmt.Printf("This is the index of v greater than 40: %d \n",v)
		}
	}
	fmt.Println(array)
	return ""
}

