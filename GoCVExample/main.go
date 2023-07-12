package main

import (
	gocvdriver "GoCVExample/gocv-driver"
	"log"
)

func main() {

	goCVDB := &gocvdriver.ImageGoCVDB{}
	err := goCVDB.AddData("./sample.png") //將sample圖片分析成特徵值後存入DB
	if err != nil {
		log.Printf("錯誤 > %v)", err)
	}

	fullPictureName := "./fullPicture.png"
	fullPictureImg, err := gocvdriver.GetImageFromFilePath(fullPictureName)
	if err != nil {
		log.Printf("錯誤 > %v)", err)
	}

	sampleIndex, maxConfidence, X, Y, err := gocvdriver.ImagesMatchTemplate(fullPictureImg, goCVDB.GoCVMat, 0.85) //通常85%的辨識率即可
	sampleName := goCVDB.PicFilePath[sampleIndex]

	log.Printf("在圖片'%v'當中尋找相似圖'%v', 執行結果 > 最大相似度%v 位置(%v,%v) 錯誤:%v\n", fullPictureName, sampleName, maxConfidence, X, Y, err)
	//2023/07/12 12:57:28 在圖片'./fullPicture.png'當中尋找相似圖'./sample.png', 執行結果 > 最大相似度0.95035076 位置(1024,140) 錯誤:<nil>
	return
}
