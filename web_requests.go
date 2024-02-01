package libdangeru

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

// Try to extract latest news and statistics from the main page.
func (client *ClientWeb) Main() (MainPage, error) {
	result := MainPage{}
	path := client.addr.PathWebMain
	data, err := client.get(path)
	if err != nil {
		return result, err
	}

	if client.debug {
		fmt.Println(string(data))
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
// Use dangeru_client_api.ThreadMetadata() to get full metadata.
func (client *ClientWeb) ArchiveIndex(page uint) ([]ArchivePageEntry, error) {
	result := []ArchivePageEntry{}
	path := client.addr.PathWebArchive
	data, err := client.get(path)
	if err != nil {
		return result, err
	}

	if client.debug {
		fmt.Println(string(data))
	}

	entries := regexp.MustCompile(`href="/(.+?)/thread/(\d+?)"`).FindAllSubmatch(data, -1)
	for i := 0; i < len(entries); i++ {
		tmp, _ := strconv.ParseUint(string(entries[i][2]), 10, 32)
		result = append(result, ArchivePageEntry{Board: string(entries[i][1]), ID: uint(tmp)})
	}

	return result, nil
}
