package storage

import "testing"

func TestCreateUser(t *testing.T) {
	user := &UserInfo{
		Name: "Alice",
	}
	us := &UserStorage{}
	if err := us.CreateUser(user); err != nil {
		if err != ErrUserExist {
			t.Error(err)
			t.FailNow()
		}
	}
	if err := us.CreateUser(user); err != nil {
		if err != ErrUserExist {
			t.Error(err)
			t.FailNow()
		}
	}
	user = &UserInfo{
		Name: "Bob",
	}
	if err := us.CreateUser(user); err != nil {
		if err != ErrUserExist {
			t.Error(err)
			t.FailNow()
		}
	}
	user = &UserInfo{
		Name: "Carol",
	}
	if err := us.CreateUser(user); err != nil {
		if err != ErrUserExist {
			t.Error(err)
			t.FailNow()
		}
	}
	t.Logf("user=%v", user)
}

func TestListAllUsers(t *testing.T) {
	us := &UserStorage{}
	users, err := us.ListAllUsers()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, user := range users {
		t.Logf("user=%v", user)
	}
}
