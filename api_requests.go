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
type thread_metadata struct {
	Post_ID           uint   `json:"post_id"`
	Board             string `json:"board"`
	Is_OP             bool   `json:"is_op"`
	Comment           string `json:"comment"`
	Date_Posted       uint   `json:"date_posted"`
	Title             string `json:"title"`
	Last_Bumped       uint   `json:"last_bumped"`
	Is_Locked         bool   `json:"is_locked"`
	Number_Of_Replies uint   `json:"number_of_replies"`
	Sticky            bool   `json:"sticky"`
	Stickyness        uint   `json:"stickyness"`
	Hash              string `json:"hash"`
}
type res_threads []thread_metadata
type reply_metadata struct {
	Post_ID     uint   `json:"post_id"`
	Board       string `json:"board"`
	Is_OP       bool   `json:"is_op"`
	Comment     string `json:"comment"`
	Date_Posted uint   `json:"date_posted"`
	Parent      uint   `json:"parent"`
	Hash        string `json:"hash"`
}
type res_replies []reply_metadata

// Get all available boards. /all/ is not filtered.
//
// Route: /api/v2/boards
func (client *dangeru_client_api) Boards() (res_boards, error) {
	res := res_boards{}
	path := client.addr.PathBoards
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

// Get details for a board.
//
// Route: /api/v2/board/$board$/detail
func (client *dangeru_client_api) BoardDetails(board string) (res_board_details, error) {
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
func (client *dangeru_client_api) Threads(board string, page uint) (res_threads, error) {
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
func (client *dangeru_client_api) ThreadMetadata(id uint) (thread_metadata, error) {
	res := thread_metadata{}
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
func (client *dangeru_client_api) ThreadReplies(id uint) (res_replies, error) {
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
