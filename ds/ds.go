package ds

const (
	ResultOK       = 0
	ErrorUnauthorized= iota + 4000
	ErrorUnmarshal
	ErrorAddModel
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
