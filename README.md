# FolderFilesWatcherNotification

A small CLI tool that watches folders or files for modifications like added, changed or deleted files. It then sends a notification to a specific endpoint, e.g. a Slack channel, an URI (HTTP request) or execute a command.

It makes use of the [fsnotify Go library](https://github.com/fsnotify/fsnotify) to provide cross-platform filesystem notifications.
