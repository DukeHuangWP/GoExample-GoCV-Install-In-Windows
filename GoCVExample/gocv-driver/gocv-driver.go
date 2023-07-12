package gocvdriver

import (
	"image"
	"image/draw"
	_ "image/jpeg" //image jpg轉檔必須載入
	_ "image/png"  //image png轉檔必須載入
	"os"

	"gocv.io/x/gocv"
)

// GoCV圖片解析後特徵值資料庫，使用slice[index]方式來獲取資料
type ImageGoCVDB struct {
	fileDictionary map[string]struct{} //用來檢查key重複
	PicFilePath    []string            //紀錄圖片路徑
	ImgWidth       []int               //紀錄圖片寬度
	ImgHeight      []int               //紀錄圖片高度
	GoCVMat        []*gocv.Mat         //儲存圖片GoCV特徵值
}

func (db *ImageGoCVDB) AddData(picFilePath string) (err error) {

	if db.fileDictionary == nil {
		db.fileDictionary = make(map[string]struct{})
	} else if _, isExsit := db.fileDictionary[picFilePath]; isExsit == true {
		return nil
	} else {
		db.fileDictionary[picFilePath] = struct{}{}
	}

	img, err := GetImageFromFilePath(picFilePath)
	if err != nil {
		return err
	}

	goCVMat, err := ImageToGoCVMat(img)
	if err != nil {
		return err
	}

	db.GoCVMat = append(db.GoCVMat, &goCVMat)
	db.PicFilePath = append(db.PicFilePath, picFilePath)
	db.ImgWidth = append(db.ImgWidth, img.Bounds().Dx())
	db.ImgHeight = append(db.ImgHeight, img.Bounds().Dy())
	return
}

func GetImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func ImageToGoCVMat(img image.Image) (gocv.Mat, error) {
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()
	bytes := make([]byte, 0, x*y*3)

	for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			bytes = append(bytes, byte(b>>8), byte(g>>8), byte(r>>8))
		}
	}
	return gocv.NewMatFromBytes(y, x, gocv.MatTypeCV8UC3, bytes)
} // image.Image轉換成gocv.Mat

func ImagesMatchTemplate(fullPictureImg image.Image, sampleGoCVMats []*gocv.Mat, interruptConfidence float32) (sampleIndex int, maxConfidence float32, X, Y int, err error) {

	fullPicGoCVMat, err := ImageToGoCVMat(fullPictureImg)
	defer fullPicGoCVMat.Close()
	if err != nil {
		//log.Printf("錯誤 > %v)", err)
		return
	}

	for index := 0; index < len(sampleGoCVMats); index++ {
		sampleGoCVMat := sampleGoCVMats[index]

		// 开始匹配
		resultMat := gocv.NewMat()
		defer resultMat.Close()

		maskGoCVMat := gocv.NewMat() //作用不明
		gocv.MatchTemplate(fullPicGoCVMat, *sampleGoCVMat, &resultMat, gocv.TmCcoeffNormed, maskGoCVMat)
		maskGoCVMat.Close()
		// 获取最大匹配度 和 匹配范围
		_, tempMaxConfidence, _, tempMaxLoc := gocv.MinMaxLoc(resultMat)
		//log.Println(fmt.Sprintf("max confidence %f, %v, %v", cacheMaxConfidence, cacheMaxLoc.X, cacheMaxLoc.Y))
		if tempMaxConfidence >= interruptConfidence { //發現極高辨識率故直接中斷後續圖片辨識
			return index, tempMaxConfidence, tempMaxLoc.X, tempMaxLoc.Y, nil
		} else if tempMaxConfidence > maxConfidence {
			sampleIndex = index
			maxConfidence = tempMaxConfidence
			X = tempMaxLoc.X
			Y = tempMaxLoc.Y
		}
	}

	return
}

func ImagesMatchTemplateMuilt(fullPictureImg image.Image, sampleGoCVMat *gocv.Mat, minConfidence float32, maxMatch int) (matchCount int, confidences []float32, Xs, Ys []int, err error) {

	fullPicGoCVMat, err := ImageToGoCVMat(fullPictureImg)
	defer fullPicGoCVMat.Close()
	if err != nil {
		//log.Printf("錯誤 > %v)", err)
		return
	}

	fullPicImg := image.NewRGBA(image.Rect(0, 0, fullPictureImg.Bounds().Dx(), fullPictureImg.Bounds().Dy())) //必須先創建一個RGBA格式，再將圖片資料複製進去
	draw.Draw(fullPicImg, fullPictureImg.Bounds(), fullPictureImg, image.ZP, draw.Src)
	fullPicImgBounds := fullPicImg.Bounds()
	fullPicImgWidth := fullPicImgBounds.Dx()
	fullPicImgHeight := fullPicImgBounds.Dy()

	for index := 0; index < maxMatch; index++ {

		resultMat := gocv.NewMat()
		defer resultMat.Close()

		maskGoCVMat := gocv.NewMat() //作用不明
		gocv.MatchTemplate(fullPicGoCVMat, *sampleGoCVMat, &resultMat, gocv.TmCcoeffNormed, maskGoCVMat)
		maskGoCVMat.Close()
		// 获取最大匹配度 和 匹配范围
		_, tempMaxConfidence, _, tempMaxLoc := gocv.MinMaxLoc(resultMat)
		//log.Println(fmt.Sprintf("max confidence %f, %v, %v", cacheMaxConfidence, cacheMaxLoc.X, cacheMaxLoc.Y))
		if tempMaxConfidence < minConfidence {
			return
		} else {
			matchCount++
			confidences = append(confidences, tempMaxConfidence)
			locX := tempMaxLoc.X
			locY := tempMaxLoc.Y
			Xs = append(Xs, locX)
			Ys = append(Ys, locY)

			overWrtieImg := image.NewRGBA(image.Rect(0, 0, fullPicImgWidth, fullPicImgHeight))
			draw.Draw(overWrtieImg, overWrtieImg.Bounds(), image.Transparent, image.Pt(0, 0), draw.Src) //創建一個具備樣本大小透明圖片，利用複蓋圖片達到移除效果
			draw.Draw(fullPicImg, image.Rect(locX, locY, fullPicImgWidth+locX, fullPicImgHeight+locY), overWrtieImg, image.ZP, draw.Src)

			fullPicGoCVMat, err = ImageToGoCVMat(fullPicImg) //移除後重新辨識
			if err != nil {
				//log.Printf("錯誤 > %v)", err)
				break //若移除錯誤則中斷辨識
			}
		}
	}

	return
}
