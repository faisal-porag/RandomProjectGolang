package QR_Code_Generator

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"log"
	"os"
)

func QRCOdeGeneratorIMG() {
	// Create the barcode
	qrCode, _ := qr.Encode("FAISAL PORAG", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("qrcode_dd.png")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	// encode the barcode as png
	err := png.Encode(file, qrCode)
	if err != nil {
		log.Println("something went wrong:",err)
		return
	}
}
