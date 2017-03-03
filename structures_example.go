package main

import "fmt"

type Credentials struct {
	Username string
	Password string
}

func changePassword(cred *Credentials) {
	cred.Password = "changedpassword"
}

//call by reference
func (cred *Credentials) String() {
	fmt.Println("Username:", cred.Username, "\nPassword:", cred.Password)
}

// factory pattern
func newCredential(username string, password string) *Credentials {
	return &Credentials{Username: username, Password: password}
}

func main() {

	cred := &Credentials{Username: "admin", Password: "abc123"}
	cred.String()
	changePassword(cred)
	cred.String()

	fmt.Println("***********")
	cred1 := newCredential("admin", "admin")
	cred1.String()
	changePassword(cred1)
	cred1.String()

}
