package util

import (
	"crypto/sha256"
	"github.com/golang/protobuf/proto"
	"github.com/sasaxie/go-client-api/core"
	"github.com/sasaxie/go-client-api/common/hexutil"
	"encoding/hex"
	"math/big"
	"strings"
	"github.com/o3labs/neo-utils/neoutils"
	"github.com/sasaxie/go-client-api/api"
)

const (
	paddingZeor16 = "0000000000000000"
)

func GetBlockHash2(block *core.Block) (string, error) {
	rawData := block.BlockHeader.RawData

	rawDataBytes, err := proto.Marshal(rawData)
	if err != nil {
		return "", nil
	}

	h256 := sha256.New()
	h256.Write(rawDataBytes)
	blockHash := h256.Sum(nil)
	hs2 := hexutil.Encode(blockHash)

	numberByte, err := IntegerToByteArray(block.BlockHeader.RawData.Number)
	if err != nil {
		return "", nil
	}
	hs1 := paddingZeor16 + numberByte
	hs1Index := len(hs1) - 16

	hs := hs1[hs1Index:]
	hs += hs2[16:]

	return hs, nil
}

func GetBlockHash2_Ext(block *api.BlockExtention) (string, error) {
	rawData := block.BlockHeader.RawData

	rawDataBytes, err := proto.Marshal(rawData)
	if err != nil {
		return "", nil
	}

	h256 := sha256.New()
	h256.Write(rawDataBytes)
	blockHash := h256.Sum(nil)
	hs2 := hexutil.Encode(blockHash)

	numberByte, err := IntegerToByteArray(block.BlockHeader.RawData.Number)
	if err != nil {
		return "", nil
	}
	hs1 := paddingZeor16 + numberByte
	hs1Index := len(hs1) - 16

	hs := hs1[hs1Index:]
	hs += hs2[16:]

	return hs, nil
}

func ByteArrayToInteger(str string) (*big.Int, error) {
	val := remove0x(str)
	b, err := hex.DecodeString(val)
	if err != nil {
		return nil, err
	}
	d := new(big.Int).SetBytes(b)
	return d, nil
}

func IntegerToByteArray(value int64) (string, error) {
	d := new(big.Int).SetInt64(value)
	b := d.Bytes()
	val := hex.EncodeToString(b)

	return val, nil
}

func remove0x(data string) (string) {
	index := strings.Index(data, "0x")
	if index == 0 {
		d := data[2:]
		return d
	}
	return data
}

// no test
func toBigEndian(data string) (string, error) {
	// if has 0x..., this is bigendian
	// else this is littleendian
	index := strings.Index(data, "0x")
	if index == 0 {
		return data, nil
	}

	d, err := hex.DecodeString(data)
	if err != nil {
		return "", nil
	}

	dr := neoutils.ReverseBytes(d)
	return hex.EncodeToString(dr), nil
}