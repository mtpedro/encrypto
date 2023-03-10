package encrypto

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)


func Encrypt(key []byte, plaintext *[]byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    ciphertext := make([]byte, aes.BlockSize+len(*plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], *plaintext)

    return ciphertext, nil
}

func Decrypt(key []byte, ciphertext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)

    plaintext := make([]byte, len(ciphertext))
    stream.XORKeyStream(plaintext, ciphertext)

    return plaintext, nil
}

