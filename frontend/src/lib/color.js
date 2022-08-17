/** @param {string} color */
export function hexToRGBW(color, fix = true) {
  // prevent errors
  if (color[0] !== "#" || color.length < 7) return [255, 255, 255, 255]

  /** @type {number[]} */
  let rgbw = []
  color = color.slice(1) // remove "#" from color

  for (let x = 1; x < color.length; x++) {
    if (x % 2 == 1) {
      rgbw.push(parseInt(`${color[x - 1]}${color[x]}`, 16))
    }
  }

  // push W to rgbw
  if (rgbw.length < 4) {
    if (!rgbw.filter(v => v != rgbw[0]).length) {
      rgbw.push(rgbw[0])
    } else {
      rgbw.push(0)
    }
  }

  // TODO: Fix color (ex: [192,191, 188]) - diff <= 5 should be all even (the highest value)

  return rgbw // placeholder
}

/** @param {number[]} rgbw */
export function rgbwToHex(...rgbw) {
  // TODO: convert rgb(w) color to hex

  return "#ffffff" // placeholder
}
