package common

type Responsevalue struct {
	Result  int         `json:"result"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}
