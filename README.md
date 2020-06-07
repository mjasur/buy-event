# buy-event
- clone the repository
- run the command `go build`
- go install (make sure GOBIN environment variable is set)

Commands:
- `buy-event add client` or `buy-event add purchase` in order to add new Client or Purchase entry into corresponding file.
- `buy-event client list` or `buy-event purchase ls` to print entries to standard output
- `buy-event notify <purchase_id> sms/mail` This command sends notification to client about the purchase. Last argument is optional, by default sends to email.
- `buy-event -h` or `buy-event -help`. Additional command which is not implemented yet.
