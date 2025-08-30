# set shell := ["pwsh", "", "-CommandWithArgs"]
# set positional-arguments
shebang := 'pwsh.exe'
# Variables
exe_name := "xmlparser"
mod_name := "xmlparser"
dist := ".dist"
ld_flags :="-s -w"

default:
  just --list

win64:
    #!{{shebang}}
    $env:Path = "C:\Go\go.124\bin;C:\go\gcc\mingw64\bin;" + $env:Path
    $env:GOARCH = "amd64"
    $env:GOOS = "windows"
    $env:CGO_ENABLED = 1
    if (-Not (Test-Path go.mod)) {
      go mod init {{mod_name}}
    }
    go mod tidy -go 1.24 -v
    if(-Not $?) { exit }
    Remove-Item {{dist}}\{{exe_name}}.exe, {{dist}}\{{exe_name}}_64.exe 2>$null
    if (-Not (Test-Path {{dist}})) { New-Item -ItemType Directory -Force {{dist}} | Out-Null }
    go build -ldflags="{{ld_flags}}" -o {{dist}}\{{exe_name}}_64.exe ./cmd
    if(-Not $?) { exit }
    if (Get-Command upx -ErrorAction SilentlyContinue) {
      upx --force-overwrite -o {{dist}}\{{exe_name}}.exe {{dist}}\{{exe_name}}_64.exe
    } else {
      Copy-Item {{dist}}\{{exe_name}}_64.exe {{dist}}\{{exe_name}}.exe
    }
