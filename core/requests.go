package core

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func GetRequest(client http.Client, url string) ([]byte, http.Header, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.Header, err
	}

	return body, resp.Header, nil
}

func GetRequestJSON[T interface{}](client http.Client, url string) (T, http.Header, error) {
	var result T

	body, header, err := GetRequest(client, url)
	if err != nil {
		return result, nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, header, err
	}

	return result, header, nil
}

func PostRequest(client http.Client, url string, form url.Values) ([]byte, http.Header, error) {
	resp, err := client.PostForm(url, form)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, resp.Header, nil
}
