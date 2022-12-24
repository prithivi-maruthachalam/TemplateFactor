const Joi = require('joi');
const checkGlob = require('is-valid-glob');
const {InvalidConfigSchema} = require('./errors');

// Joi Schema for configuration object
const CONFIGURATION_SCHEMA = Joi.object({
  name: Joi.string().alphanum(),
  saveFiles: Joi.boolean().default(false),
  saveFileContent: Joi.boolean().default(false),
  optimizeStorage: Joi.boolean().default(false),
  exclude: Joi.array().has(Joi.string()),
});

/**
 * Validates a configuration object
 *
 * @param {Object} configuration
 */
const validateConfig = (configuration) => {
  try {
    Joi.assert(configuration, CONFIGURATION_SCHEMA);
  } catch (error) {
    throw new InvalidConfigSchema(error.details[0].message);
  }
};

module.exports = {validateConfig};
