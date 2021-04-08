# GO SMS APP

## Twilio Setup

### 1. Update Config
{config/config.development.json}

    -ACCOUNT_SID
    -AUTH_TOKEN
    -FROM_NUMBER
    
### 2. Update SMS, TO & BODY environment variable 
{main.go}

    -TO="+123456789"
    -BODY="Sample sms text"
    
    Note: Env variable can also assign on the go by commenting init function in main.go

## SendGrid Setup

### 1. Update Config
{config/config.development.json}

    -SENDGRID_API_KEY
    
### 2. Update Email environment variable 
{main.go}

    -EMAIL_FROM_NAME
    -EMAIL_FROM
    -EMAIL_TO_NAME
    -EMAIL_TO
    -EMAIL_SUBJECT
    -EMAIL_BODY    

### 3. Execute program

```b
$ go run .
```
