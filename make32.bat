setlocal
  @rem x32 https://github.com/brechtsanders/winlibs_mingw/releases/download without LLVM/Clang/LLD/LLDB`
  SET PATH=C:\Go\32\go.120\bin;C:\go\gcc\mingw32\bin;%PATH%

  SET GOARCH=386

  @rem x64 scoop install mingw
  @rem SET PATH=E:\bin\apps\mingw\current\bin;%PATH%

  SET GOOS=windows
  set CGO_ENABLED=1

  copy /y guiapp\resource.syso_32 guiapp\resource.syso

  go build -ldflags "-H=windowsgui -s -w -X 'github.com/mechiko/go4zreport/pkg/entity.Mode=production'" -o distbin/4z_32.exe ./guiapp
  
  copy /y guiapp\resource.syso_64 guiapp\resource.syso

  @rem go build -o 4z.exe ./guiapp

  upx --force-overwrite -o ./distbin/4z_32_upx.exe ./distbin/4z_32.exe

endlocal
