package libdangeru

import (
	"io"
	"net/http"
)

type dangeru_client_web struct {
	addr   dangeru_addr
	client http.Client
	debug  bool
}

// Create a new dangeru_client_web with options.
func NewClientWeb(options *dangeru_options) dangeru_client_web {
	client_web := dangeru_client_web{
		addr:   options.Addr,
		client: options.Client,
		debug:  options.Debug,
	}

	return client_web
}

func (client *dangeru_client_web) get(path string) ([]byte, error) {
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
