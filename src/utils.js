/**
 * Checks if two objects have overlapping keys
 *
 * @param {Object} a
 * @param {Object} b
 * @return {boolean}
 */
const hasOverlappingKeys = (a, b) => {
  return Object.keys(a).length === Object.keys(b).length &&
    Object.keys(a).every((k) => b.hasOwnProperty(k));
};

module.exports = {hasOverlappingKeys};
