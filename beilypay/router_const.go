package beilypay

const (
	ContentType = "application/json"
)

const (
	CREATE_PAYMENT_URL = "/v2/payment/create"

	PAYMENT_QUERY_URL = "/v2/payment/query"

	CREATE_TRANS_URL = "/v2/trans/create"

	TRANS_QUERY_URL = "/v2/trans/query"
)

const (
	DEV_DOMAIN  = "http://dev.beilypay.com"     //测试环境
	PROD_DOMAIN = "http://service.beilypay.com" //生产环境
)
