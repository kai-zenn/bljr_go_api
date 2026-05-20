package utils

import (
  "github.com/golang-jwt/jwt/v5"
  "time"
  "os"
  "github.com/google/uuid"
)


var secretkey = []byte(os.Getenv("JWT_SECRET"))

func CreateToken(userID uuid.UUID) (string, error) {
  claims := jwt.MapClaims{
    "user_id": userID.String(),
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  }
  
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  tokenString, err := token.SignedString(secretkey)
  if err != nil {
    return "", err
  }

  return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return secretkey, nil
  })
  if err != nil {
    return nil, err
  }

  return token, nil
}
