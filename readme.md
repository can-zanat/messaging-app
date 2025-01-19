##  How It Works
**_docker-compose up --build_ command will be enough to run the project.**
**you have to wait about 10 second for the database to be ready.**

```bash
    docker-compose up --build
```

**or you can run this api _make run_ command.**

---

**This API is written using hexagonal architecture and consists of 3 endpoint designed to fulfill the following
requirements of the case. Below you can see this endpoint and the specific
topics of the data they provide:**

#### GetSentMessages _(it gets sent messages)_
```bash 
  curl --location 'http://127.0.0.1:96/sent-messages'
```
#### StartSending _(it starts sending messages process)_
```bash 
  curl --location --request POST 'http://127.0.0.1:96/start-sending'
```
#### StopSending _(it stops sending messages process)_
```bash 
  curl --location --request POST 'http://127.0.0.1:96/stop-sending'
```