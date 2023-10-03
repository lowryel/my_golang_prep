##### Golang Studies

* Creating Go modules
- Modules in Golang are basically functions or objects that are referenced or imported in another function
- e.g. import "fmt" - fmt is an inbuilt module which contain all the print statements such as Print, Printf, Println, Sprintf, etc.

##### Let's create a module to Greet a User

* Greeting module:
- Create a folder called greetings and another hello
- In the greetings folder, run `go mod init demo/greetings`
- Create greetings.go file and add some code 
```go 
    package greetings
    import ("fmt")

    func Hello (name string) string {
        message := fmt.Println("Hello %v" name)
        return message
    } 
```
- Now in the hello folder, run `go mod init ex/hello` 
- Create a hello.go file and add the following code
```go
    package main
    import (
        "fmt"
        "ex/greetings" //module we created earlier
        // You can give it any name that resonate with you
    )
    func main() {
        message := greetings.Hello("Eugene")
        fmt.Println(message)
    }
``` 
- In the same hello folder, in the terminal, run `go mod edit -replace ex/greetings=../greetings`to replace ex/greetings with greetings
- Then run `go mod tidy` to synchronize the ex/hello module and add those dependencies that haven't been tracked
- Then run `go run .` in same hello folder to execute the program