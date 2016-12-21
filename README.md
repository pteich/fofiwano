# Watcher and Notifications for File and Folder changes

A small CLI tool that watches folders or files for modifications like added, changed or deleted files. It then sends a notification to a specific endpoint, e.g. a Slack channel, an URI (HTTP request) or execute a command. 

It makes use of the [github.com/rjeczalik/notify Go library](https://github.com/rjeczalik/notify) to provide cross-platform filesystem notifications.

Fofiwano is available as pre-built binaries for macOS, Linux, FreeBSD and Windows on the [release page](https://github.com/pteich/fofiwano/releases).

Create a config file `.fofiwano.yml` in the current directory or in your `$HOME` with content like this (see [example](.fofiwano.example.yml)):
```yaml
watching:
  # recursive watching ./test
  - target: ./test/...
    notifications:
      - notify: slack
        event: all
        options:
          channel: "#test"
          username: "fofiwano"
          webhook_url: "https://hooks.slack.com/services/..."
          icon_emoji: ":monkey_face:"
          footer: "fofiwano"
      - notify: http
        event: write
        options:
          uri: http://my.endpoint.com

  # only watching ./test2 and NO sub-folders
  - target: ./test2
    notifications:
      - notify: http
        event: write
        options:
          uri: http://test.com

```

- `target` can be a single file or a folder. Add `/...` to a folder to create a recursive watcher that also reacts to modifications in sub-folders.`
- `notifications` is an array of notification providers with their options. *Right now only Slack notifications are available!* (more to come).
- `event` can be one of `all`, `write`, `create`, `remove` or `rename`

*Hint:* You don't have to stick with YAML if you don't like it. You can write your config in every format that [Viper](https://github.com/spf13/viper) supports (JSON, TOML, YAML, HCL, and Java properties config files).

Start your watcher like so:
```bash
fofiwano watch
```

You can also provide the path to your config file with the `--config` flag:
```bash
fofiwano --config ~/.fofiwano.yml watch
```

Get help about all available flags and commands:
```bash
fofiwano --help
```

## TODO:

- add more notification providers (command execution)
- add templates for fine grained control over notification messages
- add queue for notification events to prevent locks
- add tests

## Similar tools

Tools with similar functionality:
- https://github.com/splitbrain/Watcher
- http://incron.aiken.cz/
