package controller

import (
	"fmt"
	"pkgArch/entity"
	"testing"
	"time"
)

type mockupUserController struct {
	NotifyFn func(usr entity.User, msg string)
}

func (uc *mockupUserController) Notify(usr entity.User, msg string) {
	uc.NotifyFn(usr, msg)
}

func TestAccountController_Transfer(t *testing.T) {
	// arrange
	usr1 := &entity.User{Name: "John", Email: "j@connor.name"}
	usr2 := &entity.User{Name: "Clara", Email: "c@hiller.name"}

	acnt1 := &entity.Account{
		Balance: 100,
		Owner:   usr1,
	}

	acnt2 := &entity.Account{
		Balance: 0,
		Owner:   usr2,
	}

	uc := &mockupUserController{func(u entity.User, m string) {
		fmt.Printf("user %v got message '%v'\n", u, m)
	}}
	ac := &AccountController{uc}

	// act
	err := ac.Transfer(acnt1, acnt2, 1, time.Now().Add(time.Microsecond*50))

	// assert
	if err != nil {
		t.Error(err)
	}

	if acnt1.Balance != 99 {
		t.Errorf("withdrawal failed: new balance of source is %v", acnt1.Balance)
	}

	if acnt2.Balance != 1 {
		t.Errorf("reception failed: new balance of target is %v", acnt2.Balance)
	}
}
