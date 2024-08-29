# **MiniChatSentryBot** : <small>a Golang Telegram Bot</small>

## 1. About the Project
- This is a Telegram Bot, programmable to perform admin tasks.
- Written in Golang, aims to be stateless (e.g. no need database etc) and deploy-anywhere (with internet).
- Originally developed to manage a university Telegram group.

## 2. What Can It Do?
- `v0.1`: the bot finds messages with **3 or more** counts of "👎" reactions, and **deletes that message** `[downvote]`

## 3. How Does This Work?
First and foremost, take a gander at the [Telegram Bot documentation](https://core.telegram.org/bots).
1. on Telegram: 
    - create a bot with `@BotFather`
    - invite bot into group as admin who can access messages
    - bot has its own URL on the Telegram server (see [`API_REFERENCE`](API_REFERENCE.md))
2. in your code: 
    - bot calls its URL with a path for what information it wants (eg. `/getUpdates`)
    - bot gets that information from Telegram as a HTTP response
3. in your code: 
    - bot does something in response to the HTTP response (e.g. delete a message)
    - bot waits a while before calling its URL again *(goto step 2)*.

## 4. Other Documents
- [`API_REFERENCE`](API_REFERENCE.md): APIs used and explanations.
- [`DEVLOG`](DEVLOG.md): work done and tech debt.

## 5. Setup Project for Development
