# **Subber**

Subber is a golang CLI tool that allows you to interact with a PubSub local emulator.

## **Launch the PubSub emulator**
A Google PubSub emulator is ready-to-use in the docker-compose file in the repo. If you have docker installed on your local machine, just clone the repo and start the container.
```bash
docker-compose up -d
```

## **Installation**
The easiest way to get the tool is to download directly with `go get`.
```bash
go get github.com/arthureichelberger/subber
```


You can also get the whole repository and install the tool.
```bash
git clone git@github.com:arthureichelberger/subber.git
cd subber
go install
```

## **Usage**
_For the reference, every command comes with a help flag (`--help or -h`) that describes every flag and every possible usage._

### **Create a PubSub Topic.**
```bash
subber createTopic
```
It will create a topic, if no other exists with the same ID, inside the emulator.

### **Create a PubSub Subscription.**
```bash
subber createSub
```
It will create a subscription, if no other exists with the same ID, inside the emulator.

### **List all topics.**
```bash
subber listTopics
```
It will list all topics within the emulator.

### **List all subscriptions.**
```bash
subber listSubs
```
It will list all subscriptions within the emulator.

### **Read messages from subscription**
```bash
subber readSub
```
It will read messages from a subscription. By default, process will exit after 10 messages read. This can be configured through the `--maxMessages=10` flag.
You can also choose to read the subscription interactively with the `--interactively` flag and you'll have the choice to ack or nack every message.

### **Publish a message in a topuc**
```bash
subber publish
```
It will allow you to publish a message in a topic.

You can also pass a payload argument in case you'd like to cat a JSON file directly like this:
```bash
subber publish --payload="$(cat file.json)"
```
