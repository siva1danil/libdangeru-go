package core

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ClientAPI represents an API client for interacting with the danger/u/ server.
type ClientAPI struct {
	Address string
	Http    http.Client
}

// BoardDetails represents the details of a board.
type BoardDetails struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Rules string `json:"rules"`
}

// Post represents a post on a thread or reply.
type Post struct {
	Post_ID           uint   `json:"post_id"`           // Thread & Reply
	Board             string `json:"board"`             // Thread & Reply
	Comment           string `json:"comment"`           // Thread & Reply
	Date_Posted       uint   `json:"date_posted"`       // Thread & Reply
	Hash              string `json:"hash"`              // Thread & Reply
	IP                string `json:"ip"`                // Thread & Reply; auth-only
	Capcode           string `json:"capcode"`           // Thread & Reply; auth-only
	Is_OP             bool   `json:"is_op"`             // Thread & Reply
	Parent            uint   `json:"parent"`            // Reply
	Title             string `json:"title"`             // Thread
	Last_Bumped       uint   `json:"last_bumped"`       // Thread
	Is_Locked         bool   `json:"is_locked"`         // Thread
	Number_Of_Replies uint   `json:"number_of_replies"` // Thread
	Sticky            bool   `json:"sticky"`            // Thread
	Stickyness        uint   `json:"stickyness"`        // Thread
}

func (client *ClientAPI) init() {
	if client.Address == "" {
		client.Address = ADDRESS
	}
}

// Create a new thread.
//
// Route: /post
func (client *ClientAPI) PostThread(board string, title string, comment string, capcode bool) error {
	path := "/post"
	form := url.Values{}
	form.Add("board", board)
	form.Add("title", title)
	form.Add("comment", comment)
	if capcode {
		form.Add("capcode", "true")
	}

	client.init()
	_, _, err := PostRequest(client.Http, client.Address+"/"+path, form)
	if err != nil {
		return err
	}

	return nil
}

// Create a new reply.
//
// Route: /reply
func (client *ClientAPI) PostReply(board string, parent uint, content string, capcode bool) (int, error) {
	path := PATH_API_POST_REPLY
	form := url.Values{}
	form.Add("board", board)
	form.Add("parent", strconv.Itoa(int(parent)))
	form.Add("content", content)
	if capcode {
		form.Add("capcode", "true")
	}

	client.init()
	body, _, err := PostRequest(client.Http, client.Address+"/"+path, form)
	if err != nil {
		return 0, err
	}
	result_str, ok := strings.CutPrefix(string(body), "OK/")
	if !ok {
		return 0, fmt.Errorf("expected OK, got %s", string(body))
	}
	result, err := strconv.Atoi(result_str)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// Get all available boards. /all/ is filtered.
//
// Route: /api/v2/boards
func (client *ClientAPI) Boards() ([]string, error) {
	result := []string{}
	path := PATH_API_BOARDS

	client.init()
	body, _, err := GetRequestJSON[[]string](client.Http, client.Address+"/"+path)
	if err != nil {
		return result, err
	}
	for i := 0; i < len(body); i++ {
		if body[i] != "all" {
			result = append(result, body[i])
		}
	}

	return result, nil
}

// Get details for a board.
//
// Route: /api/v2/board/$board$/detail
func (client *ClientAPI) BoardDetails(board string) (BoardDetails, error) {
	path := fmt.Sprintf(PATH_API_BOARD_DETAILS, board)

	client.init()
	body, _, err := GetRequestJSON[BoardDetails](client.Http, client.Address+"/"+path)
	if err != nil {
		return body, err
	}

	return body, nil
}

// Get active threads for a board. First page is 0.
//
// Route: /api/v2/board/$board$?page=$page$
func (client *ClientAPI) Threads(board string, page uint) ([]Post, error) {
	path := fmt.Sprintf(PATH_API_THREADS, board, page)

	client.init()
	body, _, err := GetRequestJSON[[]Post](client.Http, client.Address+"/"+path)
	if err != nil {
		return body, err
	}

	return body, nil
}

// Get active threads for all boards. First page is 0.
//
// Route: /api/v2/board/all?page=$page$
func (client *ClientAPI) ThreadsAll(page uint) ([]Post, error) {
	return client.Threads("all", page)
}

// Get the metadata for a thread.
//
// Route: /api/v2/thread/$thread$/metadata
func (client *ClientAPI) ThreadMetadata(id uint) (Post, error) {
	path := fmt.Sprintf(PATH_API_THREAD_METADATA, id)

	client.init()
	body, _, err := GetRequestJSON[Post](client.Http, client.Address+"/"+path)
	if err != nil {
		return body, err
	}

	return body, nil
}

// Get all replies for a thread. Metadata is currently not parsed.
//
// Route: /api/v2/thread/$thread$/replies
func (client *ClientAPI) ThreadReplies(id uint) ([]Post, error) {
	path := fmt.Sprintf(PATH_API_THREAD_REPLIES, id)

	client.init()
	body, _, err := GetRequestJSON[[]Post](client.Http, client.Address+"/"+path)
	if err != nil {
		return body, err
	}

	return body, nil
}
