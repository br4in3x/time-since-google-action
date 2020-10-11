# Time Since - Golang Google Assistant Action

This simple action is my attempt to play with Google Assistant actions in Golang. It can tell you how many days have passed since a certain date, or how many days between two dates.

This action should be deployed as an HTTP server, for example, on Heroku. It acts as a simple REST API responding on Google's requests with valid JSON.

## Poject Config

In the `project_config` folder you will find Google Actions project, which should be pushed to the Google Actions console with [gactions](https://developers.google.com/assistant/actionssdk/gactions) command.

## Development

If you would like to run this action on your machine, use next command:

```bash
PORT=3000 go run ./*.go
```

Or, if you would like to have a cool live-reload feature, use [air](https://github.com/cosmtrek/air):

```bash
go get -u github.com/cosmtrek/air
PORT=3000 air
```
