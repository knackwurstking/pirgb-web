/**
 * @typedef {{
 *  pulse: number,
 *  color: number[],
 * }} PWM
 *
 * @typedef {{
 *  id: number,
 *  pulse: number,
 *  lastPulse: number,
 *  color: number[],
 * }} Section
 *
 * @typedef {{
 *  host: string,
 *  port: number,
 *  groups: string[],
 *  sections: Section[],
 * }} Device
 *
 * @typedef {Device[]} Devices
 */

export default (function () {
  const data = {};

  /**
   * @returns {Promise<Devices>}
   */
  data.getDevices = async function () {
    // TODO: load origin from localStorage if possible
    let resp = await fetch("/api/devices");
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
  data.setPWM = async function (host, section, pwmData) {
    // TODO: load origin from localStorage if possible
    const resp = await fetch(`/api/devices/${host}/${section}/pwm`, {
      method: "POST",
      body: JSON.stringify(pwmData),
      headers: {
        "Content-Type": "application/json",
      },
    });

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
  data.getPWM = async function (host, section) {
    // TODO: load origin from localStorage if possible
    const resp = await fetch(`/api/devices/${host}/${section}/pwm`);

    if (!resp.ok) {
      throw await data.responseError(resp);
    }

    return await resp.json();
  };

  /**
   * @param {Response} resp
   * @returns {Promise<string>}
   */
  data.responseError = async function (resp) {
    return `resp: ${resp.statusText}: ${await resp.text()}`;
  };

  return data;
})();
