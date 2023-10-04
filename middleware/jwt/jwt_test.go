package jwt

import (
	"testing"
	"tiktok/common"
)

func TestJWT(t *testing.T) {
	user := common.User{
		UserId:   1,
		Account:  "test_account",
		Username: "test_user",
		Password: "secret",
	}

	token, err := IssueToken(user.Account)
	if err != nil {
		t.Errorf("Get token failed:%v", err)
	}
	claims, err := ParseToken(token)
	if err != nil {
		t.Errorf("Parse token failed:%v", err)
	}

	if claims.Account != user.Account {
		t.Errorf("claims doesn`t match user`")
	}
}

func TestEmptyJWT(t *testing.T) {
	_, err := ParseToken("")
	if err == nil {
		t.Errorf("Empty token shouldn't be allowed")
	}
}
