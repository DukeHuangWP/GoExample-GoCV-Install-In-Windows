wmic os get Caption,CSDVersion /value
set workdir=%cd%
set PATH=%SYSTEMROOT%\System32\WindowsPowerShell\v1.0;
@echo �����Y�ϥ�powershell�A�Y�J����~�i�ۦ��ʸ����YGoRoot.zip��ؿ���
if not exist "%workdir%\GoRoot" @echo ���b�����YGoRoot.zip .... && powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path '%workdir%\GoRoot.zip' -DestinationPath '%workdir%'"
^
set gopath=%workdir%\GoPath
set CGO_CXXFLAGS=--std=c++11
set CGO_CPPFLAGS="-I%workdir%/build/install/include"
set CGO_LDFLAGS="-L%workdir%/build/install/x64/mingw/lib" -lopencv_core460
::CGO���|�Ѽƥ�����ʳ]�m opencv������
set PATH=%workdir%\msys64\mingw64\bin;%workdir%\build\install\x64\mingw\bin;%workdir%\GoRoot\bin;
::PATH�����]�tmingw64���|
for /f "delims=" %%i in ('go env GOCACHE') do set originalGoCache=%%i
go env -w GOCACHE="%workdir%\GoCache"
::�R��GoCache���s�קK��LGoCV�����v�T
if exist "%workdir%\GoCache" del /S /Q "%workdir%\GoCache\*"
::����GoCV�d��
cd /D "%workdir%\GoCVExample"
go run main.go
^
^
::�٭�GoCache���s�ؿ�
go env -w GOCACHE="%originalGoCache%"
cd..
pause
