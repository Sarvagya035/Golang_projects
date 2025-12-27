package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Random password Generator is Working...")

	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	lengthpwd := generateCmd.Int("length", 6, "This is the length of password")
	lower := generateCmd.Int("lower", 3, "Number of lowercase Characters")
	upper := generateCmd.Int("upper", 1, "Number of uppercase Characters")
	specialcase := generateCmd.Int("special", 1, "Number of special Characters")
	numerics := generateCmd.Int("number", 1, "Number of numeric Characters")

	generateCmd.Parse(os.Args[2:])

	fmt.Printf("The length of password is %v and number of lowercase letter is %v and the number of upercase letters is %v and special character are %v and numerics are %v", *lengthpwd, *lower, *upper, *specialcase, *numerics)

}
