const headerBanner = document.getElementById("header-banner");
const headerInfo = document.getElementById("header-info");

function dismissHeaderBanner() {
  toggleHiddenDiv(headerBanner, false);
  toggleHiddenDiv(headerInfo, true);
}

function showHeaderBanner() {
  toggleHiddenDiv(headerBanner, true);
  toggleHiddenDiv(headerInfo, false);
}

function toggleHiddenDiv(divElement, enable = true) {
  if (enable) {
    divElement.classList.remove("hidden");
  } else {
    divElement.classList.add("hidden");
  }
}
