import "./style.css";
import Api from "./lib/api";
import { sectionCard } from "./components/section-card";

(async () => {
    document.querySelector("#app").innerHTML = `
        <div id="sectionsContainer"></div>
    `;

    const sectionContainer = document.querySelector("#sectionsContainer");

    const devices = await Api.getDevices();

    // render section cards
    for (const device of devices) {
        for (const section of device.sections) {
            sectionCard(sectionContainer, section);
        }
    }
})();
