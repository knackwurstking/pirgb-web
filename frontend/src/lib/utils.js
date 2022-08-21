/** @param {string} color */
export function hexToColor(color) {
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

  let max = Math.max(...rgbw)
  let min = Math.min(...rgbw)

  if ((max - min <= 5 || max === min) && rgbw.length === 3) {
    return [max, max, max, max]
  } else if (rgbw.length < 3) {
    // fill with zero values
    const newRGBW = []
    for (let x = 0; rgbw.length < 3; x++) {
      newRGBW.push(rgbw[x] !== undefined ? rgbw[x]: 0)
    }
    rgbw = newRGBW
  }

  return [...rgbw, 0].slice(0, 4)
}

/**
 * @param {number[]} rgbw
 */
export function colorToHex(...rgbw) {
  var hex = "#"
  for (let i = 0; i < 3; i++) {
    hex += (rgbw[i] || 0).toString(16).padStart(2, "0")
  }

  return hex
}
