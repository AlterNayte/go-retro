package httpbin

type PostValue struct {
	Args struct {
	} `json:"args"`
	Data  string `json:"data"`
	Files struct {
	} `json:"files"`
	Form struct {
	} `json:"form"`
	Headers struct {
		Accept         string `json:"Accept"`
		AcceptEncoding string `json:"Accept-Encoding"`
		ContentLength  string `json:"Content-Length"`
		ContentType    string `json:"Content-Type"`
		Host           string `json:"Host"`
		PostmanToken   string `json:"Postman-Token"`
		UserAgent      string `json:"User-Agent"`
		XAmznTraceId   string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	Json struct {
		Boots int `json:"boots"`
	} `json:"json"`
	Origin string `json:"origin"`
	Url    string `json:"url"`
}

type PostInput struct {
	Boots int `json:"boots"`
}

type HttpAPI struct {
	PostExample func(input PostInput) (*PostValue, error) `method:"POST" path:"/post"`
}
