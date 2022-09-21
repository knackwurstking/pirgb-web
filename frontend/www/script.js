"use strict";

/** @type {HTMLDivElement} */
const sectionsContainer = document.getElementById("sectionsContainer");

/** @type {HTMLTemplateElement} */
const template = document.getElementById("sectionCard");

/** @type {DocumentFragment} */
const section = sectionsContainer.appendChild(template.content.cloneNode(true));
console.log("section:", section);

/** @type {HTMLParagraphElement} */
const section2 = sectionsContainer.appendChild(
  template.content.cloneNode(true).querySelector("p")
);
section2.innerText = "What's up!";
console.log("section2:", section2);
