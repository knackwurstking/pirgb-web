"use strict";

import * as events from "./lib/events.js";

import ColorPicker from "./components/color-picker.js";
import RangeSlider from "./components/range-slider.js";
import PowerSwitch from "./components/power-switch.js";

// TODO: from backend: init and startwebsocket handler...
//    ...also handle suspend if possible...
//    ...after suspend get update from server (devices/sections)
// TODO: from backend: get devices and sections stuff
// TODO: update sectionsContainer

/** @type {HTMLDivElement} */
const sectionsContainer = document.getElementById("sectionsContainer");

/** @type {HTMLTemplateElement} */
const template = document.getElementById("sectionCard");

/** @type {DocumentFragment} */
const section = sectionsContainer.appendChild(
  template.content.cloneNode(true).querySelector("fieldset")
);
section.querySelector("legend > span").innerText = "pi-bed";
section.querySelector("legend > code").innerText = "[0]";

/** @type {HTMLDivElement} */
const colorPicker = section.querySelector(".color-picker");
colorPicker.replaceWith(ColorPicker.element);

/** @type {HTMLDivElement} */
const rangeSlider = section.querySelector(".range-slider");
rangeSlider.replaceWith(RangeSlider.element);

/** @type {HTMLDivElement} */
const powerSwitch = section.querySelector(".power-switch");
PowerSwitch.setScale(0.6);
PowerSwitch.setColor(ColorPicker.color);
powerSwitch.replaceWith(PowerSwitch.element);

// NOTE: just debugging
console.log("events.global:", events.global);
