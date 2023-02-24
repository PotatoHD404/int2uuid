package main

import (
	"bytes"
	"crypto/aes"
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
)

func intToUuid(value uint64, key []byte) string {
	block, _ := aes.NewCipher(key)
	v := make([]byte, 16)
	binary.BigEndian.PutUint64(v[8:], value)
	a := make([]byte, 16)
	block.Encrypt(a, v)
	uid, _ := uuid.FromBytes(a)
	return uid.String()
}

func uuidToInt(value string, key []byte) int {
	block, _ := aes.NewCipher(key)
	parse, _ := uuid.Parse(value)
	c := make([]byte, 16)
	block.Decrypt(c, parse[:])
	c = c[8:]

	return int(binary.BigEndian.Uint64(c))
}

func main() {
	value := uint64(1 << 33)
	fmt.Println(value)
	key := bytes.Repeat([]byte{0}, 16)

	uid := intToUuid(value, key)
	fmt.Println(uid)
	fmt.Println(uuidToInt(uid, key))
}
