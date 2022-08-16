package QR_Code_Generator

import (
	"image/png"
	"log"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Page struct {
	Title string
}

func CreateQRCode(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR Code Generator Service"}

	t, _ := template.ParseFiles("./QR_Code_Generator/qr_code_generator.html")
	err := t.Execute(w, p)
	if err != nil {
		return
	}
}

func ViewQRCode(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")
	log.Println(dataString)

	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	err := png.Encode(w, qrCode)
	if err != nil {
		return
	}
}
