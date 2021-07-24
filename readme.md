# **Subber**

Subber is a golang CLI tool that interacts with a PubSub local emulator.

## **Installation**
`go get github.com/arthureichelberger/subber`

## **Commands**
### **Create a PubSub Topic.**

`subber createTopic`

### **Create a PubSub Subscription.**
`subber createSub`

### **List all topics.**
`subber listTopics`

### **List all subscriptions.**
`subber listSubs`

### **Read messages from subscription**
`subber readSub`

### **Publish a message in a topuc**
`subber publish`

You can also pass a payload argument in case you'd like to cat a JSON file directly like this:
`subber publish --payload="$(cat file.json)"`