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

  // markStops lights up the stops of completed lessons, on the index route
  // and in the lesson sidebar.
  function markStops(done) {
    for (const item of document.querySelectorAll("[data-lesson]")) {
      item.classList.toggle("is-done", done.has(item.dataset.lesson));
    }
  }

  const done = load();
  markStops(done);

  // Lesson page: the "mark as done" checkbox.
  for (const box of document.querySelectorAll("[data-done]")) {
    box.checked = done.has(box.dataset.done);
    box.addEventListener("change", () => {
      box.checked ? done.add(box.dataset.done) : done.delete(box.dataset.done);
      save(done);
      markStops(done);
    });
  }

  // Index page: how far the reader got in each module.
  for (const module of document.querySelectorAll("[data-module]")) {
    const lessons = module.querySelectorAll("[data-lesson]").length;
    const finished = module.querySelectorAll("[data-lesson].is-done").length;
    const progress = module.querySelector("[data-progress]");
    if (finished > 0) progress.textContent = `${finished} de ${lessons} hechas`;
    module.classList.toggle("is-complete", finished === lessons);
  }

  // Index page: once the walk has started, the main button picks it up at
  // the first lesson still to do instead of at the very first one.
  const resume = document.querySelector("[data-continue]");
  const next = document.querySelector("[data-lesson]:not(.is-done) a");
  if (resume && next && done.size > 0) {
    resume.href = next.href;
    resume.textContent = `Sigue con «${next.textContent.trim()}»`;
  }
})();
