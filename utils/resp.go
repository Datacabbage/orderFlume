package utils

type OrderStruct struct {
	Orderid  string `json:"orderid"`
	UID      string `json:"uid"`
	Sid      string `json:"sid"`
	Total    string `json:"total"`
	Direct   string `json:"direct"`
	Quantity string `json:"quantity"`
	Dealid   string `json:"dealid"`
	Smstitle string `json:"smstitle"`
	Paytime  string `json:"paytime"`
	Modtime  string `json:"modtime"`
}

type OrderInfo struct {
	Partner    string        `json:"partner"`
	Order_list []OrderStruct `json:"order_list"`
}

//订单回应
type Response struct {
	Is_ok bool   `json:"is_ok"`
	Error string `json:"error"`
}

//自定义返回结构
type Resp struct {
	Code    int
	Message string
	Data    interface{}
}
