package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/the-great-checkout/transactions-cdc/internal/dto"
)

type TransactionLogClient struct {
	url string
}

func NewTransactionLogClient(url string) *TransactionLogClient {
	return &TransactionLogClient{url: url}
}

func (c *TransactionLogClient) Create(input *dto.TransactionLog) (*dto.TransactionLog, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	uri := c.url + "/v1/transactions/" + input.TransactionID.String() + "/logs"
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader(payload))

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result dto.TransactionLog

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
