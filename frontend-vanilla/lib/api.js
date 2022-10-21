import Location from "./location";

/**
 * @typedef PWM
 * @type {{
 *  pulse: number,
 *  color: number[],
 * }}
 */

/**
 * @typedef Devices
 * @type {Device[]}
 */

/**
 * @typedef Device
 * @type {{
 *  host: string,
 *  port: number,
 *  groups: string[],
 *  sections: Section[],
 * }}
 */

/**
 * @typedef Section
 * @type {{
 *  id: number,
 *  pulse: number,
 *  lastPulse: number,
 *  color: number[],
 * }}
 */

export default (function() {
    console.log("[api.js] initializing...");

    const data = {};

    /**
     * @returns {Promise<Devices>}
     */
    data.getDevices = async function() {
        // TODO: load origin from localStorage if possible
        let resp = await fetch(Location.origin + "/api/devices");
        /** @type Devices */
        let devices = [];

        if (resp.status === 200) {
            /** @type {Devices} */
            devices = await resp.json();
        }

        return devices;
    };

    /**
     * @param {string} host
     * @param {number} section
     * @param {PWM} pwmData
     */
    data.setPWM = async function(host, section, pwmData) {
        // TODO: load origin from localStorage if possible
        const resp = await fetch(
            Location.origin + `/api/devices/${host}/${section}/pwm`,
            {
                method: "POST",
                body: JSON.stringify(pwmData),
                headers: {
                    "Content-Type": "application/json",
                },
            }
        );

        if (!resp.ok) {
            throw await data.responseError(resp);
        }

        return;
    };

    /**
     * @param {string} host
     * @param {number} section
     * @returns {Promise<Section>}
     */
    data.getPWM = async function(host, section) {
        // TODO: load origin from localStorage if possible
        const resp = await fetch(
            Location.origin + `/api/devices/${host}/${section}/pwm`
        );

        if (!resp.ok) {
            throw await data.responseError(resp);
        }

        return await resp.json();
    };

    /**
     * @param {Response} resp
     * @returns {Promise<string>}
     */
    data.responseError = async function(resp) {
        return `resp: ${resp.statusText}: ${await resp.text()}`;
    };

    return data;
})();
