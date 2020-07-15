package pkg

import (
	"fmt"
	"testing"
)

func TestHashAndSalt(t *testing.T) {
	pwd1 := "123456"
	pwd2 := "123abc"
	pwd3 := "ab$#$$nc,"

	after1, err := HashAndSalt(pwd1)
	if err != nil {
		t.Fatal("pwd1 encryption failed", err)
	}
	fmt.Println("pwd1 ", pwd1, "to: ", after1)

	if !ComparePassword(after1, pwd1) {
		t.Fatal("pw1 match failed")
	}
	if ComparePassword(after1, "a"+pwd1+"1") {
		t.Fatal("pw1 match failed")
	}

	after2, err := HashAndSalt(pwd2)
	if err != nil {
		t.Fatal("pw3 encryption failed", err)
	}
	fmt.Println("pwd2 ", pwd2, "to: ", after2)

	if !ComparePassword(after2, pwd2) {
		t.Fatal("pw2 match failed")
	}
	if ComparePassword(after2, pwd3) {
		t.Fatal("pw2 match failed")
	}

	after3, err := HashAndSalt(pwd3)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("pwd3 ", pwd3, "to: ", after3)
	if !ComparePassword(after3, pwd3) {
		t.Fatal("pw3 match failed")
	}
	if ComparePassword(after3, "") {
		t.Fatal("pw3 match failed")
	}
}
