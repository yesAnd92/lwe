package navicat

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/xml"
	"errors"
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

		decrPwd, decrErr := decryptPwd(cons.Conns[idx].Password)
		if decrErr != nil {
			decrPwd = "can not decrypt password!"
		}
		cons.Conns[idx].Password = decrPwd
	}
	return &cons, nil
}

func decryptPwd(encryptTxt string) (string, error) {
	key := []byte("libcckeylibcckey")
	ciphertext, _ := hex.DecodeString(encryptTxt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := []byte("libcciv libcciv ")

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	return unPadding(ciphertext), nil
}

// unPadding  remove redundant padding data
func unPadding(src []byte) string {
	length := len(src)
	unpadding := int(src[length-1])
	return string(src[:(length - unpadding)])
}
