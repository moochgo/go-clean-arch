package responses

// Standard describes standard rendered JSON for Standard Response
type Standard struct {
	RequestParam string      `json:"request_param"`
	Status       string      `json:"status"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
	Next         string      `json:"next"`
	Version      Version     `json:"version"`
}

// DataTable ...
type DataTable struct {
	Items        interface{} `json:"items"`
	TotalPage    int         `json:"total_pages"`
	TotalItems   int         `json:"total_items"`
	FilteredPage int         `json:"filtered_page"`
	FilteredItem int         `json:"filtered_item"`
}

// Version : describes standard rendered JSON for Apps Versioning
type Version struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
