# Setup Guide

## 0. Table of Contents
1. [Setup for Development](#1-setup-for-development)
2. [Setup for Deployment](#2-setup-for-deployment)

## 1. Setup for Development
1. install Go

2. then run these commands from the project root:
```bash
cd code
touch .env
touch .conf
touch deletedMessages.log
mkdir claude/prompts
touch claude/prompts/systemPrompt.txt
```

3. paste into the `.env` file:
```
API_URL=https://api.telegram.org/bot
BOT_TOKEN=<token from @BotFather>

GET_UPDATES_PATH=getUpdates
ALLOWED_UPDATES_NAME=allowed_updates
MESSAGE_REACTIONS_NAME=message_reaction

DELETE_MESSAGE_PATH=deleteMessage
CHAT_ID_NAME=chat_id
MESSAGE_ID_NAME=message_id
```

4. paste into the `.conf` file:
```
THUMBS_DOWN_COUNT_TO_DELETE_MSG=<int>
DEBUG_MODE=<true/false>
```

5. paste the system prompt into the `claude/prompts/systemPrompt.txt` file, **as one single line**.

6. then run:
`go run .`

## 2. Setup for Deployment
1. install Go

2. then run these commands from the project root:
```bash
mkdir bin
cd bin
touch .env
touch .conf
touch deletedMessages.log
mkdir claude/prompts
touch claude/prompts/systemPrompt.txt
```

3. paste into the `.env` file:
```
API_URL=https://api.telegram.org/bot
BOT_TOKEN=<token from @BotFather>

GET_UPDATES_PATH=getUpdates
ALLOWED_UPDATES_NAME=allowed_updates
MESSAGE_REACTIONS_NAME=message_reaction

DELETE_MESSAGE_PATH=deleteMessage
CHAT_ID_NAME=chat_id
MESSAGE_ID_NAME=message_id
```

4. paste into the `.conf` file:
```
THUMBS_DOWN_COUNT_TO_DELETE_MSG=<int>
DEBUG_MODE=<true/false>
```

5. paste the system prompt into the `claude/prompts/systemPrompt.txt` file, **as one single line**.

6. then run from `/`
```bash
go build -o ./bin
```

7. and run from `/` to run the bot once.
```bash
./MiniChatSentryBot
```