package cmd

import (
	"testing"
	"fmt"
)

func TestRegister(t *testing.T) {
	fmt.Println("Testing Register...")
	got := Register("yanzexin", "123", "mail", "1234")
	expect := 0

	if got != expect {
		t.Errorf("got [%s] expected [%s]", got, expect)
	}
}

func TestLogin(t *testing.T) {
	fmt.Println("Testing Login...")
	got := Login("yanzexin", "123")
	expect := 0

	if got != expect {
		t.Errorf("got [%s] expected [%s]", got, expect)
	}
}

func TestLogout(t *testing.T) {
	fmt.Println("Testing Log out...")
	Login("yanzexin", "123")
	got := Logout()
	expect := 0

	if got != expect {
		t.Errorf("got [%d] expected [%d]", got, expect)
	}
}

func TestShowInfo(t *testing.T) {
	fmt.Println("Testing show information...")
	Login("yanzexin", "123")
	got := ShowInfo("yanzexin")
	expect := 0

	if got != expect {
		t.Errorf("got [%d] expected [%d]", got, expect)
	}
}

func TestShowUsers(t *testing.T) {
	fmt.Println("Testing show users...")
	Login("yanzexin", "123")
	got := ShowUsers()
	expect := 0

	if got != expect {
		t.Errorf("got [%d] expected [%d]", got, expect)
	}
}

func TestDeleteUser(t *testing.T) {
	fmt.Println("Testing delete user...")
	Login("yanzexin", "123")
	got := DeleteUser("yanzexin")
	expect := 0

	if got != expect {
		t.Errorf("got [%s] expected [%s]", got, expect)
	}
}
