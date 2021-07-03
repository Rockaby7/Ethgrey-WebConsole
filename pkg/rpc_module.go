package pkg

type Request struct {
	Msg       string      `json:"msg"`
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
}

type InfoRequest struct {
}

type InfoResponse struct {
}

type QueryRequest struct {
}

type QueryResponse struct {
}
