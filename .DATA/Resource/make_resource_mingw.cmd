setlocal
  SET PATH=E:\bin\mingw\mingw64\bin;%PATH%
  
  @rem windres --output-format=coff --target=pe-i386 -o resource.syso alcogo.rc
  windres --target=pe-i386 -F pe-i386 -o rsrc_windows_386.syso app.rc
  
endlocal
