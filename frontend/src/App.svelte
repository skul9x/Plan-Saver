<script>
  import { onMount } from 'svelte';
  import { SelectFolder, ExtractAndSave, GetSavedPath } from '../wailsjs/go/main/App.js';
  import './assets/css/main.css';

  let content = '';
  let targetPath = '';
  let isProcessing = false;
  let toast = { show: false, message: '', type: 'success' };

  onMount(async () => {
    try {
      const savedPath = await GetSavedPath();
      if (savedPath) {
        targetPath = savedPath;
      }
    } catch (err) {
      console.error('Failed to load saved path:', err);
    }
  });

  function showToast(message, type = 'success') {
    toast = { show: true, message, type };
    setTimeout(() => {
      toast.show = false;
    }, 3000);
  }

  async function handleSelectFolder() {
    try {
      const selection = await SelectFolder();
      if (selection) {
        targetPath = selection;
      }
    } catch (err) {
      showToast(err, 'error');
    }
  }

  async function handleExtract() {
    if (!content.trim()) {
      showToast('Vui lòng nhập nội dung cần trích xuất', 'error');
      return;
    }
    if (!targetPath) {
      showToast('Vui lòng chọn thư mục lưu', 'error');
      return;
    }

    isProcessing = true;
    try {
      const result = await ExtractAndSave(content, targetPath);
      showToast(result, 'success');
      content = ''; // Clear textarea on success
    } catch (err) {
      showToast(err, 'error');
    } finally {
      isProcessing = true; // Keep it true for a moment for animation
      setTimeout(() => {
        isProcessing = false;
      }, 500);
    }
  }

  function handleClear() {
    content = '';
  }
</script>

<main class="container">
  <header>
    <div class="title-group">
      <h1>Markdown Extractor</h1>
      <p>Trích xuất code blocks từ clipboard/văn bản</p>
    </div>

    <div class="folder-selector">
      <span class="folder-path" title={targetPath || 'Chưa chọn thư mục'}>
        {targetPath || 'Chưa chọn thư mục lưu...'}
      </span>
      <button class="btn-secondary" on:click={handleSelectFolder}>
        {targetPath ? 'Change' : 'Select Folder'}
      </button>
    </div>
  </header>

  <section class="editor-container">
    <textarea
      bind:value={content}
      placeholder="Dán nội dung chứa các block markdown vào đây... (ví dụ: ```filename.md ... ```)"
      disabled={isProcessing}
    ></textarea>
  </section>

  <footer class="actions">
    <button class="btn-text" on:click={handleClear} disabled={isProcessing || !content.trim()}>
      Clear
    </button>
    <button
      class="btn-primary"
      on:click={handleExtract}
      disabled={isProcessing || !content.trim() || !targetPath}
    >
      {#if isProcessing}
        <div class="spinner"></div>
        <span>Đang xử lý...</span>
      {:else}
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="7 10 12 15 17 10"/><line x1="12" y1="15" x2="12" y2="3"/></svg>
        Paste & Extract
      {/if}
    </button>
  </footer>

  {#if toast.show}
    <div class="toast toast-{toast.type}">
      {toast.message}
    </div>
  {/if}
</main>
