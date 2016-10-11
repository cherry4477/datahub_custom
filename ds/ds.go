package ds

const (
	ResultOK          = 0
	ErrorUnauthorized = iota + 6000
	ErrorUnmarshal
	ErrorAddModel
	ErrorGetModel
	ErrorUpdateModel
	ErrorDeleteModel
	ErrorAtoi
	ErrorValidateParams
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
