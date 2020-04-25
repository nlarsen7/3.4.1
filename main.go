package main

import "fmt"

func age() {
var user string
var age int
fmt.Println("What is you name?")
fmt.Scanln(&user)
fmt.Println("what is your age in years?")
fmt.Scanln(&age)
fmt.Println("Hello",user,"you are",age,"years old")
}
func main (){
  age()
}