package database

import (
	"entity"
	"log"
	"sort"
	"testing"
	"testutil"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetAllUsers(t *testing.T) {
	WithTestDB(func() {
		t.Run("Empty Database", func(t *testing.T) {
			log.Print("Testing with Empty Databaes")
			us, err := GetAllUsers()
			if testutil.NoError(t, err) {
				testutil.ExpectDeepEq(t, us, make([]*entity.User, 0))
			}
		})
		t.Run("With Incremental Users", func(t *testing.T) {
			log.Print("Testing with Incremental Users")
			insert := []*entity.User{
				{"foo", "fooooo", "foo@", "11"}, {"bar", "barrrr", "bar@", "2"},
				{"baz", "bazzzz", "baz@", "33"}}
			now := make([]*entity.User, 0)
			for _, u := range insert {
				StoreUser(u)
				now = append(now, u)
				us, err := GetAllUsers()
				if testutil.NoError(t, err) {
					sort.Sort(entity.UserSlice(us))
					sort.Sort(entity.UserSlice(now))
					testutil.ExpectDeepEq(t, us, now)
				}
			}
		})
	})
}

func TestStoreUser(t *testing.T) {
	WithTestDB(func() {
		t.Run("With Incremental Users", func(t *testing.T) {
			log.Print("Testing with Incremental Users")
			insert := []*entity.User{
				{"foo", "fooooo", "foo@", "11"}, {"bar", "barrrr", "bar@", "2"},
				{"baz", "bazzzz", "baz@", "33"}}
			now := make([]*entity.User, 0)
			for _, u := range insert {
				StoreUser(u)
				now = append(now, u)
				us, err := GetAllUsers()
				if testutil.NoError(t, err) {
					sort.Sort(entity.UserSlice(us))
					sort.Sort(entity.UserSlice(now))
					testutil.ExpectDeepEq(t, us, now)
				}
			}
		})
	})
}

func TestRemoveUser(t *testing.T) {
	WithTestDB(func() {
		t.Run("With decremental Users", func(t *testing.T) {
			insert := []*entity.User{
				{"foo", "fooooo", "foo@", "11"}, {"bar", "barrrr", "bar@", "2"},
				{"baz", "bazzzz", "baz@", "33"}}
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
				testutil.ExpectDeepEq(t, us, now)
			}
		})
	})
}

func TestGetUser(t *testing.T) {
	WithTestDB(func() {
		t.Run("With decremental Users", func(t *testing.T) {
			insert := []*entity.User{
				{"foo", "fooooo", "foo@", "11"}, {"bar", "barrrr", "bar@", "2"},
				{"baz", "bazzzz", "baz@", "33"}}
			for i, u := range insert {
				StoreUser(u)
				for j, u2 := range insert {
					u3, _ := GetUser(u2.Username)
					var n *entity.User
					if j <= i {
						testutil.ExpectDeepEq(t, u3, u2)
					} else {
						testutil.ExpectDeepEq(t, u3, n)
					}
				}
			}
		})
	})
}
