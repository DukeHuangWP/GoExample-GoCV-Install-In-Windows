set enable_shared=ON
set workdir=%cd%
::path當中msys64順序不可以改
set PATH=%SYSTEMROOT%\System32\WindowsPowerShell\v1.0;
@echo 解壓縮使用powershell，若遇到錯誤可自行手動解壓縮msys64.zip到目錄內
@echo 建置過程需要一段時間.....
@echo 注意腳本執行目錄不可用空白符號或特殊字元，可能導致camke編譯時失敗(原因不明)

if exist "%workdir%\build" del /S /Q "%workdir%\build\*"
if not exist "%workdir%\build" mkdir "%workdir%\build"

if not exist "%workdir%\msys64" @echo 正在解壓縮msys64.zip .... && powershell -command "$ProgressPreference = 'SilentlyContinue'; Expand-Archive -Path '%workdir%\msys64.zip' -DestinationPath '%workdir%'"

set PATH=%workdir%\msys64\mingw64\bin;%workdir%\msys64\usr\bin;
@echo %PATH%

cd /D "%workdir%\build"
cmake "%workdir%\opencv-4.5.4" -G "MinGW Makefiles" -B"%workdir%\build" -DENABLE_CXX11=ON -DOPENCV_EXTRA_MODULES_PATH="%workdir%\opencv_contrib-4.5.4\modules" -DBUILD_SHARED_LIBS=%enable_shared% -DWITH_IPP=OFF -DWITH_MSMF=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=OFF -DBUILD_opencv_java=OFF -DBUILD_opencv_python=OFF -DBUILD_opencv_python2=OFF -DBUILD_opencv_python3=OFF -DBUILD_DOCS=OFF -DENABLE_PRECOMPILED_HEADERS=OFF -DBUILD_opencv_saliency=OFF -DBUILD_opencv_wechat_qrcode=OFF -DCPU_DISPATCH= -DOPENCV_GENERATE_PKGCONFIG=ON -DWITH_OPENCL_D3D11_NV=OFF -DOPENCV_ALLOCATOR_STATS_COUNTER_TYPE=int64_t -Wno-dev
mingw32-make -j%NUMBER_OF_PROCESSORS%
mingw32-make install

@echo 建置完成...
@echo 請檢查 %workdir%\msys64 是否解壓縮成功
@echo 請檢查 %workdir%\build\install\x64\mingw\bin 是否編譯完成
cd "%workdir%"
pause
