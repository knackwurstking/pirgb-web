"use strict";

// TODO: color picker
// TODO: limit the range slider min based on color value picked
// TODO: power toggle component
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
