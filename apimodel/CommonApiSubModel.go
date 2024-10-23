package apimodel

type KeyValueInfo struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}
