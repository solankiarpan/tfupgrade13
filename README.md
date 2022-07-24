# tfupgrade13
```
Upgrade terraform syntax to 0.13.0

Usage:
  tfupgrade13 [command]

Available Commands:
  d           It will read the directory
  f           It will read the file
  help        Help about any command
  o           It will read the directory only, not directory inside dorectory.

```

# Installation steps on macos
```
1) Build the project using
go build
2)Move the generated binary tfupgrade13 to a folder
example /user/bin or ${HOME}/bin
3)Add binary to the path
export PATH="${HOME}/bin:${PATH}"
```
