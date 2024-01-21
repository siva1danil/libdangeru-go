package libdangeru

import (
	"net/http"
)

type dangeru_addr struct {
	Scheme             string
	Domain             string
	PathBoards         string
	PathBoardDetails   string
	PathThreads        string
	PathThreadMetadata string
	PathThreadReplies  string
	PathWebMain        string
	PathWebArchive     string
}

type dangeru_options struct {
	Addr      dangeru_addr
	Client    http.Client
	UserAgent string
	Debug     bool
}

// Initialize a default options object.
func NewOptions() dangeru_options {
	options := dangeru_options{
		Addr: dangeru_addr{
			Scheme:             "https",
			Domain:             "dangeru.us",
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
