setlocal
  SET PATH=C:\go\gcc\mingw64\bin;%PATH%
  SET GOARCH=amd64

  SET GOOS=windows
  set CGO_ENABLED=1
  copy /y cmd\resource.syso_64 cmd\resource.syso

  go build -ldflags "-H=windowsgui -s -w -X 'zupper/config.Mode=production'" -o ./.dist/zupper.exe ./cmd

  @rem go build -o 4z.exe ./guiapp

  @rem upx --force-overwrite -o ./.dist/4z_upx.exe ./.dist/4z.exe

endlocal
