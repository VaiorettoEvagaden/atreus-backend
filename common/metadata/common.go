package metadata

import "github.com/DayDayYiDay/atreus-backend/common"

type BaseResp struct {
	Result bool   `json:"result"`
	Code   int    `json:"bk_error_code"`
	ErrMsg string `json:"bk_error_msg"`
}

type Response struct {
	BaseResp `json:",inline"`
	Data     interface{} `json:"data"`
}

func NewSuccessResp(data interface{}) *Response {
	return &Response{
		BaseResp: BaseResp{true, common.CCSuccess, common.CCSuccessStr},
		Data:     data,
	}
}