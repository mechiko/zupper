setlocal
  SET PATH=E:\bin\mingw\mingw64\bin;%PATH%
  
  @rem windres --output-format=coff --target=pe-i386 -o resource.syso_64 app.rc
  windres --target=pe-x86-64 -o rsrc_windows_amd64.syso app.rc
  
endlocal
