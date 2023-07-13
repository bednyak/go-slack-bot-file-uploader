# Go-CRUD-API-example

### Preparation

1. Create an account in Slack
2. Go to https://api.slack.com/apps?new_app=1 and create slack app ![setup-1.png](./docs/images/setup-1.png) ![setup-2.png](./docs/images/setup-2.png) ![setup-3.png](./docs/images/setup-3.png)
3. Enable sockets ![setup-4.png](./docs/images/setup-4.png) ![setup-5.png](./docs/images/setup-5.png)
4. Enable event subscriptions ![setup-6.png](./docs/images/setup-6.png)
5. Configure Slack OAuth and Permissions with adding scopes ![setup-7.png](./docs/images/setup-7.png) ![setup-8.png](./docs/images/setup-8.png) ![setup-9.png](./docs/images/setup-9.png)
6. Get Slack Bot User OAuth Token and set SLACK_BOT_TOKEN in config file ![setup-11.png](./docs/images/setup-11.png)
7. Get Slack Channel ID and set CHANNEL_ID in config file ![setup-12.png](./docs/images/setup-12.png) ![setup-13.png](./docs/images/setup-13.png)
8. Add file-bot to the channel in Slack ![setup-14.png](./docs/images/setup-14.png) ![setup-15.png](./docs/images/setup-15.png)

### Run
```
go run main.go
```

Go application should send 'upload-file-example.txt' to linked Slack channel ![result.png](./docs/images/result.png)