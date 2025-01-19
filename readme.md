##  How It Works
**_docker-compose up --build_ command will be enough to run the project.**
**you have to wait about 10 second for the database to be ready.**

```bash
    docker-compose up --build
```

**or you can run this api _make run_ command.**

---

**This API is written using hexagonal architecture and consists of 3 endpoint designed to fulfill the following
requirements of the case. This repo included unit tests and integration tests. Below you can see this endpoint and the 
specific topics of the data they provide:**

#### GetSentMessages _(it gets sent messages)_
This endpoint returns the messages that have been sent to the recipient. It filters the messages by is_sent fields.

**REQUEST**
```bash 
  curl --location 'http://127.0.0.1:96/sent-messages'
```
**200 - response**
```json
[
  {
    "id": "678d47289b919db6d6425bd8",
    "content": "Meeting at 3 PM.",
    "recipient": "user2",
    "sent_time": "2025-01-19T21:42:04+03:00",
    "is_sent": true
  }
]
```
**500 - response**
```json
{
  "message": "error message"
}
```

#### StartSending _(it starts sending messages process)_
This endpoint starts the process of sending messages to the recipient. It sends the messages to the recipient 
every 2 minutes. It starts automatically when the project is started. In this endpoint it is checked whether the
process is already running or not. If it is already running, it returns an error message. Additionally, it is checked
content length of the message. If the content length is greater than 250 characters, it logs the message and continue
to send messages process. If there is no message to send, it logs "no message to send" and continue to send messages 
process. 

**REQUEST**
```bash 
  curl --location --request POST 'http://127.0.0.1:96/start-sending'
```
**200 - response**
```json
{
  "message": "Sending process started."
}
```
**500 - response**
```json
{
  "message": "sending is already running"
}
```

#### StopSending _(it stops sending messages process)_
This endpoint stops the process of sending messages.

**REQUEST**
```bash 
  curl --location --request POST 'http://127.0.0.1:96/stop-sending'
```
**200 - response**
```json
{
  "message": "start process ended"
}
```
**500 - response**
```json
{
  "message": "sending is already not running"
}
```