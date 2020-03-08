package restClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enableMocks = false
	mocks       = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

func getMockId(httpMethod string, url string) string{
	return fmt.Sprintf("&s &s", httpMethod, url)
}

func StartMockUps() {
	enableMocks = true
}

func StopMockUps() {
	enableMocks = false
}

func FlushMockUps()  {
	mocks = make(map[string]*Mock)
}

func AddMockUp(mock Mock) {
	mocks[getMockId(mock.HttpMethod, mock.Url)]= &mock
}

func Post(url string, body interface{}, header http.Header) (*http.Response, error) {

	if enableMocks {
		// return local mock without calling any external resource
		mock := mocks[getMockId(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockUp found for given request")
		}
		return mock.Response, mock.Err
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		err.Error()
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = header

	client := http.Client{}
	return client.Do(request)
}
