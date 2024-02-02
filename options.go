package libdangeru

import (
	"net/http"
)

type Addr struct {
	Scheme             string
	Domain             string
	PathPostThread     string
	PathPostReply      string
	PathBoards         string
	PathBoardDetails   string
	PathThreads        string
	PathThreadMetadata string
	PathThreadReplies  string
	PathWebMain        string
	PathWebArchive     string
}

type ClientOptions struct {
	Addr      Addr
	Client    http.Client
	UserAgent string
	Debug     bool
}

// Initialize a default options object.
func NewOptions() ClientOptions {
	options := ClientOptions{
		Addr: Addr{
			Scheme:             "https",
			Domain:             "dangeru.us",
			PathPostThread:     "post",
			PathPostReply:      "reply",
			PathBoards:         "api/v2/boards",
			PathBoardDetails:   "api/v2/board/%s/detail",
			PathThreads:        "api/v2/board/%s?page=%d",
			PathThreadMetadata: "api/v2/thread/%d/metadata",
			PathThreadReplies:  "api/v2/thread/%d/replies",
			PathWebMain:        "",
			PathWebArchive:     "archive",
		},
		Client:    http.Client{},
		UserAgent: "libdangeru-go",
		Debug:     false,
	}
	return options
}
