setlocal
  SET PATH=C:\go\gcc\mingw64\bin;%PATH%
  
  windres --target=pe-x86-64 -o resource.syso_64 app.rc  

endlocal
