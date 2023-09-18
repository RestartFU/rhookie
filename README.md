## Getting Started

**rhookie** may be imported using `go get`:
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

embed := rhookie.Embed{}.WithType("rich").
  WithTitle("My Cool Webhook").
  WithDescription("Sent using rhookie").
  WithColor(15548997)

payload := rhookie.Payload{}.
  WithEmbeds(embed).
  WithUsername("sent with rhookie")

_ = h.SendMessage(payload)
```
Editing an existing message previously sent by the webhook
```go
h := rhookie.NewHook("id", "token")

embed := rhookie.Embed{}.WithType("rich").
  WithTitle("My New Content").
  WithDescription("Edited using rhookie").
  WithColor(5763719)

payload := rhookie.Payload{}.
  WithEmbeds(embed).
  WithUsername("edited with rhookie")

_ = h.EditMessage("id", payload)
```
