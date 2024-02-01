package libdangeru

import (
	"encoding/json"
	"fmt"
)

type res_boards []string
type res_board_details struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Rules string `json:"rules"`
}
type post struct {
	Post_ID           uint   `json:"post_id"`           // Thread & Reply
	Board             string `json:"board"`             // Thread & Reply
	Comment           string `json:"comment"`           // Thread & Reply
	Date_Posted       uint   `json:"date_posted"`       // Thread & Reply
	Hash              string `json:"hash"`              // Thread & Reply
	IP                string `json:"ip"`                // Thread & Reply; auth-only
	Capcode           string `json:"capcode"`           // Thread & Reply; auth-only
	Is_OP             bool   `json:"is_op"`             // Is the post thread or reply?
	Parent            uint   `json:"parent"`            // Reply
	Title             string `json:"title"`             // Thread
	Last_Bumped       uint   `json:"last_bumped"`       // Thread
	Is_Locked         bool   `json:"is_locked"`         // Thread
	Number_Of_Replies uint   `json:"number_of_replies"` // Thread
	Sticky            bool   `json:"sticky"`            // Thread
	Stickyness        uint   `json:"stickyness"`        // Thread
}
type res_threads []post
type res_replies []post

// Get all available boards. /all/ is filtered.
//
// Route: /api/v2/boards
func (client *ClientAPI) Boards() (res_boards, error) {
	res := res_boards{}
	path := client.addr.PathBoards
	data, err := client.get(path)
	if err != nil {
		return res, err
	}

	if client.debug {
		fmt.Println(string(data))
	}

	tmp := []string{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return res, err
	}

	for i := 0; i < len(tmp); i++ {
		if tmp[i] != "all" {
			res = append(res, tmp[i])
		}
	}

	return res, nil
}

// Get details for a board.
//
// Route: /api/v2/board/$board$/detail
func (client *ClientAPI) BoardDetails(board string) (res_board_details, error) {
	res := res_board_details{}
	path := fmt.Sprintf(client.addr.PathBoardDetails, board)
	data, err := client.get(path)
	if err != nil {
		return res, err
	}

	if client.debug {
		fmt.Println(string(data))
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Get active threads for a board. First page is 0.
//
// Route: /api/v2/board/$board$?page=$page$
func (client *ClientAPI) Threads(board string, page uint) (res_threads, error) {
	res := res_threads{}
	path := fmt.Sprintf(client.addr.PathThreads, board, page)
	data, err := client.get(path)
	if err != nil {
		return res, err
	}

	if client.debug {
		fmt.Println(string(data))
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Get the metadata for a thread.
//
// Route: /api/v2/thread/$thread$/metadata
func (client *ClientAPI) ThreadMetadata(id uint) (post, error) {
	res := post{}
	path := fmt.Sprintf(client.addr.PathThreadMetadata, id)
	data, err := client.get(path)
	if err != nil {
		return res, err
	}

	if client.debug {
		fmt.Println(string(data))
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Get all replies for a thread. Metadata is currently not parsed.
//
// Route: /api/v2/thread/$thread$/replies
func (client *ClientAPI) ThreadReplies(id uint) (res_replies, error) {
	res := res_replies{}
	path := fmt.Sprintf(client.addr.PathThreadReplies, id)
	data, err := client.get(path)
	if err != nil {
		return res, err
	}

	if client.debug {
		fmt.Println(string(data))
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
