package database

import (
	"testing"
	"testutil"
)

type ut struct {
	user  string
	token string
}

func TestGetToken(t *testing.T) {
	WithTestDB(func() {
		uts := []ut{{"foo", "fooooo"}, {"bar", "barrrrrr"}, {"baz", "bazzzzzz"}}
		t.Run("With Empty Database", func(t *testing.T) {
			for _, uT := range uts {
				tok, e := GetToken(uT.user)
				testutil.ExpectDeepEq(t, tok, "")
				if e == nil {
					t.Errorf("should get error when querying empty db")
				}
			}
		})
		t.Run("With Adding Tokens Incrementally", func(t *testing.T) {
			for _, uT := range uts {
				e := PutToken(uT.user, uT.token)
				if testutil.NoError(t, e) {
					tok, e := GetToken(uT.user)
					if testutil.NoError(t, e) {
						testutil.ExpectDeepEq(t, tok, uT.token)
					}
				}
			}
		})
	})
}

func TestPutToken(t *testing.T) {
	WithTestDB(func() {
		uts := []ut{{"foo", "fooooo"}, {"bar", "barrrrrr"}, {"baz", "bazzzzzz"}}
		for _, uT := range uts {
			e := PutToken(uT.user, uT.token)
			if testutil.NoError(t, e) {
				tok, e := GetToken(uT.user)
				if testutil.NoError(t, e) {
					testutil.ExpectDeepEq(t, tok, uT.token)
				}
			}
		}
	})
}

func TestDeleteToken(t *testing.T) {
	WithTestDB(func() {
		uts := []ut{{"foo", "fooooo"}, {"bar", "barrrrrr"}, {"baz", "bazzzzzz"}}
		t.Run("With Empty Database", func(t *testing.T) {
			for _, uT := range uts {
				err := DeleteToken(uT.token)
				testutil.NoError(t, err)
			}
		})
		t.Run("Adding data then delete", func(t *testing.T) {
			for _, uT := range uts {
				PutToken(uT.user, uT.token)
			}
			for _, uT := range uts {
				DeleteToken(uT.token)
				_, err := GetToken(uT.user)
				if err == nil {
					t.Errorf(
						"token %s should be deleted, but still has user %s",
						uT.token, uT.user)
				}
			}
		})
	})
}

func TestDeleteTokenByUsername(t *testing.T) {
	WithTestDB(func() {
		uts := []ut{{"foo", "fooooo"}, {"bar", "barrrrrr"}, {"baz", "bazzzzzz"}}
		t.Run("With Empty Database", func(t *testing.T) {
			for _, uT := range uts {
				err := DeleteTokenByUsername(uT.user)
				testutil.NoError(t, err)
			}
		})
		t.Run("Adding data then delete", func(t *testing.T) {
			for _, uT := range uts {
				PutToken(uT.user, uT.token)
			}
			for _, uT := range uts {
				DeleteTokenByUsername(uT.user)
				_, err := GetToken(uT.user)
				if err == nil {
					t.Errorf(
						"token %s should be deleted, but still has user %s",
						uT.token, uT.user)
				}
			}
		})
	})
}

func TestHasToken(t *testing.T) {
	WithTestDB(func() {
		uts := []ut{{"foo", "fooooo"}, {"bar", "barrrrrr"}, {"baz", "bazzzzzz"}}
		t.Run("With Empty Database", func(t *testing.T) {
			for _, uT := range uts {
				b, e := HasToken(uT.token)
				if testutil.NoError(t, e) {
					testutil.ExpectDeepEq(t, b, false)
				}
			}
		})
		t.Run("With Adding Tokens Incrementally", func(t *testing.T) {
			for _, uT := range uts {
				e := PutToken(uT.user, uT.token)
				if testutil.NoError(t, e) {
					b, e := HasToken(uT.token)
					if testutil.NoError(t, e) {
						testutil.ExpectDeepEq(t, b, true)
					}
				}
			}
		})
	})
}

func TestGetUsername(t *testing.T) {
	WithTestDB(func() {
		uts := []ut{{"foo", "fooooo"}, {"bar", "barrrrrr"}, {"baz", "bazzzzzz"}}
		t.Run("With Empty Database", func(t *testing.T) {
			for _, uT := range uts {
				u, e := GetUsername(uT.token)
				if testutil.NoError(t, e) {
					testutil.ExpectDeepEq(t, u, "")
				}
			}
		})
		t.Run("With Adding Tokens Incrementally", func(t *testing.T) {
			for _, uT := range uts {
				e := PutToken(uT.user, uT.token)
				if testutil.NoError(t, e) {
					u, e := GetUsername(uT.token)
					if testutil.NoError(t, e) {
						testutil.ExpectDeepEq(t, u, uT.user)
					}
				}
			}
		})
	})
}
