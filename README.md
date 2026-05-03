# 🚀 Plan-Saver

**Plan-Saver** là một công cụ desktop mạnh mẽ và tinh gọn, được thiết kế để giúp các nhà phát triển và người dùng AI nhanh chóng trích xuất, tổ chức và lưu trữ các kế hoạch (plans) hoặc mã nguồn từ các khối văn bản Markdown.

![Banner](https://img.shields.io/badge/Status-Production-success?style=for-the-badge)
![Tech](https://img.shields.io/badge/Built%20with-Wails%20|%20Go%20|%20Svelte-blue?style=for-the-badge)

---

## ✨ Tính năng chính

- 🔍 **Trích xuất thông minh**: Tự động nhận diện các khối mã (code blocks) Markdown từ văn bản được dán vào.
- 📁 **Tổ chức tự động**: Lưu trữ các file vào các thư mục được đánh dấu thời gian (timestamp), giúp quản lý phiên bản dễ dàng.
- ⚙️ **Ghi nhớ cấu hình**: Tự động lưu lại thư mục đích cuối cùng bạn đã chọn.
- 🖥️ **Giao diện hiện đại**: Sử dụng Svelte mang lại trải nghiệm mượt mà, phản hồi nhanh.
- 🛡️ **An toàn & Riêng tư**: Toàn bộ quá trình xử lý diễn ra offline trên máy tính của bạn.

---

## 🛠️ Công nghệ sử dụng

Ứng dụng được xây dựng trên nền tảng công nghệ hiện đại nhất:

- **Backend**: [Go](https://go.dev/) - Hiệu năng cao và ổn định.
- **Frontend**: [Svelte](https://svelte.dev/) & [Vite](https://vitejs.dev/) - Giao diện nhanh, nhẹ và tương tác tốt.
- **Framework**: [Wails v2](https://wails.io/) - Kết hợp sức mạnh của Go và sự linh hoạt của Web công nghệ để tạo ứng dụng Desktop.

---

## 🚀 Hướng dẫn cài đặt

### Yêu cầu hệ thống
- **Go** (v1.18+)
- **Node.js** (v16+) & **NPM**
- **Wails CLI** (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)

### Các bước cài đặt cho nhà phát triển

1. **Clone repository**:
   ```bash
   git clone https://github.com/skul9x/Plan-Saver.git
   cd Plan-Saver
   ```

2. **Chạy ở chế độ Development**:
   ```bash
   wails dev
   ```

3. **Build ứng dụng**:
   - **Windows**: `wails build -platform windows/amd64`
   - **Linux**: `wails build -platform linux/amd64`

---

## 📖 Cách sử dụng

1. **Chọn thư mục**: Nhấp vào nút "Thay đổi thư mục" để chọn nơi bạn muốn lưu các bản kế hoạch.
2. **Dán nội dung**: Sao chép văn bản chứa các khối Markdown (ví dụ: từ ChatGPT, Claude hoặc tài liệu kỹ thuật) và dán vào ô văn bản chính.
3. **Trích xuất & Lưu**: Nhấn nút "Paste & Extract". Ứng dụng sẽ tự động tách các khối mã và lưu thành các file riêng biệt trong một thư mục mới.

---

## 📂 Cấu trúc thư mục

```text
Plan-Saver/
├── frontend/         # Mã nguồn giao diện (Svelte)
├── build/            # Các tài nguyên build và icon
├── app.go            # Logic ứng dụng chính (Go)
├── parser.go         # Logic trích xuất Markdown
├── main.go           # Điểm khởi đầu của ứng dụng
└── wails.json        # Cấu hình dự án Wails
```

---

## 👨‍💻 Tác giả

Dự án được phát triển bởi **skul9x**.

---
*Chúc bạn có những trải nghiệm tuyệt vời với Plan-Saver!*
