package apimodel

type BackendResponseList struct {
	BackendResponseInfoArray []BackendResponseInfoArray `json:"backendResponseInfoArray,omitempty"`
	Size                     int                        `json:"size,omitempty"`
}

type BackendResponseInfoArray struct {
	APIName   string `json:"apiName,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
	Message   string `json:"message,omitempty"`
	System    string `json:"system,omitempty"`
	URL       string `json:"url,omitempty"`
}
