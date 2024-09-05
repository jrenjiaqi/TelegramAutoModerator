# Developer's Log

## 1. Work Done So Far
- `[doc]` Added [`README.md`](README.md), [`API_REFERENCE.md`](API_REFERENCE.md), and [`DEVLOG.md`](DEVLOG.md).
- `[doc]` Added [`.gitignore`](.gitignore) to exclude `.env`.
- `[ft.]` Added [`/code`](code/) with Golang.
- `[ft.]` Added [`/code/utils/http.go`](/code/utils/http.go) and [`/code/main.go`](code/main.go) with simple `net/http` example.
- `[doc]` Added [`Food For Thought`](DEVLOG.md#3-food-for-thought) section under [`DEVLOG`](DEVLOG.md).
- `[ft.]` Added [`/code/utils/env.go`](/code/utils/env.go) to load env files.
- `[ft.]` Added [`/code/repo/getURI.go`](/code/repo/getURI.go) to represent getting URI path from env.
- `[ft.]` Added [`/code/repo/getUpdateJSON.go`](/code/repo/getUpdateJSON.go) to get JSON of updates seen by Telegram bot (e.g. messages, message reactions).
- `[ft.]` Added [`/code/repo/getMsgToDelete.go`](/code/repo/getMsgToDelete.go) to get messages to delete from thumbs down reaction updates (as per response JSON).
- `[ft.]` Added [`/code/repo/deleteMsg.go`](/code/repo/deleteMsg.go) to delete messages that meets thumbs down count criteria.

## 2. Tech Debt
- nothing here; let's hope it stays that way.

## 3. Food For Thought
### 3.1 Golang (compiled) vs Python/NodeJS (intepreted) for JSON decoding:
- Golang (compiled language):
    - Golang is a statically typed, compiled language. It requires structs for JSON decoding because, after compilation, the executable doesn't have the capability to dynamically determine types unless this functionality is explicitly built into the code.
- Python/NodeJS (interpreted languages):
    - Python and NodeJS are dynamically typed, interpreted languages. They don't require predefined types for JSON decoding because the interpreter handles type inference at runtime, allowing them to work with JSON data more flexibly.

### 3.2 Golang Struct vs Python Dictionary
- Golang (compiled): cannot freely add new key value pair into struct.
- Python (interpreted): can dynamically add new key value pair into dictionary.
- conclusion: should use Golang Maps instead of structs for such cases.

### 3.3 Iteration over slice of structs vs slice of strings-formed-from-structs
- is O(n^2) and complex to do comparison and edits.
- if struct fields are **unique** (e.g. chat id and message id), concatenate them into 1 string ...
- ... and iterate over a slice of string would be much simpler.
- e.g. `[{chat_id: xyz, message_id: 12345678}...]` -> `["xyz_12345678"...]`