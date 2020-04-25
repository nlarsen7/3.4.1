//Nicholas Larsen
//4/25/2020
//Creates a function that asks the user's age and name and outputs it
package main

import "fmt"

func age() {
var user string
var age int
//variables for age and user name
fmt.Println("What is you name?")
fmt.Scanln(&user)
fmt.Println("what is your age in years?")
fmt.Scanln(&age)
//asks for user name and age and stores in variable
fmt.Println("Hello",user,"you are",age,"years old")
//Prints out the name and age
}
func main (){
  age()
  //another function that does what the age function does
}