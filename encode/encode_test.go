// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encode_test

import (
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"testing"

	xdraw "golang.org/x/image/draw"

	"github.com/chai2010/rawp"
	"github.com/chai2010/webp"
)

func TestFoo(t *testing.T) {
	//
}

func BenchmarkGray_webp_q75(b *testing.B) {
	m := tToGray(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := webp.Encode(ioutil.Discard, m, &webp.Options{
			Lossless: false,
			Quality:  75,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGray_webp_q90(b *testing.B) {
	m := tToGray(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := webp.Encode(ioutil.Discard, m, &webp.Options{
			Lossless: false,
			Quality:  90,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGray_webp_lossless(b *testing.B) {
	m := tToGray(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := webp.Encode(ioutil.Discard, m, &webp.Options{
			Lossless: true,
			Quality:  0,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRGBA_webp_q75(b *testing.B) {
	m := tToRGBA(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := webp.Encode(ioutil.Discard, m, &webp.Options{
			Lossless: false,
			Quality:  75,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRGBA_webp_q90(b *testing.B) {
	m := tToRGBA(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := webp.Encode(ioutil.Discard, m, &webp.Options{
			Lossless: false,
			Quality:  90,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRGBA_webp_lossless(b *testing.B) {
	m := tToRGBA(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := webp.Encode(ioutil.Discard, m, &webp.Options{
			Lossless: true,
			Quality:  0,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGray_rawp(b *testing.B) {
	m := tToGray(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := rawp.Encode(ioutil.Discard, m, &rawp.Options{
			UseSnappy: false,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGray_rawp_snappy(b *testing.B) {
	m := tToGray(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := rawp.Encode(ioutil.Discard, m, &rawp.Options{
			UseSnappy: true,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRGBA_rawp(b *testing.B) {
	m := tToRGBA(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := rawp.Encode(ioutil.Discard, m, &rawp.Options{
			UseSnappy: false,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRGBA_rawp_snappy(b *testing.B) {
	m := tToRGBA(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := rawp.Encode(ioutil.Discard, m, &rawp.Options{
			UseSnappy: true,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGray_png(b *testing.B) {
	m := tToGray(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := png.Encode(ioutil.Discard, m)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRGBA_png(b *testing.B) {
	m := tToRGBA(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := png.Encode(ioutil.Discard, m)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGray_jpeg_q75(b *testing.B) {
	m := tToGray(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := jpeg.Encode(ioutil.Discard, m, &jpeg.Options{Quality: 75})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGray_jpeg_q90(b *testing.B) {
	m := tToGray(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := jpeg.Encode(ioutil.Discard, m, &jpeg.Options{Quality: 90})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRGBA_jpeg_q75(b *testing.B) {
	m := tToRGBA(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := jpeg.Encode(ioutil.Discard, m, &jpeg.Options{Quality: 75})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRGBA_jpeg_q90(b *testing.B) {
	m := tToRGBA(tLoadImage("../testdata/lena512color.png"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := jpeg.Encode(ioutil.Discard, m, &jpeg.Options{Quality: 90})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func tLoadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("tLoadImage: os.Open(%q), err= %v", filename, err)
	}
	defer f.Close()

	m, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("tLoadImage: image.Decode, err= %v", err)
	}
	return m
}

func tToGray(m image.Image) *image.Gray {
	if p, ok := m.(*image.Gray); ok {
		return p
	}
	p := image.NewGray(m.Bounds())
	xdraw.Draw(p, p.Bounds(), m, image.Pt(0, 0), xdraw.Src)
	return p
}

func tToGray16(m image.Image) *image.Gray16 {
	if p, ok := m.(*image.Gray16); ok {
		return p
	}
	p := image.NewGray16(m.Bounds())
	xdraw.Draw(p, p.Bounds(), m, image.Pt(0, 0), xdraw.Src)
	return p
}

func tToRGBA(m image.Image) *image.RGBA {
	if p, ok := m.(*image.RGBA); ok {
		return p
	}
	p := image.NewRGBA(m.Bounds())
	xdraw.Draw(p, p.Bounds(), m, image.Pt(0, 0), xdraw.Src)
	return p
}

func tToRGBA64(m image.Image) *image.RGBA64 {
	if p, ok := m.(*image.RGBA64); ok {
		return p
	}
	p := image.NewRGBA64(m.Bounds())
	xdraw.Draw(p, p.Bounds(), m, image.Pt(0, 0), xdraw.Src)
	return p
}
