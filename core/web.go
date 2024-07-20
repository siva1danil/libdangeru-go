package core

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type ClientWeb struct {
	Address string
	Http    http.Client
}

type MainPage struct {
	News             string
	Threads          uint
	Replies          uint
	Archived_Threads uint
	Archived_Replies uint
	Burgs            uint
	Angry_Burgs      uint
}

type ArchivePageEntry struct {
	Board string
	ID    uint
}

func (client *ClientWeb) init() {
	if client.Address == "" {
		client.Address = ADDRESS
	}
}

func (client *ClientWeb) get(path string) ([]byte, error) {
	client.init()

	url := client.Address + "/" + path

	req, err := client.Http.Get(url)
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

// Try to extract latest news and statistics from the main page.
func (client *ClientWeb) MainPage() (MainPage, error) {
	result := MainPage{}
	path := PATH_WEB_MAIN
	data, err := client.get(path)
	if err != nil {
		return result, err
	}

	news := regexp.MustCompile(`<span class="redtext">(.*?)</span>`).FindSubmatch(data)
	if news == nil {
		return result, fmt.Errorf("find: expected match, got nil")
	} else {
		result.News = string(news[1])
	}

	html := strings.NewReplacer("\r", "", "\n", "", " ", "").Replace(string(data))
	threads := regexp.MustCompile(`(\d+)</span>threads`).FindStringSubmatch(html)
	replies := regexp.MustCompile(`(\d+)</span>replies`).FindStringSubmatch(html)
	archived_threads := regexp.MustCompile(`(\d+)</span><ahref="/archive">`).FindStringSubmatch(html)
	archived_replies := regexp.MustCompile(`(\d+)</span>archivedreplies`).FindStringSubmatch(html)
	burgs := regexp.MustCompile(`(\d+)</span>burgs`).FindStringSubmatch(html)
	angry_burgs := regexp.MustCompile(`(\d+)</span>angryburgs`).FindStringSubmatch(html)
	if threads == nil || replies == nil || archived_threads == nil || archived_replies == nil || burgs == nil || angry_burgs == nil {
		return result, fmt.Errorf("find: expected match, got nil")
	} else {
		tmp, _ := strconv.ParseUint(threads[1], 10, 32)
		result.Threads = uint(tmp)
		tmp, _ = strconv.ParseUint(replies[1], 10, 32)
		result.Replies = uint(tmp)
		tmp, _ = strconv.ParseUint(archived_threads[1], 10, 32)
		result.Archived_Threads = uint(tmp)
		tmp, _ = strconv.ParseUint(archived_replies[1], 10, 32)
		result.Archived_Replies = uint(tmp)
		tmp, _ = strconv.ParseUint(burgs[1], 10, 32)
		result.Burgs = uint(tmp)
		tmp, _ = strconv.ParseUint(angry_burgs[1], 10, 32)
		result.Angry_Burgs = uint(tmp)
	}

	return result, nil
}

// Try to extract thread IDs from the archive page.
// Use ClientAPI.ThreadMetadata() to get full metadata.
func (client *ClientWeb) ArchiveIndex(page uint) ([]ArchivePageEntry, error) {
	result := []ArchivePageEntry{}
	path := PATH_WEB_ARCHIVE
	data, err := client.get(path)
	if err != nil {
		return result, err
	}

	entries := regexp.MustCompile(`href="/(.+?)/thread/(\d+?)"`).FindAllSubmatch(data, -1)
	for i := 0; i < len(entries); i++ {
		tmp, _ := strconv.ParseUint(string(entries[i][2]), 10, 32)
		result = append(result, ArchivePageEntry{Board: string(entries[i][1]), ID: uint(tmp)})
	}

	return result, nil
}
