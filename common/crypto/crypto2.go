package crypto

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sasaxie/go-client-api/common/hexutil"
	"bytes"
	"fmt"
)

const TestAddressPrefix = "a0"
const MainAddressPrefix = "41"

var (
	_netPrefix = MainAddressPrefix
)

func SetNetPrefix(prefix string)  {
	_netPrefix = prefix
}
func GetNetPrefix() string {
	return _netPrefix
}

func PubkeyToAddress2(p ecdsa.PublicKey) (Address, error) {
	address := crypto.PubkeyToAddress(p)

	addressTron := make([]byte, AddressLength)

	addressPrefix, err := hexutil.Decode(GetNetPrefix())
	if err != nil {
		return Address{}, err
	}

	addressTron = append(addressTron, addressPrefix...)
	addressTron = append(addressTron, address.Bytes()...)

	return BytesToAddress(addressTron), nil
}

func VerifyAddress(data []byte) error {
	address := Address{}
	address.SetBytes(data)

	addressPrefix, err := hexutil.Decode(GetNetPrefix())
	if err != nil {
		return err
	}

	if 0 != bytes.Compare(addressPrefix, address[:len(addressPrefix)]) {
		return fmt.Errorf("prefix not right")
	}

	return nil
}