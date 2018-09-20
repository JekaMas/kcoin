package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/kowala-tech/kcoin/e2e/cluster"
	"github.com/kowala-tech/kcoin/mock-exchange/app"
)

type MockExchangeContext struct {
	globalCtx *Context
	mockAddr  string
}

func NewMockExchangeContext(parentCtx *Context) *MockExchangeContext {
	ctx := &MockExchangeContext{
		globalCtx: parentCtx,
	}
	return ctx
}

func (ctx *MockExchangeContext) Reset() {
}

func (ctx *MockExchangeContext) TheMockExchangeIsRunning() error {
	mockAddr, err := ctx.runMockExchange()
	if err != nil {
		return fmt.Errorf("error executing mock exchange: %s", err)
	}

	ctx.mockAddr = mockAddr

	return nil
}

func (ctx *MockExchangeContext) runMockExchange() (mockAddr string, err error) {
	mockSpec, err := cluster.MockExchangeSpec(ctx.globalCtx.nodeSuffix)
	if err != nil {
		return "", err
	}
	if err := ctx.globalCtx.nodeRunner.Run(mockSpec); err != nil {
		return "", err
	}
	mockIP, err := ctx.globalCtx.nodeRunner.IP(mockSpec.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v:8080", mockIP), nil
}

func (ctx *MockExchangeContext) IFetchTheExchangeWithMockData(accountsDataTable *gherkin.DataTable) error {
	request, err := getRateRequestFromTableData(accountsDataTable)
	if err != nil {
		return err
	}

	postData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to encode table data: %s", err)
	}

	response, err := http.Post(
		fmt.Sprintf("http://%s:9080/api/fetch", ctx.globalCtx.nodeRunner.HostIP()),
		"",
		bytes.NewReader(postData),
	)
	if err != nil {
		return fmt.Errorf("failed to send post request to exchange: %s", err)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("fetching the mocked exchange server returned status code %d", response.StatusCode)
	}

	return nil
}

func getRateRequestFromTableData(accountsDataTable *gherkin.DataTable) (*app.Request, error) {
	request := app.Request{}
	for i, row := range accountsDataTable.Rows {
		if i != 0 {
			typeVal := row.Cells[0].Value
			if typeVal == "buy" {
				amount, rate, err := extractAmountAndRateFromRow(row)
				if err != nil {
					return nil, err
				}

				request.Buy = append(request.Buy, app.RateValue{Amount: amount, Rate: rate})
			} else if typeVal == "sell" {
				amount, rate, err := extractAmountAndRateFromRow(row)
				if err != nil {
					return nil, err
				}

				request.Sell = append(request.Sell, app.RateValue{Amount: amount, Rate: rate})
			} else {
				continue
			}
		}
	}

	return &request, nil
}

func extractAmountAndRateFromRow(row *gherkin.TableRow) (amount float64, rate float64, err error) {
	amount, err = strconv.ParseFloat(row.Cells[1].Value, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid amount data in table for exchange mock server: %s", err)
	}

	rate, err = strconv.ParseFloat(row.Cells[2].Value, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid rate data in table for exchange mock server: %s", err)
	}

	return amount, rate, nil
}

func (ctx *MockExchangeContext) ICanQueryTheExchangeAndGetMockedResponse(exchangeName string, jsonResponse *gherkin.DocString) error {
	expectedResponse := jsonResponse.Content

	response, err := http.Get(
		fmt.Sprintf("http://%s:9080/api/%s/get", ctx.globalCtx.nodeRunner.HostIP(), exchangeName),
	)
	if err != nil {
		return fmt.Errorf("there was a problem getting mocked data from server: %s", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading body from request: %s", err)
	}

	if expectedResponse != string(responseBytes) {
		return fmt.Errorf("failed to assert equal response when getting mocked data: \nExpected: %s\nGot: %s", expectedResponse, responseBytes)
	}

	return nil
}
