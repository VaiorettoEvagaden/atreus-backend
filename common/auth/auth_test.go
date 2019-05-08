package auth

import (
	"fmt"
	"testing"
)

const SecretKey = "asdf"

func TestCreateToken(t *testing.T)  {
	token, _ := CreateToken([]byte(SecretKey), "YDQ", 2222, true)
	fmt.Println(token)
}
