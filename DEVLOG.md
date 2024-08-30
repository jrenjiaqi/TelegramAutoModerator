# Developer's Log

## 1. Work Done So Far
- `[doc]` Added [`README.md`](README.md), [`API_REFERENCE.md`](API_REFERENCE.md), and [`DEVLOG.md`](DEVLOG.md).
- `[doc]` Added [`.gitignore`](.gitignore) to exclude `.env`.
- `[ft.]` Added [`/code`](code/) with Golang.
- `[ft.]` Added [`/code/utils/http.go`](/code/utils/http.go) and [`/code/main.go`](code/main.go) with simple `net/http` example.
- `[doc]` Added [`Food For Thought`](DEVLOG.md#3-food-for-thought) section under [`DEVLOG`](DEVLOG.md).
- `[ft.]` Added [`/code/utils/env.go`](/code/utils/env.go) to load env files.

## 2. Tech Debt
- nothing here; let's hope it stays that way.

## 3. Food For Thought
### 3.1 Golang (compiled) vs Python/NodeJS (intepreted) for JSON decoding:
- Golang (compiled language):
    - Golang is a statically typed, compiled language. It requires structs for JSON decoding because, after compilation, the executable doesn't have the capability to dynamically determine types unless this functionality is explicitly built into the code.
- Python/NodeJS (interpreted languages):
    - Python and NodeJS are dynamically typed, interpreted languages. They don't require predefined types for JSON decoding because the interpreter handles type inference at runtime, allowing them to work with JSON data more flexibly.