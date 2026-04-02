package generator

import (
    "crypto/rand"
    "errors"
    "math/big"
)

const (
    letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    alphanumeric  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    allChars      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"
)

// GeneratePasswordWithType gera uma senha aleatória de acordo com o tipo
func GeneratePasswordWithType(length int, charType string) (string, error) {
    var charset string

    switch charType {
    case "letters":
        charset = letters
    case "alphanumeric":
        charset = alphanumeric
    case "all":
        charset = allChars
    default:
        return "", errors.New("tipo inválido! use letters, alphanumeric ou all")
    }

    password := make([]byte, length)
    for i := 0; i < length; i++ {
        index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        if err != nil {
            return "", err
        }
        password[i] = charset[index.Int64()]
    }

    return string(password), nil
}