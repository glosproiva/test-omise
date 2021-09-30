package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"inception.com/common"
	"inception.com/services"
)

const (
	listening_port = "25671"
	db_name        = "./transaction.db"
)

var (
	res bool
)

func main() {

	fmt.Println("start program")
	var err error
	common.DB, err = sql.Open("sqlite3", db_name)
	res = common.IsError(err)
	if res {
		return
	}

	defer common.DB.Close()

	var (
		e = echo.New()
	)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	g := e.Group("/v1")
	g.POST("/transaction/create", services.CreateTransaction)
	g.POST("/transaction/get", services.GetTransaction)

	log.Fatal(e.Start(":" + listening_port))

	// client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	// if e != nil {
	// 	log.Fatal(e)
	// }

	// source, createSource := &omise.Source{}, &operations.CreateSource{
	// 	Amount:   100000,
	// 	Currency: "thb",
	// 	Type:     "internet_banking_scb",
	// }

	// if e := client.Do(source, createSource); e != nil {
	// 	log.Fatalln(e)
	// }

	// log.Printf("created source: %#v\n", source)

	// token := "tokn_xxxxxxxxxxxxx"

	// // Creates a charge from the token
	// charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
	// 	Amount:   100000, // ฿ 1,000.00
	// 	Currency: "thb",
	// 	Card:     token,
	// }
	// if e := client.Do(charge, createCharge); e != nil {
	// 	log.Fatal(e)
	// }

	// fmt.Println("test2")
	// log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)
}
