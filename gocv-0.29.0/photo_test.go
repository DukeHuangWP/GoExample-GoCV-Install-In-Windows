package gocv

import (
	"image"
	"testing"
)

func TestColorChange(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	mask := src.Clone()
	defer mask.Close()

	ColorChange(src, mask, &dst, 1.5, .5, .5)
	if dst.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invlalid ColorChange test")
	}
}

func TestSeamlessClone(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst := NewMatWithSize(30, 30, MatTypeCV8UC3)
	defer dst.Close()
	blend := NewMatWithSize(dst.Rows(), dst.Cols(), dst.Type())
	defer blend.Close()
	mask := src.Clone()
	defer mask.Close()

	center := image.Point{dst.Cols() / 2, dst.Rows() / 2}
	SeamlessClone(src, dst, mask, center, &blend, NormalClone)
	if blend.Empty() || dst.Rows() != blend.Rows() || dst.Cols() != blend.Cols() {
		t.Error("Invlalid SeamlessClone test")
	}
}

func TestIlluminationChange(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	mask := src.Clone()
	defer mask.Close()

	IlluminationChange(src, mask, &dst, 0.2, 0.4)
	if dst.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invlalid IlluminationChange test")
	}
}

func TestTextureFlattening(t *testing.T) {
	src := NewMatWithSize(20, 20, MatTypeCV8UC3)
	defer src.Close()
	dst := NewMat()
	defer dst.Close()
	mask := src.Clone()
	defer mask.Close()

	TextureFlattening(src, mask, &dst, 30, 45, 3)
	if dst.Empty() || dst.Rows() != src.Rows() || dst.Cols() != src.Cols() {
		t.Error("Invlalid TextureFlattening test")
	}
}

func TestFastNlMeansDenoisingColoredMultiWithParams(t *testing.T) {
	var src [3]Mat
	for i := 0; i < 3; i++ {
		src[i] = NewMatWithSize(20, 20, MatTypeCV8UC3)
		defer src[i].Close()
	}

	dst := NewMat()
	defer dst.Close()

	FastNlMeansDenoisingColoredMultiWithParams([]Mat{src[0], src[1], src[2]}, &dst, 1, 1, 3, 3, 7, 21)

	if dst.Empty() || dst.Rows() != src[0].Rows() || dst.Cols() != src[0].Cols() {
		t.Error("Invalid FastNlMeansDenoisingColoredMultiWithParams test")
	}
}

func TestMergeMertens(t *testing.T) {
	var src [3]Mat
	for i := 0; i < 3; i++ {
		src[i] = NewMatWithSize(20, 20, MatTypeCV8UC3)
		defer src[i].Close()
	}

	dst := NewMat()
	defer dst.Close()

	mertens := NewMergeMertens()
	defer mertens.Close()

	mertens.Process([]Mat{src[0], src[1], src[2]}, &dst)

	if dst.Empty() || dst.Rows() != src[0].Rows() || dst.Cols() != src[0].Cols() {
		t.Error("Invalid TestMergeMertens test")
	}
}

func TestNewAlignMTB(t *testing.T) {
	var src [3]Mat
	for i := 0; i < 3; i++ {
		src[i] = NewMatWithSize(20, 20, MatTypeCV8UC3)
		defer src[i].Close()
	}

	alignwtb := NewAlignMTB()
	defer alignwtb.Close()

	var dst []Mat
	alignwtb.Process([]Mat{src[0], src[1], src[2]}, &dst)

	sizedst := len(dst)
	t.Logf(" Size Dst slice : %d ", sizedst)
	if sizedst > 0 {
		if dst[0].Empty() || dst[0].Rows() != src[0].Rows() || dst[0].Cols() != src[0].Cols() {
			t.Error("Invalid TestNewAlignMTB test")
		}
	}
	if sizedst <= 0 {
		t.Error("Invalid TestNewAlignMTB test : empty result")
	}
}
