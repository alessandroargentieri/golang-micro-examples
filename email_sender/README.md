
# Golang email sender

Little program to send emails via Go.

## Usage example

```bash
go run main.go \
  --smtp-host=smtp.myprovider.com \
  --smtp-port=465 \
  --sender=myemailhere \
  --sender-name="Your name" \
  --password='mypasswordhere' \
  --recipient=recipient@email.com \
  --subject="Hello there" \
  --plain-body="This is a test email sent from my Go program."
```