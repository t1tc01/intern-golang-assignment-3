package request

type ApiReq struct {
	ID          int
	ReqTime     int64                  `json:"time"`
	ReqParams   map[string]interface{} `json:"req_param"`
	ReqBody     map[string]interface{} `json:"req_body"`
	ReqHeaders  map[string]interface{} `json:"req_headers"`
	RespTime    int64                  `json:"resp_time"`
	RespStatus  int                    `json:"resp_status"`
	RespBody    map[string]interface{} `json:"resp_body"`
	RespHeaders map[string]interface{} `json:"resp_headers"`
	CreatedAt   int64                  `json:"createdTime"`
	UpdatedAt   int64                  `json:"updatedTime"`
	DeletedAt   int64                  `json:"deletedTime"`
}
