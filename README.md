# deleteme mebe

* Scratchpad for trying out buf and other things
* Working hello-world gRPC server and client
* Buf dependency achievement unlocked

# Buf commands used

```shell
buf mod init buf.build/wk/deleteme
```

```shell
buf registry login buf.build --username 'wk'
```

Example: 
```shell
wk@MAC deleteme % buf registry login buf.build --username 'wk'
Token: 
Credentials saved to /Users/wk/.netrc.
```

Created `protos` file and then pushed it to the Buf registry

```shell
buf lint
buf build
buf push --draft will-draft-buf-branch
```
