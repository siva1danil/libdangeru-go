package libdangeru

import (
	"io"
	"net/http"
)

type ClientWeb struct {
	addr   Addr
	client http.Client
	debug  bool
}

// Create a new dangeru_client_web with options.
func NewClientWeb(options *ClientOptions) ClientWeb {
	client_web := ClientWeb{
		addr:   options.Addr,
		client: options.Client,
		debug:  options.Debug,
	}

	return client_web
}

func (client *ClientWeb) get(path string) ([]byte, error) {
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
