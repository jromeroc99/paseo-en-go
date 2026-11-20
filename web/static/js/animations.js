// Loads the component behind each animation tag on the page, e.g.
// <pointer-demo> loads animations/pointer-demo.js. A lesson downloads only
// its own animations, and the path is relative to this file, so it works
// under any base URL the site is published with.
const tags = new Set();
for (const element of document.querySelectorAll(":not(:defined)")) {
  tags.add(element.localName);
}

for (const tag of tags) {
  import(new URL(`animations/${tag}.js`, import.meta.url)).catch((error) => {
    // The text fallback inside the tag is still on the page.
    console.error(`animation <${tag}> could not be loaded:`, error);
  });
}
