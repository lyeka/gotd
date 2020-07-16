package auth

import (
	"fmt"
	"testing"
	"time"
)

const key = "12344%^&>"

func TestNewJwtToken(t *testing.T) {
	uid := "asd123"
	expiredAt := time.Now().Unix() + 2
	token, err := NewJwtToken(uid, expiredAt, key)
	if err != nil {
		t.Fatal("new jwt token failed, error: ", err)
	}
	fmt.Println(token)

	claims, err := ParseToken(token, key)
	if err != nil {
		t.Fatal("parse token failed, error: ", err)
	}
	if uid != claims.GetUserID() {
		t.Fatal("parse token content failed, content: ", claims.GetUserID())
	}

	time.Sleep(3 * time.Second)
	_, err = ParseToken(token, key)
	if err == nil {
		t.Fatal("token  过期机制失效")
	}

}
