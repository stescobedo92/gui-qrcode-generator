package main

import (
	"bytes"
	_ "embed"
	"github.com/stescobedo92/gui-qrcode-generator/qrutils"
	"image"
	"image/jpeg"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

//go:embed resources/qr-code-app-icon.png
var qrCodeIconData []byte

func main() {
	a := app.New()
	w := a.NewWindow("QR Code Generator")
	w.Resize(fyne.NewSize(400, 400))

	// Centrar la ventana en la pantalla
	w.CenterOnScreen()

	// Cargar y establecer el ícono de la aplicación
	qrCodeIconResource := fyne.NewStaticResource("qr-code-app-icon.png", qrCodeIconData)
	w.SetIcon(qrCodeIconResource)

	// UI Components
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text for QR")
	qrImage := canvas.NewImageFromResource(nil)
	qrImage.SetMinSize(fyne.NewSize(256, 256))

	// Function to generate QR and update the image
	generateQR := func() {
		text := input.Text
		if text == "" {
			dialog.ShowInformation("Error", "Please enter some text.", w)
			return
		}

		data, err := qrutils.GenerateQRCode(text, 256)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		// Update the image in the UI
		qrImage.Resource = fyne.NewStaticResource("qr.png", data)
		qrImage.Refresh()
	}

	generateButton := widget.NewButton("Generate QR", generateQR)

	// Menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Save", func() {
			if input.Text == "" {
				dialog.ShowInformation("Error", "Please generate a QR code before saving.", w)
				return
			}

			saveDialog := dialog.NewFileSave(
				func(write fyne.URIWriteCloser, err error) {
					if err != nil {
						dialog.ShowError(err, w)
						return
					}
					if write == nil {
						// User canceled the operation
						return
					}
					defer write.Close()

					// Get the text and generate the QR
					text := input.Text
					if text == "" {
						dialog.ShowInformation("Error", "Please enter some text.", w)
						return
					}

					data, err := qrutils.GenerateQRCode(text, 256)
					if err != nil {
						dialog.ShowError(err, w)
						return
					}

					// Determine the format based on the file extension
					uri := write.URI()
					ext := uri.Extension()

					var imgData []byte
					switch ext {
					case ".png":
						imgData = data
					case ".jpg", ".jpeg":
						img, _, err := image.Decode(bytes.NewReader(data))
						if err != nil {
							dialog.ShowError(err, w)
							return
						}
						buf := new(bytes.Buffer)
						err = jpeg.Encode(buf, img, nil)
						if err != nil {
							dialog.ShowError(err, w)
							return
						}
						imgData = buf.Bytes()
					default:
						dialog.ShowInformation("Unsupported Format", "Please choose PNG or JPG format.", w)
						return
					}

					// Write the image data directly to the file
					_, err = write.Write(imgData)
					if err != nil {
						dialog.ShowError(err, w)
						return
					}

					dialog.ShowInformation("Saved", "QR Code saved successfully.", w)

					// Clear all fields after saving
					input.SetText("")
					qrImage.Resource = nil
					qrImage.Refresh()
				}, w)

			saveDialog.SetFileName("qr.png")
			saveDialog.SetFilter(
				storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}),
			)
			saveDialog.Show()
		}),
		fyne.NewMenuItem("Quit", func() {
			a.Quit()
		}),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowInformation("About", "QR Code Generator\n\nCreated by:\nSergio Triana Escobedo\nv1.0", w)
		}),
	)

	menu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	w.SetMainMenu(menu)

	// UI Layout
	content := container.NewVBox(
		input,
		generateButton,
		qrImage,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
