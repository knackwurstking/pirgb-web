"use strict";

export default (function () {
  const data = {};

  const rootHTML = `
    <div class="slider-value-display">100</div>
    <input
      type="range"
      min="5"
      max="100"
      value="100"
      class="slider"
      oninput="
        this.parentElement.querySelector('div').innerText = this.value;
      "
    />
  `;

  data.element = (function () {
    const el = document.createElement("div");
    el.classList.add("slider-container");
    el.innerHTML = rootHTML;

    el.querySelector(".slider").oninput = (ev) => {
      setDisplayValue(ev.currentTarget.value);
    };

    return el;
  })();

  data.value = 100;

  data.setValue = function (value) {
    data.element.querySelector(".slider").value = value;
    setDisplayValue(value);
  };

  function setDisplayValue(value) {
    data.value = parseInt(value);
    data.element.querySelector(".slider-value-display").innerText = value;
  }

  return data;
})();
