# Random Password Generator CLI 

A simple Random password generator CLI built using **Golang**.  
This is a simple CLI to generate random password as per user input and uses basic golang packages like rand/math, flag and os.

---

## ✨ Features
- Generate passwords with custom length
- Control minimum lowercase, uppercase, numeric and special characters
- Command-line interface using Go standard library

---

## ▶️ How to Run the Project
Clone the repository, move into the project directoryand  run the application using the commands below:

### Usage
```bash
git clone https://github.com/Sarvagya035/Golang_projects.git

cd Password_Generator_CLI

/*
    A sample input is shown below.
    please make sure that sum of lower character, upper character, special character and numeric character must be less than or equal to length of password
*/

go run main.go generate --length 12 --minlower 4 --minupper 5 --minspecial 1 --minnumber 2

// Another example includes usage of default values if flags are not provides

go run main.go generate --length 16