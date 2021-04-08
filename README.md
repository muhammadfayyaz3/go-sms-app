# GO SMS/EMAIL APP

## Developer Setup

### 1. Update Config
{config/config.development.json}

#### Twilio Setup

    -ACCOUNT_SID
    -AUTH_TOKEN
    -FROM_NUMBER

#### SendGrid Setup    

    -SENDGRID_API_KEY

### 2. Update SMS, TO & BODY environment variable 
{main.go}

#### For SMS

    -TO="+123456789"
    -BODY="Sample sms text"

#### For Email

    -EMAIL_FROM_NAME
    -EMAIL_FROM
    -EMAIL_TO_NAME
    -EMAIL_TO
    -EMAIL_SUBJECT
    -EMAIL_BODY  

    Note: Env variable can also assign on the go by commenting init function in main.go

```b
$ go run .
```
