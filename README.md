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
    -BODY="Sapmple sms text"
    
    Note: Env variable can also assign on the go by commenting init function in main.go
    
### 3. Execute program

```b
$ go run .
```
