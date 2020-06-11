# configs

----

config 项目是用来读取及配置应用程序的配置文件的，目前支持yaml、yml、json、toml格式的配置文件。
个人的项目严格来说不是原创，是来至参考资料①的源码。感谢[qianlnk](https://github.com/qianlnk)

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

程序案例请查看[demo](example/main.go)

## 参考资料

- [https://github.com/qianlnk/config](https://github.com/qianlnk/config)
- [Golang学习--TOML配置处理](https://www.cnblogs.com/CraryPrimitiveMan/p/7928647.html)