wmic os get Caption,CSDVersion /value
set workdir=%cd%
set PATH=%SYSTEMROOT%\System32\WindowsPowerShell\v1.0;
@echo �����Y�ϥ�powershell�A�Y�J����~�i�ۦ��ʸ����YGoRoot.zip��ؿ���
if not exist "%workdir%\GoRoot" @echo ���b�����YGoRoot.zip .... && powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path '%workdir%\GoRoot.zip' -DestinationPath '%workdir%'"
^
set gopath=%workdir%\GoPath
set CGO_CXXFLAGS=--std=c++11
set CGO_CPPFLAGS="-I%workdir%/build/install/include"
set CGO_LDFLAGS="-L%workdir%/build/install/x64/mingw/lib" -lopencv_core460 -lopencv_face460 -lopencv_videoio460 -lopencv_imgproc460 -lopencv_highgui460 -lopencv_imgcodecs460 -lopencv_objdetect460 -lopencv_features2d460 -lopencv_video460 -lopencv_dnn460 -lopencv_xfeatures2d460 -lopencv_plot460 -lopencv_tracking460 -lopencv_img_hash460 -lopencv_calib3d460 -lopencv_bgsegm460 -lopencv_photo460 -lopencv_aruco460 -lopencv_ximgproc460
::CGO���|�Ѽƥ�����ʳ]�m opencv������
set PATH=%workdir%\msys64\mingw64\bin;%workdir%\build\install\x64\mingw\bin;%workdir%\GoRoot\bin;
::PATH�����]�tmingw64���|
for /f "delims=" %%i in ('go env GOCACHE') do set originalGoCache=%%i
go env -w GOCACHE="%workdir%\GoCache"
::�R��GoCache���s�קK��LGoCV�����v�T
if exist "%workdir%\GoCache" del /S /Q "%workdir%\GoCache\*"
cd /D "%workdir%\gocv@0.31.0"
go run cmd\version\main.go
^
::�٭�GoCache���s�ؿ�
go env -w GOCACHE="%originalGoCache%"
cd..
pause
