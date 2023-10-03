package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)


func Hello (name string) (string, error) {
	if name == "" {
		return "", errors.New("name not provided")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

func Hellos (names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name]=message
	}
	fmt.Println(time.Now())
	return messages, nil
}

func randomFormat () string {
	format := []string {
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
	return format[rand.Intn(len(format))]
}


func Sum (a, b int) int {
	div := a/b
	if div ==0 {
		fmt.Println("Errored Out")
	}else{
		// fmt.Println(div)
		return div
	}
	return 0	
}

