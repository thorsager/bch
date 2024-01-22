package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var (
	flagPwd  string
	flagCost int
)

func init() {
}

func main() {
	flag.StringVar(&flagPwd, "p", "", "password to hash")
	flag.IntVar(&flagCost, "c", bcrypt.DefaultCost, "cost of bcrypt hashing")
	flag.Parse()

	if flagCost < bcrypt.MinCost || flagCost > bcrypt.MaxCost {
		fmt.Println("invalid cost")
		return
	}
	if flagPwd == "" {
		fmt.Println("invalid password")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(flagPwd), flagCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(hash))
}
