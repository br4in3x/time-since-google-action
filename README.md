# Time Since - Golang Google Assistant Action

This simple action is my attempt to play with Google Assistant actions in Golang. It can tell you how many days have passed since a certain date, or how many days between two dates.

This action should be deployed as an HTTP server, for example, on Heroku. It acts as a simple REST API responding on Google's requests with valid JSON.

## Poject Config

In the `project_config` folder you will find Google Actions project, which should be pushed to the Google Actions console with [gactions](https://developers.google.com/assistant/actionssdk/gactions) command.
