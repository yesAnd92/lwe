package navicat

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"github.com/spf13/cobra"
)

const (
	//Navicat加密时使用的key和iv
	AES_KEY = "libcckeylibcckey"
	AES_IV  = "libcciv libcciv "
)

type NxcConnections struct {
	Conns   []NxcConn `xml:"Connection"`
	Version string    `xml:"Ver,attr"`
}

type NxcConn struct {
	ConnectionName string `xml:"ConnectionName,attr"`
	ConnType       string `xml:"ConnType,attr"`
	Host           string `xml:"Host,attr"`
	UserName       string `xml:"UserName,attr"`
	Port           string `xml:"Port,attr"`
	Password       string `xml:"Password,attr"`
}

func ParseNcx(data []byte) (*NxcConnections, error) {

	cons := NxcConnections{}
	err := xml.Unmarshal(data, &cons)
	if err != nil {
		return nil, errors.New("ncx file format is incorrect！")
	}
	for idx := range cons.Conns {

		decrPwd := decryptPwd(cons.Conns[idx].Password)

		cons.Conns[idx].Password = decrPwd
	}
	return &cons, nil
}

//decryptPwd navicat的加密规则可以参照这个文档
//https://github.com/HyperSine/how-does-navicat-encrypt-password/blob/master/doc/how-does-navicat-encrypt-password.md
func decryptPwd(encryptTxt string) string {
	key := []byte(AES_KEY)
	ciphertext, _ := hex.DecodeString(encryptTxt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		cobra.CheckErr("Decrypt password failed! Please confirm export connection with password.")
	}
	iv := []byte(AES_IV)

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(ciphertext, ciphertext)

	return unPadding(ciphertext)
}

// unPadding  remove redundant padding data
func unPadding(src []byte) string {
	length := len(src)
	unpadding := int(src[length-1])
	return string(src[:(length - unpadding)])
}
