package beilypay

import (
	"bytes"
	"cool.com/beilypay-sdk-go/util" //根据自己的目录查找
	"crypto/md5"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type BeilypayClient struct {
	AppID      int32  `json:"appId"`
	MerchantId int32  `json:"merchantId"`
	AppSecret  string `json:"appSecret"`
	/**
	 * api域名
	 *
	 * 正式环境 http://service.beilypay.com
	 * 测试环境 http://dev.beilypay.com
	 */
	Domain string `json:"domain"`
}

//createPayment 入参
type PaymentReq struct {
	AppID         int32  `json:"appId,omitempty"`
	Email         string `json:"email,omitempty"`
	FrontCallback string `json:"frontCallback,omitempty"`
	MerchantID    int32  `json:"merchantId,omitempty"`
	Mobile        string `json:"mobile,omitempty"`
	NotifyURL     string `json:"notifyUrl,omitempty"`
	OutOrderNo    string `json:"outOrderNo,omitempty"`
	PayAmount     int32  `json:"payAmount,omitempty"`
	Sign          string `json:"sign,omitempty"`
	UserID        string `json:"userId,omitempty"`
	UserName      string `json:"userName,omitempty"`
}

//queryPayment 入参
type PaymentQueryReq struct {
	AppID      int32  `json:"appId,omitempty"`
	MerchantID int32  `json:"merchantId,omitempty"`
	Sign       string `json:"sign,omitempty"`
	OrderNo    string `json:"orderNo,omitempty"` //订单号不能重复
}

//createTrans 入参
type TransReq struct {
	Account      string `json:"account,omitempty"`
	AccountOwner string `json:"accountOwner,omitempty"`
	AccountType  string `json:"accountType,omitempty"`
	Address      string `json:"address,omitempty"`
	AppID        int32  `json:"appId,omitempty"`
	BankCode     string `json:"bankCode,omitempty"`
	Email        string `json:"email,omitempty"`
	Ifsc         string `json:"ifsc,omitempty"`
	MerchantID   int32  `json:"merchantId,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	NotifyURL    string `json:"notifyUrl,omitempty"`
	OutOrderNo   string `json:"outOrderNo,omitempty"` //订单号不能重复
	PayAmount    int32  `json:"payAmount,omitempty"`
	Sign         string `json:"sign,omitempty"`
}

//queryTrans
type TransQueryReq struct {
	AppID      int32  `json:"appId,omitempty"`
	MerchantID int32  `json:"merchantId,omitempty"`
	Sign       string `json:"sign,omitempty"`
	OrderNo    string `json:"orderNo,omitempty"`
}

type Response struct {
	Code int32       `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type PaymentReqData struct {
	OutOrderNo string `json:"outOrderNo"`
	PayAmount  int32  `json:"payAmount"`
	PayURL     string `json:"payUrl"`
	Sign       string `json:"sign"`
	OrderNo    string `json:"orderNo"`
}

type PaymentQueryReqData struct {
	OutOrderNo string `json:"outOrderNo"`
	Paid       int32  `json:"paid"`
	PayAmount  int32  `json:"payAmount"`
	Sign       string `json:"sign"`
	Status     int32  `json:"status"`
	OrderNo    string `json:"orderNo"`
	TransTime  string `json:"transTime"`
}

type TransReqData struct {
	AppID      int32  `json:"appId"`
	MerchantID int32  `json:"merchantId"`
	OutOrderNo string `json:"outOrderNo"`
	PayAmount  int32  `json:"payAmount"`
	Sign       string `json:"sign"`
	Status     int32  `json:"status"`
	OrderNo    string `json:"orderNo"`
	TransTime  string `json:"transTime"`
}

type TransQueryReqData struct {
	AppID      int32  `json:"appId"`
	MerchantID int32  `json:"merchantId"`
	OutOrderNo string `json:"outOrderNo"`
	PayAmount  int32  `json:"payAmount"`
	Sign       string `json:"sign"`
	Status     int32  `json:"status"`
	OrderNo    string `json:"orderNo"`
}

//创建代收交易

func (bc *BeilypayClient) CreatePayment(pr PaymentReq) (*PaymentReqData, error) {
	var err error
	url := bc.Domain + CREATE_PAYMENT_URL
	pr.AppID = bc.AppID
	pr.MerchantID = bc.MerchantId
	sign := GenerateSignature(JSONToMap(pr), bc.AppSecret)
	pr.Sign = sign
	data, err := util.Post(url, pr, ContentType)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	response := Response{}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	if response.Code != 200 {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  response.Msg,
		}
	}
	rd := PaymentReqData{}
	b, err := json.Marshal(response.Data)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	err = json.Unmarshal(b, &rd)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	return &rd, nil
}

//代收交易查询

func (bc *BeilypayClient) QueryPayment(pqr PaymentQueryReq) (*PaymentQueryReqData, error) {
	var err error
	url := bc.Domain + PAYMENT_QUERY_URL
	pqr.AppID = bc.AppID
	pqr.MerchantID = bc.MerchantId
	sign := GenerateSignature(JSONToMap(pqr), bc.AppSecret)
	pqr.Sign = sign
	data, err := util.Post(url, pqr, ContentType)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	response := Response{}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	if response.Code != 200 {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  response.Msg,
		}
	}
	rd := PaymentQueryReqData{}
	b, err := json.Marshal(response.Data)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	err = json.Unmarshal(b, &rd)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	return &rd, nil
}

//创建代付交易

func (bc *BeilypayClient) CreateTrans(tr TransReq) (*TransReqData, error) {
	var err error
	url := bc.Domain + CREATE_TRANS_URL
	tr.AppID = bc.AppID
	tr.MerchantID = bc.MerchantId
	sign := GenerateSignature(JSONToMap(tr), bc.AppSecret)
	tr.Sign = sign
	data, err := util.Post(url, tr, ContentType)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	response := Response{}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	if response.Code != 200 {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  response.Msg,
		}
	}
	rd := TransReqData{}
	b, err := json.Marshal(response.Data)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	err = json.Unmarshal(b, &rd)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	return &rd, nil
}

//代付交易查询

func (bc *BeilypayClient) QueryTrans(tqr TransQueryReq) (*TransQueryReqData, error) {
	var err error
	url := bc.Domain + TRANS_QUERY_URL
	tqr.AppID = bc.AppID
	tqr.MerchantID = bc.MerchantId
	sign := GenerateSignature(JSONToMap(tqr), bc.AppSecret)
	tqr.Sign = sign
	data, err := util.Post(url, tqr, ContentType)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	response := Response{}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	if response.Code != 200 {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  response.Msg,
		}
	}
	rd := TransQueryReqData{}
	b, err := json.Marshal(response.Data)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	err = json.Unmarshal(b, &rd)
	if err != nil {
		return nil, &BeilypayError{
			ErrorCode: 500,
			ErrorMsg:  err.Error(),
		}
	}
	return &rd, nil
}

//签名相关

type Values map[string]string

const SIGN = "sign"

//签名规则
//第一步，设所有发送或者接收到的数据为集合M，将集合M内非空参数值的参数按照参数名ASCII码从小到大排序（字典序），使用URL键值对的格式（key1=value1&key2=value2…）拼接成字符串stringA。 特别注意以下重要规则：

//参数名ASCII码从小到大排序（字典序）；
//参数名区分大小写；
//如果参数的值为空不参与签名；
//验证调用返回或平台主动通知签名时，传送的sign参数不参与签名，将生成的签名与该sign值作校验。
//接口可能增加字段，验证签名时必须支持增加的扩展字段。
//第二步，在stringA最后拼接上key得到stringSignTemp字符串，并对stringSignTemp进行MD5运算，再将得到的字符串所有字符转换为大写，得到sign值signValue。​

//待签名字符串：stringSignTemp= a=123&b=123&c=123&key=key 签名: sign = MD5(stringSignTemp) .toUpperCase();

func GenerateSignature(data map[string]interface{}, key string) string {
	return strings.ToUpper(JoinStringsInASCII(data, "&", key, SIGN))
}

//验证签名

func Verify(data map[string]interface{}, key string) bool {
	sign := data[SIGN]
	signature := GenerateSignature(data, key)
	return reflect.DeepEqual(sign, signature)
}

//JoinStringsInASCII 按照规则，参数名ASCII码从小到大排序后拼接
//data 待拼接的数据
//sep 连接符
//exceptKeys 被排除的参数名，不参与排序及拼接
func JoinStringsInASCII(data map[string]interface{}, sep string, key string, exceptKeys ...string) string {
	var list []string
	m := make(map[string]int)
	if len(exceptKeys) > 0 {
		for _, except := range exceptKeys {
			m[except] = 1
		}
	}
	for k := range data {
		if _, ok := m[k]; ok {
			continue
		}
		value := data[k]
		if value == "" || value == nil {
			continue
		}
		list = append(list, fmt.Sprintf("%s=%s", k, value))
	}
	sort.Strings(list)
	list = append(list, fmt.Sprintf("%s=%s", "key", key))
	src := strings.Join(list, sep)
	sum := md5.Sum([]byte(src))
	return fmt.Sprintf("%x", sum)
}

// json转map函数，通用

func JSONToMap(v interface{}) map[string]interface{} {
	prb, _ := json.Marshal(v)

	var personFromJSON interface{}
	decoder := json.NewDecoder(bytes.NewReader(prb))
	decoder.UseNumber()
	decoder.Decode(&personFromJSON)
	r := personFromJSON.(map[string]interface{})
	return r
}
