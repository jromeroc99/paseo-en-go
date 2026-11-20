// Adds a "copy" button to every code block, so commands and examples can be
// pasted into a terminal or editor without selecting them by hand.
(() => {
  const ICONS = `
    <svg class="icon-copy" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" aria-hidden="true">
      <rect x="9" y="9" width="12" height="12" rx="2"/>
      <path d="M5 15V5a2 2 0 0 1 2-2h8"/>
    </svg>
    <svg class="icon-done" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" aria-hidden="true">
      <path d="M5 12l5 5L20 7"/>
    </svg>`;

  const COPY_LABEL = "Copiar código";
  const COPIED_LABEL = "Copiado";
  const FEEDBACK_MS = 1500;

  function addButton(pre) {
    const button = document.createElement("button");
    button.type = "button";
    button.className = "copy-button";
    button.innerHTML = ICONS;
    button.setAttribute("aria-label", COPY_LABEL);
    button.title = COPY_LABEL;
    button.addEventListener("click", () => copy(pre, button));
    pre.append(button);
  }

  async function copy(pre, button) {
    const code = pre.querySelector("code") || pre;
    try {
      await navigator.clipboard.writeText(code.innerText.trimEnd());
    } catch {
      return; // Clipboard blocked (insecure origin, permissions): nothing to show.
    }
    button.classList.add("is-copied");
    button.setAttribute("aria-label", COPIED_LABEL);
    setTimeout(() => {
      button.classList.remove("is-copied");
      button.setAttribute("aria-label", COPY_LABEL);
    }, FEEDBACK_MS);
  }

  if (!navigator.clipboard) return;
  document.querySelectorAll("pre").forEach(addButton);
})();
