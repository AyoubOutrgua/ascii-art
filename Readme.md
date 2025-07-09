# ASCII‑Art

**ASCII‑Art** is a Go CLI tool that transforms any printable characters (letters, numbers, spaces, special symbols, and `\n`) into large, 8‑line‑tall ASCII art banner. It supports style: `standard`.

---

## 🚀 Features

- Converts input strings into banner-style ASCII art  
- Handles multi-line input via `\n`  
- Processes all printable ASCII characters (space to `~`)  
- Uses modular banner file (`standard.txt`)

---

## ⚙️ Installation & Usage

To run the ASCII-Art program:

```bash
git clone https://learn.zone01oujda.ma/git/boulhaj/ascii-art.git
cd ascii-art
go run . "your text here"
