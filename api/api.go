package api

import "github.com/QOSGroup/qos-mobile-wallet-sdk/wallet"

const (
	SUCCESS = "success"
	FAIL = "fail"
)

type Response struct {
	Status string `json:"status:omitempty"`
	ErrorMessage string `json:"message:omitempty"`
	Data interface{} `json:"data:omitempty"`	
}

func NewSuccessResponse(data interface{}) Response {
	return Response{
		Status: SUCCESS,
		Data:   data,
	}
}


func NewErrResponse(err error) Response  {
	return Response{
		Status:       FAIL,
		ErrorMessage: err.Error(),
	}
}


type API struct {
	Wallet *wallet.Wallet
}