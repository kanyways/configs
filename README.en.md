# configs

----

Config project is used to read and configure the configuration files of the application. Currently, it supports the configuration files in yaml, yml, JSON and toml formats.
Personal projects are not strictly original, but source code from Reference 1.Tks [qianlnk](https://github.com/qianlnk)

```yml
server:
  name: config.yaml
```

```yml
server:
  name: config.yml
```

```json
{
  "server": {
    "name": "config.json"
  }
}
```

```toml
[server]
name="config.toml"
```

See [demo](example/main.go) for a program case


## References

- [https://github.com/qianlnk/config](https://github.com/qianlnk/config)
- [Golang Learning - TOML Configuration Processing](https://www.cnblogs.com/CraryPrimitiveMan/p/7928647.html)