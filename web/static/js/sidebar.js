// Keeps the current lesson in view inside the sidebar: with the whole
// syllabus listed, a lesson near the end would start out of sight.
(() => {
  const sidebar = document.querySelector(".sidebar");
  const current = sidebar?.querySelector('li:has([aria-current="page"])');
  if (!current) return;

  // The sidebar is positioned (sticky or a popover), so offsetTop is the
  // item's distance from the sidebar's top.
  function centerCurrent() {
    sidebar.scrollTop = current.offsetTop - sidebar.clientHeight / 2;
  }

  centerCurrent();
  // On a narrow screen the sidebar is hidden until the menu opens it.
  sidebar.addEventListener("toggle", (event) => {
    if (event.newState === "open") centerCurrent();
  });
})();
