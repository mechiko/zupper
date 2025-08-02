setlocal
  SET PATH=E:\bin\mingw32\bin;%PATH%
  
  windres --target=pe-i386 -o resource.syso_32 alcogo.rc

endlocal
