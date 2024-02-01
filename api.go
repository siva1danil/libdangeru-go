package libdangeru

import (
	"io"
	"net/http"
)

type ClientAPI struct {
	addr   Addr
	client http.Client
	debug  bool
}

// Create a new dangeru_client_api with options.
func NewClientAPI(options *ClientOptions) ClientAPI {
	client_api := ClientAPI{
		addr:   options.Addr,
		client: options.Client,
		debug:  options.Debug,
	}

	return client_api
}

func (client *ClientAPI) get(path string) ([]byte, error) {
	url := client.addr.Scheme + "://" + client.addr.Domain + "/" + path

	req, err := client.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
