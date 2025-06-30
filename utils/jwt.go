package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("q5HRZCOp3SsfZ2bCjz1vF+KDKcVLh7D1NpEFL1OY3dI=")

func GenerateJWT(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString(JwtKey)
}
