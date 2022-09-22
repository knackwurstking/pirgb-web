"use strict";

export default (function () {
  const data = {};

  /**
   * Our color set
   * @type {string[][]}
   */
  data.colors = [
    [
      "#ff0000", // 255 (ff)
      "#ffff00", // 255 (ff)
      "#00ff00", // 255 (ff)
      "#00ffff", // 255 (ff)
      "#0000ff", // 255 (ff)
      "#ff00ff", // 255 (ff)
      "#ffffff",
    ],
    [
      "#ff3f00", // 63  (3f)
      "#3fff00", // 63  (3f)
      "#00ff3f", // 63  (3f)
      "#003fff", // 63  (3f)
      "#3f00ff", // 63  (3f)
      "#ff003f", // 63  (3f)
      "#ff7fbf", // 7f,bf
    ],
    [
      "#ff7f00", // 127 (7f)
      "#7fff00", // 127 (7f)
      "#00ff7f", // 127 (7f)
      "#007fff", // 127 (7f)
      "#7f00ff", // 127 (7f)
      "#ff007f", // 127 (7f)
      "#7fffbf", // 7f,bf
    ],
    [
      "#ffbf00", // 191 (bf)
      "#bfff00", // 191 (bf)
      "#00ffbf", // 191 (bf)
      "#00bfff", // 191 (bf)
      "#bf00ff", // 191 (bf)
      "#ff00bf", // 191 (bf)
      "#bfff7f", // 7f,bf
    ],
    // ...
    [
      "#ff7f7f", // 7f
      "#ffff7f", // 7f
      "#7fff7f", // 7f
      "#7fffff", // 7f
      "#7f7fff", // 7f
      "#ff7fff", // 7f
      "#bf7fff", // 7f,bf
    ],
    [
      "#ff3f7f", // 3f,7f
      "#3fff7f", // 3f,7f
      "#7fff3f", // 3f,7f
      "#7f3fff", // 3f,7f
      "#3f7fff", // 3f,7f
      "#ff7f3f", // 3f,7f
      "#7fbfff", // 7f,bf
    ],
    [
      "#ff3fbf", // 3f,bf
      "#3fffbf", // 3f,bf
      "#bfff3f", // 3f,bf
      "#bf3fff", // 3f,bf
      "#3fbfff", // 3f,bf
      "#ffbf3f", // 3f,bf
      "#ffbf7f", // 7f,bf
    ],
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
    </div>
  `;

  data.element = (function () {
    const el = document.createElement("div");
    el.classList.add("color-picker-holder");
    el.innerHTML = rootHTML;

    const dd = el.querySelector(".values-dropdown");
    el.onclick = async () => {
      dd.classList.toggle("active");
      if (dd.classList.contains("active")) {
        const el = document.createElement("div");
        el.classList.add("values-dropdown-grid");
        appendColors(el);
        dd.appendChild(el);
      } else {
        while (dd.firstChild) {
          dd.removeChild(dd.firstChild);
        }
      }
    };

    return el;
  })();

  data.color = "#ffffff"; // NOTE: placeholder

  function appendColors(dst) {
    for (let cs of data.colors) {
      for (let c of cs) {
        const el = document.createElement("button");
        el.classList.add("color-block");
        el.style.backgroundColor = c;
        dst.appendChild(el);
      }
    }
  }

  return data;
})();
