# Phase 02: Backend Logic (Go)
Status: ✅ Completed
Dependencies: Phase 01

## Objective
Xây dựng các hàm xử lý phía Go: Parser markdown, Dialog chọn folder, Lưu file.

## Requirements
### Functional
- [x] Hàm trích xuất các block ` ```md ` từ chuỗi string.
- [x] Hàm mở dialog chọn thư mục (Wails Runtime).
- [x] Logic tạo thư mục con theo timestamp.
- [x] Logic lưu file `.md` với tên lấy từ dòng đầu của block.

### Non-Functional
- [x] Xử lý lỗi (invalid path, file permission).
- [x] Xử lý ký tự đặc biệt trong tên file.

## Implementation Steps
1. [x] Viết struct `App` và đăng ký các method cho Wails.
2. [x] Viết hàm `SelectFolder()` sử dụng `runtime.OpenDirectoryDialog`.
3. [x] Viết hàm `ExtractAndSave(content string, path string)`:
    - Regex hoặc string split để tìm các block.
    - Tạo folder con.
    - Ghi file.
4. [x] Unit test cho logic extract.

## Files to Create/Modify
- `app.go` - Chứa logic chính.
- `parser.go` - (Tùy chọn) Tách riêng logic xử lý text.

## Test Criteria
- [x] Hàm extract trả về đúng số lượng block từ input mẫu (`answer.txt`).
- [x] Folder được tạo đúng format timestamp.
- [x] File `.md` được lưu với nội dung chính xác.

---
Next Phase: [Phase 03: Frontend UI](file:///home/skul9x/Desktop/Test_code/Plan-Saver/plans/260503-1910-markdown-extractor/phase-03-frontend.md)
