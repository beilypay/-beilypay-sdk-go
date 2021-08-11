package main

import (
	"fmt"
	"github.com/beilypay/beilypay-sdk-go/beilypay"
)

func main() {
	//appId merchantId appSecret 需要换成自己的
	client := beilypay.BeilypayClient{111, 111, "***", beilypay.DEV_DOMAIN}
	CreatePayment(client)
	//QueryPayment(client)
	//CreateTrans(client)
	//QueryTrans(client)
}

//创建代收交易
func CreatePayment(client beilypay.BeilypayClient) {
	pr := beilypay.PaymentReq{}
	pr.NotifyURL = "https://www.baidu.com"
	pr.FrontCallback = "https://www.aaa.com"
	pr.OutOrderNo = "123456a4a8a6aaa26"
	pr.PayAmount = 500
	pr.Email = "4567546845@qq.com"
	pr.Mobile = "1385458454"
	pr.UserName = "sdaf"
	pr.UserID = "fsad"
	payment, err := client.CreatePayment(pr)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(payment)
	}
}

//代收交易查询
func QueryPayment(client beilypay.BeilypayClient) {
	pr := beilypay.PaymentQueryReq{}
	pr.OrderNo = "210811939***4"
	payment, err := client.QueryPayment(pr)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(payment)
	}
}

//创建代付交易
func CreateTrans(client beilypay.BeilypayClient) {
	pr := beilypay.TransReq{}
	pr.Account = "sdafasd"
	pr.AccountOwner = "lisi"
	pr.AccountType = "Card"
	pr.Address = "浙江省杭州市"
	pr.BankCode = "5345"
	pr.Email = "54325@qq.com"
	pr.Ifsc = "4325a"
	pr.Mobile = "13856856540"
	pr.NotifyURL = "https://www.baidu.com"
	pr.OutOrderNo = "45645***23123a3" //
	pr.PayAmount = 500
	payment, err := client.CreateTrans(pr)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(payment)
	}

}

//代付交易查询
func QueryTrans(client beilypay.BeilypayClient) {
	pr := beilypay.TransQueryReq{}
	pr.OrderNo = "2108***75153152"
	payment, err := client.QueryTrans(pr)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(payment)
	}
}
