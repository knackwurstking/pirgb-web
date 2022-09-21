"use strict";

export default (function () {
  const data = {};

  /**
   * Our color set
   * @type {string[][]}
   */
  data.colors = [
    ["#ff0000", "#ffff00", "#00ff00", "#00ffff", "#0000ff", "#ff00ff"], // 255 (ff)
    ["#ff3f00", "#3fff00", "#00ff3f", "#003fff", "#3f00ff", "#ff003f"], // 63  (3f)
    ["#ff7f00", "#7fff00", "#00ff7f", "#007fff", "#7f00ff", "#ff007f"], // 127 (7f)
    ["#ffbf00", "#bfff00", "#00ffbf", "#00bfff", "#bf00ff", "#ff00bf"], // 191 (bf)
    // ...
    ["#ff7f7f", "#ffff7f", "#7fff7f", "#7fffff", "#7f7fff", "#ff7fff"], // 7f
    ["#ff3f7f", "#3fff7f", "#7fff3f", "#7f3fff", "#3f7fff", "#ff7f3f"], // 3f,7f
    ["#ff3fbf", "#3fffbf", "#bfff3f", "#bf3fff", "#3fbfff", "#ffbf3f"], // 3f,bf
    ["#ff7fbf", "#7fffbf", "#bfff7f", "#bf7fff", "#7fbfff", "#ffbf7f"], // 7f,bf
    ["#ffffff"],
  ];

  const rootHTML = `
    <div class="color-picker-inner">
      <button class="select-color">
        <div>
          <div class="color-block"></div>
          <!-- div class="caret"></div -->
        </div>
      </button>
    </div>

    <div class="values-dropdown">
      <div class="values-dropdown-grid">
      </div>
    </div>
  `;

  const colorGridItemHTML = `
    <div class="values-dropdown-grid">
      <!-- TODO: button is just a placeholder for now -->
      <button class="color-block">
      </button>
      <button class="color-block">
      </button>
      <button class="color-block">
      </button>
      <button class="color-block">
      </button>
      <button class="color-block">
      </button>
      <button class="color-block">
      </button>
      <button class="color-block">
      </button>
      <button class="color-block">
      </button>
      <button class="color-block">
      </button>
      <button class="color-block">
      </button>
      <button class="color-block">
      </button>
      <button class="color-block">
      </button>
    </div>

  `;

  data.element = (function () {
    const el = document.createElement("div");
    el.classList.add("color-picker-holder");
    el.innerHTML = rootHTML;

    const dd = el.querySelector(".values-dropdown");
    el.onclick = () => {
      dd.classList.toggle("active");
      if (dd.classList.contains("active")) {
        dd.innerHTML = colorGridItemHTML;
      } else {
        // remove
        dd.innerHTML = ``;
      }
    };

    // TODO: render colors into '.values-dropdown-grid'

    return el;
  })();

  return data;
})();
