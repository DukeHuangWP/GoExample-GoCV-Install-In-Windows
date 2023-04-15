wmic os get Caption,CSDVersion /value
set workdir=%cd%
set gopath=%workdir%\GoPath
set PATH=%SYSTEMROOT%\System32\WindowsPowerShell\v1.0;
if not exist "%workdir%\GoRoot" @echo ¥¿¦b¸ÑÀ£ÁYGoRoot.zip .... && powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path '%workdir%\GoRoot.zip' -DestinationPath '%workdir%'"
^
set PATH=%workdir%\msys64\mingw64\bin;%workdir%\msys64\usr\bin;%workdir%\build\install\x64\mingw\bin;%workdir%\GoRoot\bin;
cd /D "%workdir%\gocv-0.29.0"
go run cmd\version\main.go
^
cd..
pause





