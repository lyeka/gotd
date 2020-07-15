package pkg

import "golang.org/x/crypto/bcrypt"

// ref  https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72

// HashAndSalt 加密
// 同样明文每次加密后的密文都不一致
func HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword 比对
func ComparePassword(hashedPwd, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if err != nil {
		return false
	}

	return true
}
