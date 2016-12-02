# FolderFilesWatcherNotification

A small CLI tool that watches folders or files for modifications like added, changed or deleted files. It then sends a notification to a specific endpoint, e.g. a Slack channel, an URI (HTTP request) or execute a command.

It makes use of the [github.com/rjeczalik/notify Go library](https://github.com/rjeczalik/notify) to provide cross-platform filesystem notifications.
