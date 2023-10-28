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

###### Command to install go
`[ ! -d "/usr/local/go" ] && cd /tmp && wget https://go.dev/dl/go1.17.4.linux-amd64.tar.gz && tar -C /usr/local/ -xzf go1.17.4.linux-amd64.tar.gz && cd /usr/local/ && echo "export PATH=\$PATH:/usr/local/go/bin:\$HOME/go/bin" >> ~/.bashrc && echo "export GOROOT=/usr/local/go" >> ~/.bashrc && echo "export PATH=\$PATH:/usr/local/go/bin:\$HOME/go/bin" >> /home/*/.bashrc && echo "export GOROOT=/usr/local/go" >> /home/*/.bashrc && source ~/.bashrc && source /home/*/.bashrc`

###### Twilio SMS service

`go
    package validation

    import (
        "fmt"
        "locus-api/pkg/logger"
        "os"
        "regexp"

        // "strconv"

        "github.com/twilio/twilio-go"
        openapi "github.com/twilio/twilio-go/rest/verify/v2"
        "golang.org/x/exp/rand"
    )

    // IsValidPIN validates a phone number
    func IsValidPIN(pin string) bool {
        // Define a regular expression to match phone numbers.
        regex := `^\d{4}$`

        // Compile the regular expression.
        re := regexp.MustCompile(regex)

        // Test if the phone number matches the regular expression.
        return re.MatchString(pin)
    }


    var TWILIO_ACCOUNT_SID string = os.Getenv("TWILIO_ACCOUNT_SID")
    var TWILIO_AUTH_TOKEN string = os.Getenv("TWILIO_AUTH_TOKEN")
    var TWILIO_SERVICES_ID string = os.Getenv("TWILIO_SERVICES_ID")
    var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: TWILIO_ACCOUNT_SID,
        Password: TWILIO_AUTH_TOKEN,
    })

    // send OTP to user for phone verification
    func SendOtp(to string) string {
    params := &openapi.CreateVerificationParams{}
    params.SetTo(to)
    params.SetChannel("sms")
    //    params.SetCustomCode()


    resp, err := client.VerifyV2.CreateVerification(TWILIO_SERVICES_ID, params)

    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Printf("Sent verification '%s'\n", *resp.Sid)
    }
    return *resp.Sid
    }

    var letterRunes = []rune("0123456789")

    func RandStringRunes(n int) string {
        b := make([]rune, n)
        for i := range b {
            b[i] = letterRunes[rand.Intn(len(letterRunes))]
        }
        return string(b)
    }

    // verify OTP with phone
    func CheckOtp(to string) string {
    code := RandStringRunes(4)
    //    fmt.Println("Please check your phone and enter the code:")
    //    fmt.Scanln(&code)
            logger.DevLog(code)
    //    params := &openapi.CreateVerificationCheckParams{}
    //    params.SetTo(to)
    //    params.SetCode(code)
            logger.DevLog(to)
    //    resp, err := client.VerifyV2.CreateVerificationCheck(TWILIO_SERVICES_ID, params)

    //    if err != nil {
    //        fmt.Println(err.Error())
    //    } else if *resp.Status == "approved" {
    //        fmt.Println("Correct!")
    //    } else {
    //        logger.DevLog("Incorrect pin!")
    //        panic("Invalid pin")
    //    }
        return string(code)
    }
`

