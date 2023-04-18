wmic os get Caption,CSDVersion /value
set workdir=%cd%
set PATH=%SYSTEMROOT%\System32\WindowsPowerShell\v1.0;
@echo 解壓縮使用powershell，若遇到錯誤可自行手動解壓縮GoRoot.zip到目錄內
if not exist "%workdir%\GoRoot" @echo 正在解壓縮GoRoot.zip .... && powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path '%workdir%\GoRoot.zip' -DestinationPath '%workdir%'"
^
set gopath=%workdir%\GoPath
set CGO_CXXFLAGS=--std=c++11
set CGO_CPPFLAGS="-I%workdir%/build/install/include"
set CGO_LDFLAGS="-L%workdir%/build/install/x64/mingw/lib" -lopencv_core460 -lopencv_face460 -lopencv_videoio460 -lopencv_imgproc460 -lopencv_highgui460 -lopencv_imgcodecs460 -lopencv_objdetect460 -lopencv_features2d460 -lopencv_video460 -lopencv_dnn460 -lopencv_xfeatures2d460 -lopencv_plot460 -lopencv_tracking460 -lopencv_img_hash460 -lopencv_calib3d460 -lopencv_bgsegm460 -lopencv_photo460 -lopencv_aruco460 -lopencv_ximgproc460
::CGO路徑參數必須手動設置 opencv版本號
set PATH=%workdir%\msys64\mingw64\bin;%workdir%\build\install\x64\mingw\bin;%workdir%\GoRoot\bin;
::PATH必須包含mingw64路徑
for /f "delims=" %%i in ('go env GOCACHE') do set originalGoCache=%%i
go env -w GOCACHE="%workdir%\GoCache"
::刪除GoCache佔存避免其他GoCV版本影響
if exist "%workdir%\GoCache" del /S /Q "%workdir%\GoCache\*"
cd /D "%workdir%\gocv@0.31.0"
go run cmd\version\main.go
^
::還原GoCache佔存目錄
go env -w GOCACHE="%originalGoCache%"
cd..
pause
