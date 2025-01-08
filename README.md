# 🚀 GUI QR Code Generator

## 📝 Description
A desktop application developed in **Golang** using the **Fyne** framework that allows generating QR codes from user-inputted text. Additionally, generated QR codes can be saved in **PNG** or **JPG** formats.

## 🎯 Features
- **User-friendly graphical interface** built with Fyne.
- Dynamic QR code generation.
- Ability to save QR codes in **PNG** and **JPG** formats.

## 🛠️ Technologies Used
- **Golang**: Core language of the project.
- **Fyne**: GUI framework for Golang.
- **image/jpeg**: Image encoding for JPG format.

## 📦 Installation
1. Clone the repository:
   ```bash
   git clone git@github.com:stescobedo92/gui-qrcode-generator.git
   cd qr-code-generator
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

## 💻 Usage
1. Enter text into the input field.
2. Click the **"Generate QR"** button to create the QR code.
3. Save the QR code from the **File > Save** menu in **PNG** or **JPG** format.
4. Access the **Help > About** section for more details about the application.

## 📂 Project Structure
```
qr-code-generator/
├── resources/
│   └── qr-code-app-icon.png
├── main.go
├── qrutils/
│   ├── qrutils.go
├── go.mod
├── go.sum
└── README.md
```

## 📜 License
This project is licensed under the **GNU General Public License v3.0 (GPLv3)**. You can find the full license text [here](https://www.gnu.org/licenses/gpl-3.0.en.html).

## 🤝 Contributing
Contributions are welcome! Open an **Issue** or submit a **Pull Request**.
