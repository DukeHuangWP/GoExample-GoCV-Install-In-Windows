wmic os get Caption,CSDVersion /value
set workdir=%cd%
set PATH=%SYSTEMROOT%\System32\WindowsPowerShell\v1.0;
@echo 解壓縮使用powershell，若遇到錯誤可自行手動解壓縮GoRoot.zip到目錄內
if not exist "%workdir%\GoRoot" @echo 正在解壓縮GoRoot.zip .... && powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path '%workdir%\GoRoot.zip' -DestinationPath '%workdir%'"
^
set gopath=%workdir%\GoPath
set CGO_CXXFLAGS=--std=c++11
set CGO_CPPFLAGS=-I%workdir%/build/install/include
set CGO_LDFLAGS=-L%workdir%/build/install/x64/mingw/lib -lopencv_core454 -lopencv_face454 -lopencv_videoio454 -lopencv_imgproc454 -lopencv_highgui454 -lopencv_imgcodecs454 -lopencv_objdetect454 -lopencv_features2d454 -lopencv_video454 -lopencv_dnn454 -lopencv_xfeatures2d454 -lopencv_plot454 -lopencv_tracking454 -lopencv_img_hash454 -lopencv_calib3d454 -lopencv_bgsegm454 -lopencv_photo454 -lopencv_aruco454
::CGO路徑參數必須手動設置
set PATH=%workdir%\msys64\mingw64\bin;%workdir%\build\install\x64\mingw\bin;%workdir%\GoRoot\bin;
::PATH必須包含mingw64路徑
go env -w GOCACHE=%workdir%\GoCache
::刪除GoCache佔存避免其他GoCV版本影響
if exist "%workdir%\GoCache" del /S /Q "%workdir%\GoCache\*"
cd /D "%workdir%\gocv-0.29.0"
go run cmd\version\main.go
^
cd..
pause
