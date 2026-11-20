// Stores completed lessons in localStorage, per browser. No server involved.
(() => {
  const KEY = "paseo:done";

  function load() {
    try {
      return new Set(JSON.parse(localStorage.getItem(KEY)) || []);
    } catch {
      return new Set();
    }
  }

  function save(done) {
    try {
      localStorage.setItem(KEY, JSON.stringify([...done]));
    } catch {
      // Storage unavailable (private mode, blocked): progress is not kept.
    }
  }

  const done = load();

  // Lesson page: the "mark as done" checkbox.
  for (const box of document.querySelectorAll("[data-done]")) {
    box.checked = done.has(box.dataset.done);
    box.addEventListener("change", () => {
      box.checked ? done.add(box.dataset.done) : done.delete(box.dataset.done);
      save(done);
    });
  }

  // Index page: highlight completed lessons.
  for (const item of document.querySelectorAll("[data-lesson]")) {
    item.classList.toggle("is-done", done.has(item.dataset.lesson));
  }
})();
