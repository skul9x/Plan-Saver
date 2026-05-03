# 💡 BRIEF: Markdown Extractor (Plan-Saver)

**Ngày tạo:** 2026-05-03
**Brainstorm cùng:** skul9x

---

## 1. VẤN ĐỀ CẦN GIẢI QUYẾT
Người dùng cần trích xuất nhanh các đoạn nội dung Markdown nằm trong các block ` ```md ` từ một đoạn văn bản dài (ví dụ: kết quả từ AI) và lưu chúng thành các file `.md` riêng biệt một cách tự động.

## 2. GIẢI PHÁP ĐỀ XUẤT
Một Desktop App sử dụng **Wails (Go + Svelte)**:
- Giao diện đơn giản với Textarea để dán nội dung.
- Nút "Paste" để dán nhanh từ clipboard.
- Dialog chọn thư mục lưu trữ và ghi nhớ thư mục này.
- Tự động tạo thư mục con theo thời gian và lưu các file `.md` với tên file là dòng đầu tiên của mỗi block.

## 3. ĐỐI TƯỢNG SỬ DỤNG
- Lập trình viên, người viết lách, người dùng AI cần lưu trữ kết quả có cấu trúc.

## 4. TÍNH NĂNG

### 🚀 MVP (Bắt buộc có):
- [ ] Giao diện dán văn bản (Svelte).
- [ ] Dialog chọn thư mục (Go/Wails).
- [ ] Ghi nhớ thư mục đã chọn (Persistence).
- [ ] Logic tách các block ` ```md ... ``` `.
- [ ] Tạo thư mục con theo timestamp.
- [ ] Lưu file .md với tên trích xuất từ dòng đầu của block.

### 🎁 Phase 2 (Làm sau):
- [ ] Xem trước (Preview) danh sách file sẽ được tạo.
- [ ] Chỉnh sửa tên file trước khi lưu.
- [ ] Lịch sử các lần extract.

## 5. ƯỚC TÍNH SƠ BỘ
- **Độ phức tạp:** Đơn giản - Trung bình (do cần tích hợp Wails).
- **Rủi ro:** Quyền ghi file trên OS, xử lý tên file chứa ký tự đặc biệt.

## 6. BƯỚC TIẾP THEO
→ Chạy `/plan` để tạo các file phase chi tiết.
