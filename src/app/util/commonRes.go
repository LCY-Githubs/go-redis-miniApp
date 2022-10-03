package util

//返回给前端的基础数据
type Response struct {
	Code   int32  `json:"code"`   //状态码
	Msg    string `json:"msg"`    //返回的消息文本
	Status string `json:"status"` //请求的状态
}

//如果返回有数据的时候，需要加上Data域用于存放数据
type response struct {
	Response
	Data interface{} `json:"data"`
}

func (r *response) SetData(data interface{}) {
	r.Data = data
}

func (r *response) SetMsg(msg string) {
	r.Msg = msg
}

func (r *response) SetCode(code int32) {
	r.Code = code
}

func (r *response) SetSuccess(success bool) {
	if !success {
		r.Status = "error"
	}
}

//提供拷贝功能，用于提供response的副本
func (r response) Clone() *response {
	return &r
}
