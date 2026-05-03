# Phase 04: Integration & Persistence
Status: ✅ Completed
Dependencies: Phase 02, Phase 03

## Objective
Kết nối Frontend với Backend và lưu cấu hình người dùng.

## Requirements
### Functional
- [x] Gọi hàm Go từ Svelte.
- [x] Lưu đường dẫn thư mục đã chọn vào file config (ví dụ: `config.json` hoặc sử dụng Wails store).
- [x] Tự động load thư mục đã chọn khi mở app.

### Non-Functional
- [x] Trải nghiệm người dùng liền mạch.

## Implementation Steps
1. [x] Kết nối nút "Paste & Extract" với hàm `ExtractAndSave` trong Go.
2. [x] Thực hiện logic đọc/ghi config file phía Go.
3. [x] Cập nhật UI dựa trên kết quả trả về từ Go.

## Files to Create/Modify
- `app.go` - Thêm logic config.
- `frontend/src/App.svelte` - Gọi Wails bindings.

## Test Criteria
- [ ] Chọn folder một lần, tắt app mở lại vẫn thấy folder đó.
- [ ] Dán text và bấm nút -> File xuất hiện trong folder đích.

---
Next Phase: [Phase 05: Testing & Polish](file:///home/skul9x/Desktop/Test_code/Plan-Saver/plans/260503-1910-markdown-extractor/phase-05-testing.md)
