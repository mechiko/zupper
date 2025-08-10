setlocal
  SET PATH=E:\bin\mingw\mingw64\bin;%PATH%
  
  windres --target=pe-i386 -F pe-i386 -o rsrc_windows_386.syso app.rc
  windres --target=pe-x86-64 -o rsrc_windows_amd64.syso app.rc
  
endlocal
