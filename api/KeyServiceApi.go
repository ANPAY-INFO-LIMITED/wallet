package api

import (
	"fmt"
	"github.com/wallet/service"
)

func GenerateKeys(mnemonic string) string {
	keys, err := service.GenerateKeys(mnemonic, "")
	if err != nil {
		fmt.Printf("GenerateKeys->errl1:%x", err)
		return ""
	}
	publicKeyHex := fmt.Sprintf("%x", keys)
	fmt.Printf("publ1:%x", publicKeyHex)
	fmt.Printf("errl1:%x", err)
	return publicKeyHex
}
func GetPrivateKey() (string, error) {
	return service.GetPrivateKey()
}

func GetPublicKey() string {
	pb, err := service.GetPublicRawKey()
	if err != nil {
		fmt.Printf("GetPublicKey->errl1:%x", err)
		return ""
	}
	return pb
}
func GetRawPublicKey() string {
	pb, _, err := service.GetPublicKey()
	if err != nil {
		fmt.Printf("GetPublicKey->errl1:%x", err)
		return ""
	}
	return pb
}
func GetNPublicKey() string {
	_, nPub, err := service.GetPublicKey()
	if err != nil {
		fmt.Printf("GetNPublicKey->errl1:%x", err)
		return ""
	}
	return nPub
}
func GetNBPublicKey() string {
	_, nPub, err := service.GetNewPublicKey()
	if err != nil {
		fmt.Printf("GetNPublicKey->errl1:%x", err)
		return ""
	}
	return nPub
}
func GetExistPublicKey() string {
	_, nPub, err := service.GetExistPublicKey()
	if err != nil {
		fmt.Printf("GetExistPublicKey->errl1:%x", err)
		return ""
	}
	return nPub
}

func GetJsonPublicKey() string {
	keyInfo, err := service.GetJsonPublicKey()
	if err != nil {
		fmt.Printf("GetNPublicKey->errl1:%x", err)
		return ""
	}
	return keyInfo
}

func SignMess(message string) string {
	sign, err := service.SignMessage(message)
	if err != nil {
		fmt.Printf("SignMess->errl1:%x", err)
		return ""
	}
	return sign
}
func Main111() {
	GenerateKeys("about system unfair bullet bean endorse control wrist isolate pulse warfare risk picnic owner fantasy chase truck amazing clip marine rib baby attract blade")
	fmt.Printf("NBPublicKey:->" + GetNBPublicKey())
}
