# installing on windows 10 pro x64

this is subject to change and my first approach.

1. install golang
    1. install to `c:\app\Go` (or where ever you like)
    1. `set GOROOT = c:\app\Go`
    1. create dir `c:\Users\[username]\go`
    1. `set GOPATH = c:\Users\[username]\go`
1. install mingw (cygwin does not work)
1. install oracle instant client
    1. install to `C:\app\oracle\instantclient_12_2` (select your version of ic)
        1. add base package
        1. add sdk package
1. `set PKG_CONFIG_PATH=c:\Users\oli\go\src\github.com\codestoke\oracle_exporter\`

## IMPORTANT
Go does not support Cygwin :(
https://github.com/golang/go/wiki/InstallFromSource#install-c-tools