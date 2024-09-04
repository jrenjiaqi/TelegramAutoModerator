# API REFERENCE

From the [Telegram Bot documentation](https://core.telegram.org/bots/api)

## 1. Base URL `<base url>`
`https://api.telegram.org/bot<token>/METHOD_NAME`, where `<token>` can be found after creating bot or asking `@BotFather`.

## 2. URL Paths
### 2.1. Get Message with Reaction `[HTTP GET]` `<base url>/getUpdates?allowed_updates=message_reaction` 
From the [documentation](https://core.telegram.org/bots/api#making-requests):
- `/getUpdates` is one of two ways bots can interact with Telegram (the other being `/setWebhook`).
- `/getUpdates` retrieves all updates to anything in the bot's groups from the past 24 hours. 
- Bot must configured in its Telegram group to admin, and be able to access messages.
- **TO RETRIEVE REACTIONS**: retrieved iif request specifies `allowed_updates` with e.g. `message_reaction` in the request body ([see here](https://core.telegram.org/bots/api#update)).
#### 2.1.1 Expected Response
```
[
    {
    "update_id": ...,
    "message": {
        "message_id": 5,
        "from": {
        "id": ...,
        "is_bot": false,
        "first_name": "XXX",
        "username": "XXX"
        },
        "chat": {
        "id": -...,
        "title": "ABC",
        "type": "supergroup"
        },
        "date": 1724913943,
        "text": "hihi"
    }
    },
    {
    "update_id": ...,
    "message_reaction": {
        "chat": {
        "id": -...,
        "title": "ABC",
        "type": "supergroup"
        },
        "message_id": 5,
        "user": {
        "id": ...,
        "is_bot": false,
        "first_name": "XXX",
        "username": "XXX"
        },
        "date": 1724913955,
        "old_reaction": [],
        "new_reaction": [
        {
            "type": "emoji",
            "emoji": "⚡"
        }
        ]
    }
    },
    {
    "update_id": ...,
    "message": {
        "message_id": 6,
        "from": {
        "id": ...,
        "is_bot": false,
        "first_name": "XXX",
        "username": "XXX"
        },
        "chat": {
        "id": -...,
        "title": "ABC",
        "type": "supergroup"
        },
        "date": 1724913981,
        "text": "gonna"
    }
    }
]
```
### 2.2. Delete Message `[HTTP GET]` `<base url>/deleteMessage?chat_id=<chat_id>&message_id=<message_id>` 
Where `chat_id` and `message_id` are fields found in Get Message with Reaction's Expected Response ([2.1.1](#211-expected-response)) 
#### 2.2.1 Message deleted:
```
{
  "ok": true,
  "result": true
}
```

#### 2.2.2. Something went wrong:
```
{
  "ok": false,
  "result": false
}
```

#### 2.2.3 Message not found / already deleted:
```
{
  "ok": false,
  "error_code": 400,
  "description": "Bad Request: message to delete not found"
}
```