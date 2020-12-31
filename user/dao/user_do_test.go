package dao

import (
	"testing"

)

func TestUserDAOImpl_Save(t *testing.T) {
	userDAO := &UserDAOImpl{}
	err := InitMysql("127.0.0.1", "3306", "root", "root", "user")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	user := &UserEntity{
		Username: "pawn",
		Password: "pawn",
		Email: "pawn@gamil.com",
	}
	err = userDAO.Save(user)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("new User ID is %d", user.ID)
}

func TestUserDAOImpl_SelectByEmail(t *testing.T) {
	userDAO := &UserDAOImpl{}
	err := InitMysql("127.0.0.1", "3306", "root", "root", "user")
	if err != nil{
		t.Error(err)
		t.FailNow()
	}
	user, err := userDAO.SelectByEmail("pawn@gmai.com")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("result username is %s", user.Username)
}