## Getting Started

Gophig may be imported using `go get`:
```
go get github.com/restartfu/rhookie
```

## Usage

Creating a new `Hook` instance:
```go
h := rhookie.NewHook("id", "token")
```
Sending a message using the `Hook`
```go
h := rhookie.NewHook("id", "token")

var payload rhookie.Payload
payload.Embeds = []rhookie.Embed {
  {
    Title:       "My Cool Webhook",
    Description: "Sent using rhookie",
    Color:       15548997,
  },
}
_ = h.SendMessage(payload)
```
Editing an existing message previously sent by the webhook
```go
h := rhookie.NewHook("id", "token")

var payload rhookie.Payload
payload.Embeds = []rhookie.Embed {
  {
    Title:       "My New Content",
    Description: "Edited using rhookie",
    Color:       5763719,
  },
}
_ = h.EditMessage("id", payload)
```
