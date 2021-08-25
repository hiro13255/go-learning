package main

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Name string
	Age int
	jwt.StandardClaims
}

// jwtトークン作成
func createTokenString() string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &User{
		Name: "Taro",
		Age: 30,
	})

	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		log.Fatalln(err)
	}

	return tokenstring
}

func main() {
	// token作成
	tokenstring := createTokenString()

	log.Println(tokenstring)
	user := User{}
	
	// Jwt検証
	token, err := jwt.ParseWithClaims(tokenstring, &user, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})

	log.Println(token.Valid, user, err)
}