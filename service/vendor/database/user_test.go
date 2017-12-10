package database

import (
	"entity"
	"log"
	"reflect"
	"sort"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func withDB(file string, f func()) {
	InitializeDB(file)
	log.Print("database set up")
	f()
	ClearDB()
	log.Print("database tear down")
}

func withTestDB(f func()) {
	withDB("Test.db", f)
}

func noError(t *testing.T, err error) bool {
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err.Error())
		return false
	}
	return true
}

func expectDeepEq(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Fail: Expected: %v, Actual: %v\n", b, a)
	} else {
		t.Log("Pass")
	}
}

func TestGetAllUsers(t *testing.T) {
	withTestDB(func() {
		t.Run("Empty Database", func(t *testing.T) {
			log.Print("Testing with Empty Databaes")
			us, err := GetAllUsers()
			if noError(t, err) {
				expectDeepEq(t, us, make([]*entity.User, 0))
			}
		})
		t.Run("With Incremental Users", func(t *testing.T) {
			log.Print("Testing with Incremental Users")
			insert := []*entity.User{
				{"foo", "fooooo", "foo@"}, {"bar", "barrrr", "bar@"},
				{"baz", "bazzzz", "baz@"}}
			now := make([]*entity.User, 0)
			for _, u := range insert {
				StoreUser(u)
				now = append(now, u)
				us, err := GetAllUsers()
				if noError(t, err) {
					sort.Sort(entity.UserSlice(us))
					sort.Sort(entity.UserSlice(now))
					for _, u := range now {
						log.Printf("%s %s %s\n", u.Username, u.Password, u.Email)
					}
					for _, u := range us {
						log.Printf("%s %s %s\n", u.Username, u.Password, u.Email)
					}
					expectDeepEq(t, us, now)
				}
			}
		})
	})
}

func TestStoreUser(t *testing.T) {
	type args struct {
		user *entity.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StoreUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("StoreUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemoveUser(t *testing.T) {

}

func TestGetUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUser(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
