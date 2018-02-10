package global

import (
	"bytes"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"git.oschina.net/dxvgef/go-lib/convert"
	"strconv"
	"time"
)

type (
	AuthTokenKey struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Type     string `json:"type"`
	}
	AuthTokenValue struct{} //redis里登录value
)

//var EncryptSalt string = "ocpB8nZG5yBWrMfJDsM2fRB5L5LERMF47A6PWAC4wpM="

//从[]byte生成md5密文
func MD5ByBytes(value []byte, salt ...[]byte) (result string, err error) {
	var s []byte
	if len(salt) > 0 {
		s = salt[0]
	}
	h := md5.New()
	_, err = h.Write(value)
	if err != nil {
		GlobalLogger.Error("error:%s", err.Error())
		return
	}
	if s != nil {
		_, err = h.Write(s)
		if err != nil {
			GlobalLogger.Error("error:%s", err.Error())
			return
		}
	}
	result = hex.EncodeToString(h.Sum(nil))
	return
}

//从str生成md5密文
func MD5ByStr(value string, salt ...string) (result string, err error) {
	var s []byte
	if len(salt) > 0 {
		s = convert.StrToByte(salt[0])
	}
	h := md5.New()
	_, err = h.Write(convert.StrToByte(value))
	if err != nil {
		GlobalLogger.Error("error:%s", err.Error())
		return
	}
	if s != nil {
		_, err = h.Write(s)
		if err != nil {
			GlobalLogger.Error("error:%s", err.Error())
			return
		}
	}
	result = hex.EncodeToString(h.Sum(nil))
	return
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptAuthKey(key *AuthTokenKey) (token string, err error) {
	token, err = MD5ByStr(strconv.Itoa(key.Id), key.Username, key.Type, EncryptSalt, time.Now().String())
	return
}

func EncryptAuthToken(value *AuthTokenValue) (token string, err error) {
	authTokenBytes, err := json.Marshal(value)
	if err != nil {
		return
	}
	token = base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(authTokenBytes)
	return
}

func ParseAuthToken(tokenStr string) (*AuthTokenValue, error) {
	var token AuthTokenValue
	var err error
	authTokenBytes, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(tokenStr)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(authTokenBytes, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

//从str生成Base64密文
func EncryptBase64Str(val, SiteId string) (token string) {
	val = val + SiteId
	valarr := []byte(val)
	token = base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(valarr)
	return
}

//从Base64密文中解密字符串
func ParseBase64Str(val, SiteId string) (token string, err error) {
	tokenStr, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(val)
	var tokenlen = len(tokenStr) - len(SiteId)
	token = string(tokenStr[:tokenlen])
	return
}

func DesEncrypt(src, key []byte) (string, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	//src = ZeroPadding(src, bs)
	src = PKCS5Padding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}

	return base64.StdEncoding.EncodeToString(out), nil
}

//func DesDecrypt(src, key []byte) ([]byte, error) {
//	src, err := Base64Decode_(string(src))
//	if err != nil {
//		return nil, err
//	}
//	block, err := des.NewCipher(key)
//
//	if err != nil {
//		return nil, err
//	}
//	out := make([]byte, len(src))
//	dst := out
//	bs := block.BlockSize()
//	if len(src) % bs != 0 {
//		return nil, errors.New("crypto/cipher: input not full blocks")
//	}
//	for len(src) > 0 {
//		block.Decrypt(dst, src[:bs])
//		src = src[bs:]
//
//		dst = dst[bs:]
//
//	}
//	//out = ZeroUnPadding(out)
//
//	out = PKCS5UnPadding(out)
//
//
//	return out, nil
//}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
