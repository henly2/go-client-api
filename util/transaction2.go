package util

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"github.com/golang/protobuf/proto"
	"github.com/sasaxie/go-client-api/common/crypto"
	"github.com/sasaxie/go-client-api/core"
	"time"
	"github.com/sasaxie/go-client-api/common/hexutil"
	"github.com/sasaxie/go-client-api/api"
)

func SignTransaction2(transaction *core.Transaction, key *ecdsa.PrivateKey) error {
	transaction.GetRawData().Timestamp = time.Now().UnixNano() / 1000000

	rawData, err := proto.Marshal(transaction.GetRawData())
	if err != nil {
		return err
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	contractList := transaction.GetRawData().GetContract()

	for range contractList {
		signature, err := crypto.Sign(hash, key)
		if err != nil {
			return err
		}

		transaction.Signature = append(transaction.Signature, signature)
	}

	return nil
}

func GetTxHash(tx *core.Transaction) (string, error) {
	rawData := tx.RawData

	rawDataBytes, err := proto.Marshal(rawData)
	if err != nil {
		return "", err
	}

	h256 := sha256.New()
	h256.Write(rawDataBytes)
	txHash := h256.Sum(nil)

	return hexutil.Encode(txHash), nil
}

func GetTxHash_Ext(tx *api.TransactionExtention) (string, error) {
	rawData := tx.Transaction.RawData

	rawDataBytes, err := proto.Marshal(rawData)
	if err != nil {
		return "", err
	}

	h256 := sha256.New()
	h256.Write(rawDataBytes)
	txHash := h256.Sum(nil)

	return hexutil.Encode(txHash), nil
}

func ToString(transaction *core.Transaction) (string, error) {
	b, err := ToData(transaction)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(b), nil
}

func ToData(transaction *core.Transaction) ([]byte, error) {
	b, err := proto.Marshal(transaction)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func FromString(text string) (*core.Transaction, error) {
	b, err := hexutil.Decode(text)
	if err != nil {
		return nil, err
	}

	return FromData(b)
}

func FromData(data []byte) (*core.Transaction, error) {
	transaction := &core.Transaction{}
	err := proto.Unmarshal(data, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}