package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"strconv"
)

func main() {
	BirthdayExportWallet("0000")
}

func generateKey(privateKeys string) (address common.Address, privKey string) {
	privateKey, err := crypto.HexToECDSA(privateKeys)
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address = crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, hexutil.Encode(privateKeyBytes)[2:]

}

func BirthdayExportWallet(birthday string) {
	if len(birthday) != 4 {
		panic("birthday number length is 4 ")
	}
	distInt64, _ := strconv.ParseInt(birthday, 10, 64)
	birthday = fmt.Sprintf("%x", distInt64)
	if len(birthday) == 3 {
		birthday = "0" + birthday
	} else if len(birthday) == 2 {
		birthday = "00" + birthday
	}

	var birthday16 = ""
	for i := 0; i < 16; i++ {
		birthday16 += birthday
	}
	address, privateKye := generateKey(birthday16)
	fmt.Println("private key", privateKye)
	fmt.Println("address", address.Hex())
}
