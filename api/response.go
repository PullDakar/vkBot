package api

type ClientResponse struct {
	StatusCode int
	Content    string
	Headers    map[string]string
}

func NewClientResponse(statusCode int, content string, headers map[string]string) *ClientResponse {
	clientResponse := new(ClientResponse)

	clientResponse.StatusCode = statusCode
	clientResponse.Content = content
	clientResponse.Headers = headers

	return clientResponse
}
