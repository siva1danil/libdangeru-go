package libdangeru

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type res_statistics struct {
	Threads          uint
	Replies          uint
	Archived_Threads uint
	Archived_Replies uint
	Burgs            uint
	Angry_Burgs      uint
}
type archive_entry struct {
	Board string
	ID    uint
}
type res_archive_index []archive_entry

// Try to extract the latest news entry from the main page.
func (client *dangeru_client_web) News() (string, error) {
	res := ""
	path := client.addr.PathWebMain
	data, err := client.get(path)
	if err != nil {
		return res, err
	}

	if client.debug {
		fmt.Println(string(data))
	}

	news := regexp.MustCompile(`<span class="redtext">(.*?)</span>`).FindSubmatch(data)
	if news == nil {
		return res, fmt.Errorf("find: expected match, got nil")
	} else {
		res = string(news[1])
	}

	return res, nil
}

// Try to extract the statistics from the main page.
func (client *dangeru_client_web) Statistics() (res_statistics, error) {
	res := res_statistics{}
	path := client.addr.PathWebMain
	data, err := client.get(path)
	if err != nil {
		return res, err
	}

	if client.debug {
		fmt.Println(string(data))
	}

	html := strings.NewReplacer("\r", "", "\n", "", " ", "").Replace(string(data))

	if client.debug {
		fmt.Println(html)
	}

	threads := regexp.MustCompile(`(\d+)</span>threads`).FindStringSubmatch(html)
	replies := regexp.MustCompile(`(\d+)</span>replies`).FindStringSubmatch(html)
	archived_threads := regexp.MustCompile(`(\d+)</span><ahref="/archive">`).FindStringSubmatch(html)
	archived_replies := regexp.MustCompile(`(\d+)</span>archivedreplies`).FindStringSubmatch(html)
	burgs := regexp.MustCompile(`(\d+)</span>burgs`).FindStringSubmatch(html)
	angry_burgs := regexp.MustCompile(`(\d+)</span>angryburgs`).FindStringSubmatch(html)

	if threads == nil || replies == nil || archived_threads == nil || archived_replies == nil || burgs == nil || angry_burgs == nil {
		return res, fmt.Errorf("find: expected match, got nil")
	} else {
		tmp, _ := strconv.ParseUint(threads[1], 10, 32)
		res.Threads = uint(tmp)
		tmp, _ = strconv.ParseUint(replies[1], 10, 32)
		res.Replies = uint(tmp)
		tmp, _ = strconv.ParseUint(archived_threads[1], 10, 32)
		res.Archived_Threads = uint(tmp)
		tmp, _ = strconv.ParseUint(archived_replies[1], 10, 32)
		res.Archived_Replies = uint(tmp)
		tmp, _ = strconv.ParseUint(burgs[1], 10, 32)
		res.Burgs = uint(tmp)
		tmp, _ = strconv.ParseUint(angry_burgs[1], 10, 32)
		res.Angry_Burgs = uint(tmp)
	}

	return res, nil
}

// Try to extract thread IDs from the archive page.
// Use dangeru_client_api.ThreadMetadata() to get full metadata.
func (client *dangeru_client_web) ArchiveIndex(page uint) (res_archive_index, error) {
	res := res_archive_index{}
	path := client.addr.PathWebArchive
	data, err := client.get(path)
	if err != nil {
		return res, err
	}

	if client.debug {
		fmt.Println(string(data))
	}

	entries := regexp.MustCompile(`href="/(.+?)/thread/(\d+?)"`).FindAllSubmatch(data, -1)
	for i := 0; i < len(entries); i++ {
		tmp, _ := strconv.ParseUint(string(entries[i][2]), 10, 32)
		res = append(res, archive_entry{Board: string(entries[i][1]), ID: uint(tmp)})
	}

	return res, nil
}
