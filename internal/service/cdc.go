package service

import (
	"encoding/json"

	"github.com/the-great-checkout/transactions-cdc/internal/dto"
)

type TransactionLogClient interface {
	Create(input *dto.TransactionLog) (output *dto.TransactionLog, err error)
}

type CDCService struct {
	client TransactionLogClient
}

func NewCDCService(client TransactionLogClient) *CDCService {
	return &CDCService{client: client}
}

func (service *CDCService) ChangeDataCapture(bytes []byte) error {
	var transaction dto.Transaction
	err := json.Unmarshal(bytes, &transaction)
	if err != nil {
		return err
	}

	var transactionLog = dto.TransactionLog{
		TransactionID: transaction.ID,
		Timestamp:     transaction.UpdatedAt,
		Status:        transaction.Status,
		Value:         transaction.Value,
	}

	_, err = service.client.Create(&transactionLog)
	if err != nil {
		return err
	}

	return nil
}
