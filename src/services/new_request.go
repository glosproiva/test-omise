package services

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"inception.com/common"
	"inception.com/configuration"
)

type response_new_request struct {
	Id     int64   `json:"transaction"`
	Amount float64 `json:"amount"`
}

func CreateTransaction(c echo.Context) (err error) {
	var res bool
	var amount float64

	now := time.Now().Format("2006/01/02 15:04:05")

	str_value := c.FormValue("amount")
	amount, err = strconv.ParseFloat(str_value, 64)
	if err != nil {
		c.JSON(common.HTTP_SUCCESS, common.Responsevalue{
			Result:  -1,
			Message: "Please select amount",
			Value:   nil,
		})
		return

	}

	client, e := omise.NewClient(configuration.OmisePublicKey, configuration.OmiseSecretKey)
	res = common.IsError(e)
	if res {
		c.JSON(common.HTTP_SUCCESS, common.Responsevalue{
			Result:  -1,
			Message: "request failed",
			Value:   nil,
		})
		return
	}

	source, createSource := &omise.Source{}, &operations.CreateSource{
		Amount:   common.FloadAmount2Int(amount),
		Currency: "thb",
		Type:     "internet_banking_scb",
	}

	if e := client.Do(source, createSource); e != nil {
		res = common.IsError(e)
		c.JSON(common.HTTP_SUCCESS, common.Responsevalue{
			Result:  -2,
			Message: "request failed",
			Value:   nil,
		})
		return
	}

	// log.Printf("created source: %#v\n", source)

	strsql := "INSERT INTO payment_transaction(transaction_id,amount,type,currentcy,timestamp) VALUES(?,?,?,?,?)"
	stmt, err := common.DB.Prepare(strsql)
	res = common.IsError(err)
	if res {
		c.JSON(common.HTTP_SUCCESS, common.Responsevalue{
			Result:  -3,
			Message: "request failed",
			Value:   nil,
		})
		return
	}

	defer stmt.Close()

	result, err := stmt.Exec(source.ID, source.Amount, source.Type, source.Currency, now)
	res = common.IsError(err)
	if res {
		c.JSON(common.HTTP_SUCCESS, common.Responsevalue{
			Result:  -4,
			Message: "request failed",
			Value:   nil,
		})
		return
	}

	insertID, err := result.LastInsertId()
	res = common.IsError(err)
	if res {
		c.JSON(common.HTTP_SUCCESS, common.Responsevalue{
			Result:  -4,
			Message: "request failed",
			Value:   nil,
		})
		return
	}

	c.JSON(common.HTTP_SUCCESS, common.Responsevalue{
		Result:  1,
		Message: "successfully",
		Value: response_new_request{
			Id:     insertID,
			Amount: common.CalculateAmount(source.Amount),
		},
	})
	return
}
