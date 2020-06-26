package nbxplorer

// Response struct
type Response struct {
	Success        bool   `json:"success"`
	RPCCode        int    `json:"rpcCode"`
	RPCCodeMessage string `json:"rpcCodeMessage"`
	RPCMessage     string `json:"rpcMessage"`
}

// ErrorResponse struct
type ErrorResponse struct {
	HTTPCode int    `json:"httpCode"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}
