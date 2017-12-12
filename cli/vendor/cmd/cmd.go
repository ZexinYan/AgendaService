package cmd

import(
  "fmt"
  "model"
  // "net/http"
  //"io/ioutil"
  //"bufio"
  //"os"
  //"net/http"
)

//var url = "http://www.baidu.com"

// Register a user
func Register(user, pass, mail, phone string) int {
  fmt.Println("register")
  // Send HTTP
  fmt.Println(user + " " + pass + " " + mail + " " + phone)

  return 0
}

// Login Command
func Login(user, pass string) int {
  fmt.Println("Login")
  // Send HTTP

  fmt.Println(user + " " + pass)
  saveToken("123")
  return 0
}

// Logout Command
func Logout() int {
  // token := GetToken()
  // resp, err := http.Get(url + "/v1/auth?token=" + GetToken())
  fmt.Println("Logout")
  // HTTP

  if loginCheck() {
  	model.Logout()
  	fmt.Println("Log Out succeed!")
  } else {
  	fmt.Println("Log Out Failed")
  }
  return 0
}

// ShowUsers print all users when logged in
func ShowUsers() int {
  fmt.Println("ShowUsers")
  // HTTP

  if loginCheck() {
  	// HTTP
  	fmt.Println("Succeed")

  } else {
  	fmt.Println("Please Login First")
  }

  return 0
}

// DeleteUser delete current login user, and removed from its meeting
func DeleteUser() int {
  fmt.Println("delete User")

  if loginCheck() {
  	// HTTP

  } else {
  	fmt.Println("Please login first")
  }
  return 0
}

func getToken() string {
	t := model.LoadLoginFile()
	if t != nil {
		fmt.Println(t.Token)
		return t.Token
	} else {
		return ""
	}
}

func saveToken(token string) {
 	model.WriteLoginFile(token)
}

func loginCheck() bool {
	var t = getToken()
	if len(t) != 0 {
		return true
	} else {
		return false
	}
}
