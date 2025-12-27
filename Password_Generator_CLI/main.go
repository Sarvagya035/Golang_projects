package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	lowercharset     = "abcdefghijklmnopqrstuvwxyz"
	uppercharset     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialcharset   = "!@#$%^&*()_+-="
	numericalcharset = "0123456789"
	randomcharset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+-=0123456789"
)

func main() {

	var (
		passwordlength int
		lowercasechar  int
		uppercasechar  int
		specialchar    int
		numericalchar  int
	)

	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	lengthpwd := generateCmd.Int("length", 10, "This is the length of password")
	lower := generateCmd.Int("minlower", 3, "Number of lowercase Characters")
	upper := generateCmd.Int("minupper", 2, "Number of uppercase Characters")
	specialcase := generateCmd.Int("minspecial", 2, "Number of special Characters")
	numerics := generateCmd.Int("minnumber", 3, "Number of numeric Characters")

	if len(os.Args) < 2 {
		fmt.Println("expected 'generate' subcommand")
		os.Exit(1)
	}

	generateCmd.Parse(os.Args[2:])

	passwordlength = *lengthpwd
	lowercasechar = *lower
	uppercasechar = *upper
	specialchar = *specialcase
	numericalchar = *numerics
	remainingchar := passwordlength - (lowercasechar + uppercasechar + specialchar + numericalchar)

	if remainingchar < 0 {

		fmt.Println("Please provide a valid length of password")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	generatedpassword := generatePassword(lowercasechar, uppercasechar, specialchar, numericalchar, remainingchar)

	fmt.Println("The generated password is", generatedpassword)

}

func generatePassword(lowercasechar int, uppercasechar int, specialchar int, numericalchar int, remainingchar int) string {

	password := ""

	//generating lowercase char

	for i := 0; i < lowercasechar; i++ {
		n := len(lowercharset)
		idx := rand.Intn(n)
		password = password + string(lowercharset[idx])
	}

	//generatingspecialchar

	for i := 0; i < specialchar; i++ {
		n := len(specialcharset)
		idx := rand.Intn(n)
		password = password + string(specialcharset[idx])
	}

	//generating uppercase letter

	for i := 0; i < uppercasechar; i++ {
		n := len(uppercharset)
		idx := rand.Intn(n)
		password = password + string(uppercharset[idx])
	}

	//genrating numerics

	for i := 0; i < numericalchar; i++ {
		n := len(numericalcharset)
		idx := rand.Intn(n)
		password = password + string(numericalcharset[idx])
	}

	//adding remaining char

	for i := 0; i < remainingchar; i++ {
		n := len(randomcharset)
		idx := rand.Intn(n)
		password = password + string(randomcharset[idx])
	}

	//shuffle the password string

	passwordRune := []rune(password)
	rand.Shuffle(len(password), func(i int, j int) {
		passwordRune[i], passwordRune[j] = passwordRune[j], passwordRune[i]
	})

	password = string(passwordRune)

	return password

}
