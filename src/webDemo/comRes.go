package webDemo

type CommonRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessRes(data interface{}) CommonRes {
	a := CommonRes{200, "success", data}
	return a
}
