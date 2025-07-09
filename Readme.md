# ASCII‑Art

**ASCII‑Art** is a Go CLI tool that transforms any printable characters (letters, numbers, spaces, special symbols, and `\n`) into large, 8‑line‑tall ASCII art banners. It supports multiple styles: `standard`, `shadow`, and `thinkertoy`.

---

## 🚀 Features

- Converts input strings into banner-style ASCII art  
- Handles multi-line input via `\n`  
- Processes all printable ASCII characters (space to `~`)  
- Uses modular banner files (e.g. `standard.txt`, `shadow.txt`, `thinkertoy.txt`)

---

## ⚙️ Installation & Usage

To run the ASCII-Art program:

```bash
git clone <YOUR_REPO_URL>
cd ascii-art
go mod tidy
