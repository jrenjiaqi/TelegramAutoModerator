# **MiniChatSentryBot** : <small>a Golang Telegram Bot</small>

## 1. About the Project ([Github Repository](https://github.com/jrenjq/MiniChatSentryBot))
- This is a Telegram Bot, programmable to perform admin tasks.
- Written in Golang, aims to be stateless (e.g. no need database etc) and deploy-anywhere (with internet).
- Originally developed to manage a university Telegram group.

## 2. What Can It Do?
- `v0.1`: the bot finds messages with **3 or more** counts of "👎" reactions, and **deletes that message** `[downvote]`

## 3. How Does the Code Work?
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

## 4. Can I Just Invite This Bot into My Telegram Group?
### 4.1. NO 
You should **CREATE YOUR OWN BOT** via Telegram's @BotFather, get your own **BOT API TOKEN**, and host the code online to manage this bot (if you really want to).

You could invite this bot as an existing instance on Telegram into your own Telegram group. And you will have to make it an admin, where it'll be able to do many, many things indeed.

But you will **HAVE NO CONTROL OVER THE BOT**, because you don't have the **BOT API TOKEN**. Sometimes in the corner of your eye, you see the bot ... twitching ... talking to itself ... mumbling as it does its job.

The bot that is an admin in your group that you can't control. You can't know for sure what it'll do. And people who make Telegram bots rarely pass the opportunity to play tricks on silly folks. 

Don't be a silly folk :) Fork the code and host your own bot.

## 5. Other Documents
- [`API_REFERENCE`](API_REFERENCE.md): APIs used and explanations.
- [`DEVLOG`](DEVLOG.md): work done and tech debt.

## 5. Setup Project for Development
