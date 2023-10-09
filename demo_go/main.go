package main

import (
	"fmt"
	"time"
    "strconv"
	// "golang.org/x/text/cases"
)

// func main() {
//     fmt.Println("Hello world")
//     fmt.Printf("Please enter your age: ")
//     var age int;
//     fmt.Scanf("%d", &age)
//     const height int = 15
//     if age < height {
//         fmt.Printf("%d\n", age)
//         fmt.Printf("Height: %d is too short\n", height)
//         fmt.Println("Try another time")
//         hello()
//     } else {
//         fmt.Printf("%d Not in range\n", age)
//         hello()
//         runc()
//      }
// }

// func hello () {
//     fmt.Printf("Hello there!!!\n")
// }

// func runc()  {
//     count :=10
//     for i := 0; i < count; i++ {
//         fmt.Printf("%d\n", i)
//     }
//  }

// func (file *File) Write(b []byte) (n int, err error)

// func trace(s string) string {
//     fmt.Println("entering:", s)
//     return s
// }

// func un(s string) {
//     fmt.Println("leaving:", s)
// }

// func a() {
//     defer un(trace("a"))
//     fmt.Println("in a")
// }

// type SyncedBuffer struct {
//     name string
//     value int
// }

// func b() {
//     var v SyncedBuffer
//     defer un(trace("b"))
//     fmt.Println("in b")
//     a()
//     v.name = "Eugene"
//     v.value = 18
//     fmt.Printf("%s %d\n", v.name, v.value)
// }

func main() {
    // b()
    fmt.Print(nil)
    // Announce("hey there, this a goroutine", 2)
    fmt.Printf("hello yadayada...\n")
    // calculator(4,0)
    case_switch()
    num_array(4,5,6,8,9)
    // times_table(12)
    D_array()
}



func D_array() {
    var table [5][6]string
    for row := 0; row<5; row++ {
        for col := 0;col<6; col++{
            table[row][col] = strconv.Itoa(row) + "," + strconv.Itoa(col)
        }
    }
    fmt.Println(len(table))
}



func times_table(n int) {
    Outerloop:
        for i :=1; i<=n; i++ {
            for j :=1; j<=n; j++ {
                if i>=3 {
                    break Outerloop
                }
                fmt.Printf("%d * %d = %d\n", i, j, i*j)
            }
            fmt.Println("-------------------------------------")
            fmt.Println(time.Now().Date())
        }
}

func num_array(nums ... int) []int {
    // var nums [5]int
    // nums[0] = 4
    // nums[1] = 8
    // nums[2] = 6
    // nums[3] = 10
    labels:
        for k, v := range nums {
            // fmt.Println(v)
            if v == 0{
                break labels
            }
            // break labels

            if v%2==0 {
                fmt.Printf("The num %d at index %d is even!\n", v, k)
            }
        }
        return nil
}


// func Announce(message string, delay time.Duration) {
//     go func() {
//         time.Sleep(delay)
//         fmt.Println(message)
//     }()  // Note the parentheses - must call the function.
// }


// func calculator(a, b int)  {
//     if b == 0{
//         return 
//     }
//     mul := a*b
//     div := a/b
//     fmt.Printf("%d %d", &mul, &div)
//     // return mul, div, nil
// }


func case_switch() string{
    num := 4
    dayofWeek := ""
    switch num {
    case 1:
        dayofWeek = "Monday"
    case 2:
        dayofWeek = "Tuesday"
    case 3:
        dayofWeek = "Wednesday"
    case 4:
        dayofWeek = "Thursday"
    case 5:
        dayofWeek = "Friday"
    case 6:
        dayofWeek = "Saturday"
    case 7:
        dayofWeek = "Sunday"
    default:
        dayofWeek = "--error--"
    }
    fmt.Println(dayofWeek)
    return dayofWeek
}





