const Joi = require('joi');

const TEMPLATE_CONFIG = Joi.object();

const validateTemplateConfig = (configuration) => {
  try {
    Joi.assert(configuration, TEMPLATE_CONFIG);
  } catch (error) {

  }
};

module.exports = {validateTemplateConfig};
