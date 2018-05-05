package main

import (
	"os"
	"log"
	"bufio"
	"image/png"
	"image/color"
	"image"
	"sort"

	_ "math"
	"math"
)

func main () {
	f, err := os.Open("modified_7_b.png")
	defer f.Close()
	handleError(err)
	r := bufio.NewReader(f)
	img, err := png.Decode(r)
	handleError(err)

	firefoxred := color.RGBA{175, 78, 55, 255}
	brown := color.RGBA{89, 80, 71, 255}
	startlightgrey := color.RGBA{210, 210, 210, 255}
	stripegrey := color.RGBA{181, 184, 89, 255}
	black := color.RGBA{0, 0, 0, 255}
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	white := color.RGBA{255, 255, 255, 255}
	green := color.RGBA{0, 255, 0, 255}
	nodegreen := color.RGBA{118, 154, 108, 255}
	browncenter := color.RGBA{60, 59, 65, 255}
	brownbigsticker := color.RGBA{68, 62, 62, 255}
	_ = firefoxred
	_ = brown
	_ = startlightgrey
	_ = stripegrey
	_ = black
	_ = red
	_ = blue
	_ = white
	_ = green
	_ = nodegreen
	_ = browncenter
	color0 := startlightgrey
	color1 := stripegrey
	color2 := black
	color3 := white
	color4 := brownbigsticker
	//color5 := browncenter
	// color6 :=

	myPalette := color.Palette{color0, color1, color2, color3, color4}
	newImage := image.NewRGBA(img.Bounds())
	var rows = make([]row, 0, img.Bounds().Max.Y)
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		if (y < 43 ) {
			rows = append(rows, row{y: y, minX: int(math.MinInt64)})
			continue
		}
		if (y > 202) {
			rows = append(rows, row{y: y, minX: int(math.MaxInt64)})
			continue
		}

		// nWhites := 0
		//totalX := 0
		// indexes := []int{0, 0, 0, 0, 0, 0}
		isSet := false
		/*for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if x < 413 || x > 570 {
				continue
			}

			// indexes[myPalette.Index(img.At(x, y))] += 1

			if myPalette.Index(img.At(x, y)) == 2 || myPalette.Index(img.At(x, y)) == 3 {
				fmt.Println(x, y)
				rows = append(rows, row{y: y, minX: -x})
				isSet = true
				//nWhites +=1
				break
			}
			//if myPalette.Index(img.At(x, y)) == 5 {
			//	totalX += x
			//	nWhites += 1
			//	break
			//}

		}*/
		for x := img.Bounds().Max.X - 1; x >= img.Bounds().Min.X; x-- {
			if x < 557 || x > 780 {
				continue
			}
			// indexes[myPalette.Index(img.At(x, y))] += 1

			if myPalette.Index(img.At(x, y)) == 2 || myPalette.Index(img.At(x, y)) == 3{
				rows = append(rows, row{y: y, minX: x})
				isSet = true
				break
			}
			//if myPalette.Index(img.At(x, y)) == 0{
			//	nWhites += 1
			//}

		}
		if !isSet {
			rows = append(rows, row{y: y, minX: int(math.MinInt64)})
		}
		/*if nWhites > 0 {
			rows = append(rows, row{y: y, minX: -totalX / nWhites})
		}*/
		//rows = append(rows, row{y: y, minX: -nWhites})
	}
	sorter := rowsSorter(rows)
	sort.Stable(sorter)

	for i, row := range (rows) {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			originalColor := img.At(x, row.y)
			newImage.Set(x, img.Bounds().Min.Y + i, originalColor)
		}
	}

	// log.Println(rows)

	newFileName := "modified.png"
	newfile, err := os.Create(newFileName)
	handleError(err)
	defer newfile.Close()
	png.Encode(newfile, newImage)
}

type row struct{
	y, minX int
}



func handleError (err error){
	if err != nil {
		log.Fatal(err)
	}
}

func maxIndex (idxs []int) int {
	max := idxs[0]
	i := 0
	for j, v := range(idxs) {
		if v > max {
			max = v
			i = j
		}
	}
	return i

}