// The reader picks how a lesson is shown: one card at a time ("slides") or
// the whole lesson as one scrolling page ("article"). The choice is kept per
// browser; the lesson template reads it before the first paint.
(() => {
  const KEY = "paseo:view";

  function load() {
    try {
      return localStorage.getItem(KEY) === "article" ? "article" : "slides";
    } catch {
      return "slides";
    }
  }

  function save(view) {
    try {
      localStorage.setItem(KEY, view);
    } catch {
      // Storage unavailable (private mode, blocked): the choice is not kept.
    }
  }

  const current = load();

  for (const option of document.querySelectorAll('input[name="view"]')) {
    option.checked = option.value === current;
    option.addEventListener("change", () => {
      save(option.value);
      // A lesson is laid out once, on load, so it reloads to show the new view.
      if (document.querySelector(".deck")) location.reload();
    });
  }
})();
