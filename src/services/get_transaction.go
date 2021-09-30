package services

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"inception.com/common"
)

type response_get_transction struct {
	Id             int64   `json:"id"`
	Transaction_id string  `json:"token"`
	Amount         float64 `json:"amount"`
	Mtype          string  `json:"type"`
	Currentcy      string  `json:"currentcy"`
	Timestamp      string  `json:"timestamp"`
}

func GetTransaction(c echo.Context) (err error) {
	var res bool
	var bindings []interface{}
	var transaction_id int64

	str_id := c.FormValue("transaction_id")
	transaction_id, err = strconv.ParseInt(str_id, 10, 64)
	if err != nil {
		transaction_id = -1
		c.JSON(common.HTTP_SUCCESS, common.Responsevalue{
			Result:  -1,
			Message: "Please select payment",
			Value:   nil,
		})

	}

	fmt.Println("test transaction_id : ", transaction_id)

	strsql := "SELECT id,transaction_id,amount,type,currentcy,timestamp FROM payment_transaction "
	strwhere := ""

	bindings = make([]interface{}, 0)
	if transaction_id != -1 {
		if strwhere != "" {
			strwhere += " AND "
		} else {
			strwhere += " WHERE "
		}

		strwhere += " id = ? "
		bindings = append(bindings, transaction_id)
	}

	strsql += strwhere
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

	rows, err := stmt.Query(bindings...)
	res = common.IsError(err)
	if res {
		c.JSON(common.HTTP_SUCCESS, common.Responsevalue{
			Result:  -4,
			Message: "request failed",
			Value:   nil,
		})
		return
	}

	rlist := []response_get_transction{}
	for rows.Next() {
		r := response_get_transction{}
		err = rows.Scan(&r.Id, &r.Transaction_id, &r.Amount, &r.Mtype, &r.Currentcy, &r.Timestamp)
		res = common.IsError(err)
		if res {
			c.JSON(common.HTTP_SUCCESS, common.Responsevalue{
				Result:  -5,
				Message: "request failed",
				Value:   nil,
			})
			return
		}

		r.Amount = common.CalculateAmount(int64(r.Amount))

		rlist = append(rlist, r)

	}

	r := common.Responsevalue{}

	r.Result = 1
	r.Message = "get transaction"
	r.Value = rlist
	c.JSON(common.HTTP_SUCCESS, r)
	return
}
