package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	math_rand "math/rand"
	"os"
	"testing"

	"github.com/golang/glog"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()-=+_{}[];,.<>?~")
var CipherKey = "Lr7N>(g/D4g&mu>=6@!>aBAXY7.W<t5p"

func TestMain(m *testing.M) {
	defer glog.Flush()
	flag.Set("logtostderr", "true")
	flag.Parse()
	ret := m.Run()
	os.Exit(ret)
}

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[math_rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func encrypt(key []byte, message string) (encmess string, err error) {
	plainText := []byte(message)

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encmess = base64.URLEncoding.EncodeToString(cipherText)
	return
}

// GenEncryptedToken ...
func genEncryptedToken(userID int64, channelID uint16, token, cipherKey string) (string, error) {
	// 生成聊天的加密 token
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(userID))
	c := make([]byte, 2)
	binary.BigEndian.PutUint16(c, channelID)
	glog.Infof("b: %s\n", string(b))
	glog.Infof("c: %s\n", string(c))
	b = append(c[:], b[:]...)
	glog.Infof("string: %s\n", string(b))
	b = append(b[:], []byte(token)[:]...)
	glog.Infof("token: %s\n", token)
	glog.Infof("string: %s\n", string(b))
	return encrypt([]byte(cipherKey), string(b))
}

func genRedisKey(redisPrefix string, userID int64) string {
	return fmt.Sprintf("%s%d", redisPrefix, userID)
}

func TestTemp(t *testing.T) {
	var userID int64
	var channelID uint16
	userID = 1234
	channelID = 2
	redisPrefix := "chatToken_"
	token := randString(32)
	encryptedS, _ := genEncryptedToken(userID, channelID, token, CipherKey)
	glog.Infof("token:%s, key:%s", encryptedS, genRedisKey(redisPrefix, userID))
}
