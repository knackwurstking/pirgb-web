"use strict";

export default (function () {
  const data = {};

  const rootHTML = `
    <input type="checkbox" />

    <div class="button">
      <svg class="power-off">
        <use xlink:href="#line" class="line" />
        <use xlink:href="#circle" class="circle" />
      </svg>
      <svg class="power-on">
        <use xlink:href="#line" class="line" />
        <use xlink:href="#circle" class="circle" />
      </svg>
    </div>

    <svg
      xmlns="http://www.w3.org/2000/svg"
      style="display: none;"
    >
      <symbol xmlns="http://www.w3.org/2000/svg" viewBox="0 0 150 150" id="line">
        <line x1="75" y1="34" x2="75" y2="58" />
      </symbol>
      <symbol xmlns="http://www.w3.org/2000/svg" viewBox="0 0 150 150" id="circle">
        <circle cx="75" cy="80" r="35" />
      </symbol>
    </svg>
  `;

  data.element = (function () {
    const el = document.createElement("div");
    el.classList.add("power-switch");
    el.style.setProperty("--color-invert", "#000000");
    el.style.transform = "scale(1)";
    el.innerHTML = rootHTML;

    return el;
  })();

  /** @param {string} color */
  data.setColor = function (color) {
    data.element.style.setProperty("--color-invert", color);
  };

  /** @param {number} factor */
  data.setScale = function (factor) {
    data.element.style.transform = `scale(${factor})`;
  };

  return data;
})();
