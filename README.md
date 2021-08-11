# BeilyPay 支付API（v2） 的 GO 版 SDK

## 一、API 文档

API文档详见 https://github.com/beilypay/beilypay-sdk-java/wiki

## 二、创建 Beilypay 实例

```go
// 加载依赖
import "/beilypay-sdk-go/beilypay"


// 以下对接参数通过商务对接获得
AppID          = <AppID>;
MerchantId     = <MerchantId>;
AppSecret      = <AppSecret>;
Domain         = <Domain>;        // 注意 测试环境 和 生产环境 地址不同

// 创建实例
client := beilypay.BeilypayClient{AppID, MerchantId, AppSecret, Domain}

```

## 三、创建代收订单


```go
//创建代收交易
func CreatePayment(client beilypay.BeilypayClient) {
    pr := beilypay.PaymentReq{}
    pr.NotifyURL = "https://www.baidu.com/"
    pr.FrontCallback = "https://www.baidu.com/"
    pr.OutOrderNo = "outOrderNo"//订单号不能重复
    pr.PayAmount = 500
    pr.Email = "4567546845@qq.com"
    pr.Mobile = "13875025410"
    pr.UserName = "helloName"
    pr.UserID = "uid1"
    payment, err := client.CreatePayment(pr)
    if err != nil {
        fmt.Print(err.Error())
    } else {
        fmt.Print(payment)
    }
}
```

- 返回的是 PaymentReqData 结构体
- API: https://github.com/beilypay/beilypay-sdk-java/wiki#%E4%BB%A3%E4%BB%98%E4%B8%8B%E5%8D%95


## 四、创建代付订单

```go
//创建代付交易
func CreateTrans(client beilypay.BeilypayClient) {
    pr := beilypay.TransReq{}
    pr.Account = "account"
    pr.AccountOwner = "owner"
    pr.AccountType = "Card"
    pr.Address = "浙江省杭州市"
    pr.BankCode = "bankcode"
    pr.Email = "email"
    pr.Ifsc = "ifsc"
    pr.Mobile = "mobile"
    pr.NotifyURL = "https://www.baidu.com"
    pr.OutOrderNo = "outOrderNo" //订单号不能重复
    pr.PayAmount = 500
    payment, err := client.CreateTrans(pr)
    if err != nil {
        fmt.Print(err.Error())
    } else {
        fmt.Print(payment)
    }
}
```
- 返回的是 TransReqData 结构体
- API: https://github.com/beilypay/beilypay-sdk-java/wiki#%E4%BB%A3%E4%BB%98%E4%B8%8B%E5%8D%95


## 五、查询代收订单

```go
func QueryPayment(client beilypay.BeilypayClient) {
    pr := beilypay.PaymentQueryReq{}
    pr.OrderNo = "orderNo"
    payment, err := client.QueryPayment(pr)
    if err != nil {
        fmt.Print(err.Error())
    } else {
        fmt.Print(payment)
    }
}
```
- 返回的是 PaymentQueryReqData 结构体，字段 status 表示订单状态
- API: https://github.com/beilypay/beilypay-sdk-java/wiki#%E4%BB%A3%E6%94%B6%E5%8D%95%E7%8A%B6%E6%80%81%E6%9F%A5%E8%AF%A2


## 五、查询代付订单

```go
func QueryTrans(client beilypay.BeilypayClient) {
    pr := beilypay.TransQueryReq{}
    pr.OrderNo = "orderNo"
    payment, err := client.QueryTrans(pr)
    if err != nil {
        fmt.Print(err.Error())
    } else {
        fmt.Print(payment)
    }
}
```
- 返回的是 TransQueryReqData 结构体，字段 status 表示订单状态
- API: https://github.com/beilypay/beilypay-sdk-java/wiki#%E4%BB%A3%E4%BB%98%E7%8A%B6%E6%80%81%E6%9F%A5%E8%AF%A2


## 六、回调校验 示例

```go
//回调签名校验
beilypay.Verify(data map[string]interface{}, key string) bool 

```
- API: https://github.com/beilypay/beilypay-sdk-java/wiki#%E4%BB%A3%E6%94%B6%E5%BC%82%E6%AD%A5%E9%80%9A%E7%9F%A5

## 七、测试代码

见 sample.go

## 八、其它

无````