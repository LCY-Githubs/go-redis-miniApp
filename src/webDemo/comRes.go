package webDemo

type CommonRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessRes(data interface{}) CommonRes {
	return CommonRes{200, "success", data}
}

func NewErrorRes(data interface{}) CommonRes {
	return CommonRes{500, "failed", data}
}
