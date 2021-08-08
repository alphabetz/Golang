package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("/home/jean/go/src/github.com/alphabetz/Golang/go_qrcode/lost_tag.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
			code_to_gen := colCell
			log.Println("Code to generate: ", code_to_gen)
			qr_code, err := qr.Encode(code_to_gen, qr.L, qr.Auto)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("QR CODE: ", qr_code.Content())

			if code_to_gen != qr_code.Content() {
				log.Fatal("data differs")
			}

			qr_code, err = barcode.Scale(qr_code, 250, 250)
			if err != nil {
				log.Fatal(err)
			}

			writePng(code_to_gen, qr_code)
		}
		fmt.Println()
	}
}

func writePng(filename string, img image.Image) {
	file, err := os.Create(filename + ".png")
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
