package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}

	switch os.Args[1] {
	case "hash":
		// hash the password
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid command: %v\n ", os.Args[1])
	}
}

func hash(password string) {
	// fmt.Printf("TODO: Hash the password %q\n", password)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error hashing %v\n", password)
		return
	}
	hash := string(hashedBytes)
	fmt.Println(hash)
}

func compare(password, hash string) {
	// fmt.Printf("TODO: Compare the password %q with the hash %q\n", password, hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("Password is invalid %v\n", password)
		return
	}

	fmt.Println("Password is correct!")
}
