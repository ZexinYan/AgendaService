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
					expectDeepEq(t, us, now)
				}
			}
		})
	})
}

func TestStoreUser(t *testing.T) {
	withTestDB(func() {
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
					expectDeepEq(t, us, now)
				}
			}
		})
	})
}

func TestRemoveUser(t *testing.T) {
	withTestDB(func() {
		t.Run("With decremental Users", func(t *testing.T) {
			insert := []*entity.User{
				{"foo", "fooooo", "foo@"}, {"bar", "barrrr", "bar@"},
				{"baz", "bazzzz", "baz@"}}
			now := make([]*entity.User, 0)
			for _, u := range insert {
				StoreUser(u)
				now = append(now, u)
			}
			for _, u := range insert {
				RemoveUser(u.Username)
				now = now[1:]
				us, _ := GetAllUsers()
				sort.Sort(entity.UserSlice(us))
				sort.Sort(entity.UserSlice(now))
				expectDeepEq(t, us, now)
			}
		})
	})
}

func TestGetUser(t *testing.T) {
	withTestDB(func() {
		t.Run("With decremental Users", func(t *testing.T) {
			insert := []*entity.User{
				{"foo", "fooooo", "foo@"}, {"bar", "barrrr", "bar@"},
				{"baz", "bazzzz", "baz@"}}
			for i, u := range insert {
				StoreUser(u)
				for j, u2 := range insert {
					u3, _ := GetUser(u2.Username)
					var n *entity.User
					if j <= i {
						expectDeepEq(t, u3, u2)
					} else {
						expectDeepEq(t, u3, n)
					}
				}
			}
		})
	})
}
