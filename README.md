# libdangeru-go

## Overview
`libdangeru-go` is an advanced client library for interacting with Danger/u/, implemented in Go. This library is designed to provide developers with easy-to-use interfaces to access and manipulate data on Danger/u/.

## Features
Currently, `libdangeru-go` offers two main components:

### API Client
- PostThread(): Post a new thread
- PostReply(): Post a reply to a thread
- Boards(): Retrieve a list of all boards.
- BoardDetails(): Get detailed information about a specific board.
- Threads(): Access threads from a specific board.
- ThreadsAll(): Access threads from all boards.
- ThreadMetadata(): Fetch metadata for a given thread.
- ThreadReplies(): Retrieve replies to a specific thread.

### Web Client
- MainPage(): Access the latest news entry and statistics from main page.
- ArchiveIndex(): Fetch the archive index of threads.