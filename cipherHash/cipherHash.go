package cipherHash

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

var (
	KEY, _ = hex.DecodeString("421A69BC2B99BEB97AA4BF13BE39D0344C9E31B853E646812F123DFE909F3D63")
)

func Encrypt(key []byte, message []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	b := message
	b = PKCS5Padding(b, aes.BlockSize)
	encMessage := make([]byte, len(b))
	iv := key[:aes.BlockSize]
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encMessage, b)

	return encMessage, nil
}

func Decrypt(key []byte, encMessage []byte) ([]byte, error) {
	iv := key[:aes.BlockSize]
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	if len(encMessage) < aes.BlockSize {
		return nil, errors.New("encMessage слишком короткий")
	}

	decrypted := make([]byte, len(encMessage))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, encMessage)

	return PKCS5UnPadding(decrypted), nil
}

func PKCS5Padding(cipher []byte, blockSize int) []byte {
	padding := blockSize - len(cipher)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipher, padText...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])

	return src[:(length - unPadding)]
}
