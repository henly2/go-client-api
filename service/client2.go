package service

import (
	"context"
	"github.com/sasaxie/go-client-api/api"
	"github.com/sasaxie/go-client-api/common/base58"
	"github.com/sasaxie/go-client-api/common/hexutil"
	"github.com/sasaxie/go-client-api/core"
	"fmt"
)

func (g *GrpcClient) ListNodes2() (*api.NodeList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), GrpcTimeout)
	defer cancel()

	nodeList, err := g.Client.ListNodes(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, err
	}

	return nodeList, nil
}

func (g *GrpcClient) GetAccount2(address string) (*core.Account, error) {
	account := new(core.Account)
	account.Address = base58.DecodeCheck(address)

	ctx, cancel := context.WithTimeout(context.Background(), GrpcTimeout)
	defer cancel()

	result, err := g.Client.GetAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (g *GrpcClient) GetNowBlock2_Ext() (*api.BlockExtention, error) {
	ctx, cancel := context.WithTimeout(context.Background(), GrpcTimeout)
	defer cancel()

	result, err := g.Client.GetNowBlock2(ctx, new(api.EmptyMessage))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (g *GrpcClient) GetBlockByNum2_Ext(num int64) (*api.BlockExtention, error) {
	numMessage := new(api.NumberMessage)
	numMessage.Num = num

	ctx, cancel := context.WithTimeout(context.Background(), GrpcTimeout)
	defer cancel()

	result, err := g.Client.GetBlockByNum2(ctx, numMessage)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (g *GrpcClient) GetTransactionById2(id string) (*core.Transaction, error) {
	transactionId := new(api.BytesMessage)
	var err error

	transactionId.Value, err = hexutil.Decode(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), GrpcTimeout)
	defer cancel()

	result, err := g.Client.GetTransactionById(ctx, transactionId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (g *GrpcClient) GetTransactionByInfoId2(id string) (*core.TransactionInfo, error) {
	transactionId := new(api.BytesMessage)
	var err error

	transactionId.Value, err = hexutil.Decode(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), GrpcTimeout)
	defer cancel()

	result, err := g.Client.GetTransactionInfoById(ctx, transactionId)
	if err != nil {
		return nil, err
	}

	return result, nil
}


// 3 steps supported
func (g *GrpcClient) BuildTransaction(fromAddress, toAddress string,
	amount int64) (*core.Transaction, error) {

	transferContract := new(core.TransferContract)
	transferContract.OwnerAddress = base58.DecodeCheck(fromAddress)
	transferContract.ToAddress = base58.DecodeCheck(toAddress)
	transferContract.Amount = amount

	ctx, cancel := context.WithTimeout(context.Background(), GrpcTimeout)
	defer cancel()

	transferTransaction, err := g.Client.CreateTransaction(ctx, transferContract)
	if err != nil {
		return nil, err
	}

	if transferTransaction == nil || len(transferTransaction.GetRawData().GetContract()) == 0 {
		return nil, fmt.Errorf("transfer error: invalid transaction")
	}

	return transferTransaction, nil
}

// sign, see util.SignTransaction2

func (g *GrpcClient) PostTransaction(transferTransaction *core.Transaction) (*api.Return, error){
	ctx, cancel := context.WithTimeout(context.Background(), GrpcTimeout)
	defer cancel()

	result, err := g.Client.BroadcastTransaction(ctx, transferTransaction)
	if err != nil {
		return nil, err
	}

	return result, nil
}