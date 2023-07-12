## 前言
由於[GoCV對於Windows環境下使用糟糕的腳本撰寫方式]( https://github.com/hybridgroup/gocv/blob/release/win_build_opencv.cmd )，在參考了[moonchant12](https://github.com/moonchant12/install-gocv)建置腳本撰寫方式後，在此示範如何完整建置GoCV與其相依套件OpenCV，此git已包含所有需要工具無需額外再下載其他元件。

## 包含工具
 * msys64
    * [msys2-x86_64-20230318](https://github.com/msys2/msys2-installer/releases/download/2023-03-18/msys2-x86_64-20230318.exe)
        * make-4.4.1-1
        * cmake-3.26.0-1
        * gcc-11.3.0-3
        * mingw-w64-x86_64-cmake-3.26.0-1
        * mingw-w64-x86_64-gcc-12.2.0-10
 * opencv套件
    * [opencv-4.6.0](https://github.com/opencv/opencv/archive/4.6.0.zip)
    * [opencv_contrib-4.6.0](https://github.com/opencv/opencv_contrib/archive/4.6.0.zip)
 * Go & GoCV
    * [Go-1.20.3](https://go.dev/dl/go1.20.3.windows-amd64.zip)
    * [GoCV@0.31.0](https://github.com/hybridgroup/gocv/archive/refs/tags/v0.31.0.zip)

## 測試通過
   * ``Windwos 7 x64`` ✔
   * ``Windwos 10 x64`` ✔
   * ``Windwos 11 x64`` ✔
 
![Alt text](/參考資料/Windows%207%20x64%20Pass.png)


## 建置方式
   0. cmake建置過程可以能會下載缺失元件，網路需要打開
   1. 確認``cmd.exe``可以正常執行
   2. 將本git複製到不含有空白符號或特殊字元的目錄(可能導致camke編譯時失敗)
   3. 執行``01. build_OpenCV_bin.cmd``，進行編譯OpenCV相依套件
   4. 執行``02. check_GoCV_version.cmd``，若正確回傳``gocv version``和``opencv lib version``，則表示OpenCV與GoCV編譯成功
   5. 執行``03. runExample.cmd``，若正確回傳圖片辨識座標，則表示OpenCV與GoCV執行成功
## 檔案樹說明
   1. 可依照自身需求修改版跟本檔案樹的內容
   2. 編譯過程會將 ``GoRoot.zip`` 和 ``msys64.zip`` 分別解壓縮到 ``./GoRoot`` 和 ``./msys64``目錄
   3. 編譯完成的OpenCV相依套件需要將``./build/install/x64/mingw/bin`` 設置到``Path``環境變數後再配合``mingw64``終端機使用
   4. 編譯GoCV前需要將環境變數``CGO_CXXFLAGS``和``CGO_CPPFLAGS``指定OpenCV相依套件目錄(可參考腳本``02. check_GoCV_version.cmd ``和``03. runExample.cmd``)

   ```bash
  ├─ /build #編譯OpenCV後產生檔案的佔存目錄
  ├─ /GoCVExample #gocv尋找相似圖範例
  │   ├─ /gocv@0.31.0 #官方PKG
  │   ├─ /gocv-driver #gocv應用方式範例PKG
  │   ├─ fullPicture.png #完整圖片
  │   ├─ sample.png #印本圖片
  │   └─ main.go #範例執行入口
  │
  ├─ /GoCache #執行check_GoCV_version.cmd產生Golang編譯佔存檔案目錄
  ├─ /GoPath #執行check_GoCV_version.cmd產生Golang PKG佔存檔案目錄
  ├─ /opencv-4.6.0
  ├─ /opencv_contrib-4.6.0
  ├─ /參考資料
  ├─ 01. build_OpenCV_bin.cmd #編譯OpenCV相依套件腳本
  ├─ 02. check_GoCV_version.cmd #確認OpenCV與GoCV版本腳本
  ├─ 03. runExample.cmd #編譯完OpenCV後，執行gocv尋找相似圖範例
  ├─ GoRoot.zip #編譯時所使用的Golang，配合Git LFS政策做成壓縮檔
  └─ msys64.zip #編譯時所使用的終端機參考上方文件，配合Git LFS政策做成壓縮檔
   ```
