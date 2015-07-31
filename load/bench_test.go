// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package load

import (
	"image/jpeg"
	"image/png"
	"os"
	"testing"

	xtiff "golang.org/x/image/tiff"

	"github.com/chai2010/gdal"
	"github.com/chai2010/tiff"
	"github.com/chai2010/webp"
)

func TestLoad_gdal(t *testing.T) {
	if _, err := gdal.LoadImage("../testdata/video-001.tiff"); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkLoad_video_001_tiff_gdal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := gdal.LoadImage("../testdata/video-001.tiff"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLoad_video_001_tiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := tiff.Load("../testdata/video-001.tiff"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLoad_video_001_tiff_x(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open("../testdata/video-001.tiff")
		if err != nil {
			b.Fatal(err)
		}
		_, err = xtiff.Decode(f)
		if err != nil {
			b.Fatal(err)
		}
		f.Close()
	}
}

func BenchmarkLoad_video_001_webp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := webp.Load("../testdata/video-001.webp"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLoad_video_001_png(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open("../testdata/video-001.png")
		if err != nil {
			b.Fatal(err)
		}
		_, err = png.Decode(f)
		if err != nil {
			b.Fatal(err)
		}
		f.Close()
	}
}

func BenchmarkLoad_video_001_jpeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open("../testdata/video-001.jpeg")
		if err != nil {
			b.Fatal(err)
		}
		_, err = jpeg.Decode(f)
		if err != nil {
			b.Fatal(err)
		}
		f.Close()
	}
}
