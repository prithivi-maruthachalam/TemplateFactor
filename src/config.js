const fs = require('fs');
const path = require('path');
const Joi = require('joi');
const checkGlob = require('is-valid-glob');
const globToRegexp = require('glob-to-regexp');
const {FileDoesNotExist, InvalidConfigSchema} = require('./errors');

const DEFAULT_FILE = 'tf.config.json'; // the default name for the config file
// Default configuration values
const STANDARD_CONFIG = {
  saveFiles: false,
  saveFilesContent: false,
  optimizeStorage: false,
};

// Joi Schema for configuration
const CONFIGURATION_SCHEMA = Joi.object({
  name: Joi.string().alphanum(),
  saveFiles: Joi.boolean().default(false),
  saveFileContent: Joi.boolean().default(false),
  optimizeStorage: Joi.boolean().default(false),
  exclude: Joi.array().items(Joi.string().custom((value, helper) => {
    return checkGlob(value) || helper.message(`\'${value}\' is not a valid glob pattern`);
  })).when('saveFiles', {
    is: true,
    then: Joi.optional(),
    otherwise: Joi.forbidden(),
  }),
  include: Joi.array().items(Joi.string().custom((value, helper) => {
    return checkGlob(value) || helper.message(`\'${value}\' is not a valid glob pattern`);
  })).when('saveFiles', {
    is: false,
    then: Joi.optional(),
    otherwise: Joi.forbidden(),
  }),
  excludeContent: Joi.array().items(Joi.string().custom((value, helper) => {
    return checkGlob(value) || helper.message(`\'${value}\' is not a valid glob pattern`);
  })).when('saveFileContent', {
    is: true,
    then: Joi.optional(),
    otherwise: Joi.forbidden(),
  }),
  includeContent: Joi.array().items(Joi.string().custom((value, helper) => {
    return checkGlob(value) || helper.message(`\'${value}\' is not a valid glob pattern`);
  })).when('saveFileContent', {
    is: false,
    then: Joi.optional(),
    otherwise: Joi.forbidden(),
  }),
});


/**
 * Config for the template
 * @class
 */
class Config {
  /**
     * @param {Object} cmdOptions Commandline options
     * @param {string} sourceDirectory Source directory for template
     * @constructor
     */
  constructor(cmdOptions, sourceDirectory) {
    this.absoluteSourceDirectory = sourceDirectory;
    const configFile = cmdOptions.config;

    // resolve configuration file
    const absConfigFilePath = this.#resolveConfigFile(configFile);

    // read and validate configuration
    const userConfig = (absConfigFilePath) ?
        this.#readAndValidateConfig(absConfigFilePath) : STANDARD_CONFIG;

    // combine user configuratino and command line options (priority command line)
    const fullConfig = this.#combineConfiguration(cmdOptions, userConfig);
    if (!fullConfig.name) {
      console.warn('Template name specified, using folder name.');
      this.jsonConfig.name = path.basename(sourceDirectory);
    }

    // convert glob pattern strings to regexp
    this.jsonConfig = this.#convertGlobToRegexp(fullConfig);
  }

  /* PUBLIC MEMBERES */
  /**
   * @return {string[]} array of excluded patterns
   */
  get exclude() {
    return this.jsonConfig.exclude;
  }

  /* PRIVATE MEMBERS */

  /**
   * Resolves the absolute path to the right configuration file
   *
   * @param {string} configFile path to the configuration file
   * @return {string} resolved configuration file location
   */
  #resolveConfigFile(configFile) {
    if (!configFile) {
      // if the config file is not specific, recurse with the default file as the parameter
      console.warn(`No config file specified, looking for '${DEFAULT_FILE}' in template srouce directory`);
      return this.#resolveConfigFile(
          path.join(this.absoluteSourceDirectory, DEFAULT_FILE),
      );
    }

    // resovle absolute path
    const absConfigFile = path.resolve(configFile);

    if (!fs.existsSync(absConfigFile)) {
      if (absConfigFile != path.join(
          this.absoluteSourceDirectory, DEFAULT_FILE,
      )) {
        // if this is the user specified file, throw an error
        throw new FileDoesNotExist(absConfigFile);
      } else {
        // if this is the default file, issue warning and use standard configuration
        console.warn('No configuration file found. Using standard configuration');
      }
    } else {
      // configuration file exists
      console.log(`Configuration file found at '${absConfigFile}'`);
      return absConfigFile;
    }

    return false;
  }

  /**
   * Reads and Validates a configuration file
   *
   * @param {string} configFile fully resolved path to the config file
   * @return {Object} Validated configuration
   */
  #readAndValidateConfig(configFile) {
    if (!fs.existsSync(configFile)) {
      throw new FileDoesNotExist(configFile);
    }

    let configObject = {};

    try {
      const configContent = fs.readFileSync(configFile);
      configObject = JSON.parse(configContent);
    } catch (error) {
      if (error instanceof SyntaxError) {
        throw new InvalidJsonfile(configFile);
      } else {
        throw new UnknownError(error);
      }
    }

    // Validate configuration
    this.#validateConfiguration(configObject);

    return configObject;
  }

  /**
   * Merges the commandline option and the configuration file with
   * a priority for the commandline options.
   *
   * @param {Object} cmdOptions
   * @param {Object} userConfig
   * @return {Object} Merged configuration
   */
  #combineConfiguration(cmdOptions, userConfig) {
    const coreCmdOptions = (({config, ...obj}) => obj)(cmdOptions);

    Object.keys(coreCmdOptions).forEach((optionsKey) => {
      userConfig[optionsKey] = coreCmdOptions[optionsKey];
    });

    return userConfig;
  }

  /**
   * Validates configuration with the Joi schema
   *
   * @param {Object} configuration configuration object
   */
  #validateConfiguration(configuration) {
    try {
      Joi.assert(configuration, CONFIGURATION_SCHEMA);
    } catch (error) {
      throw new InvalidConfigSchema(error.details[0].message);
    }
  }

  /**
   * Convert glob pattern strings to regexp strings
   *
   * @param {Objet} fullConfiguration the merged configuration
   * @return {Object} configuration with regexp strings
   */
  #convertGlobToRegexp(fullConfiguration) {
    /**
     * Convert glob string to regexp string
     * @param {string} globPattern glob pattern string
     * @return {string} regexp string
     */
    const convertGlob = (globPattern) => {
      return globToRegexp(path.join(this.absoluteSourceDirectory, globPattern.replace(/\/$/, '')));
    };

    const exclude = fullConfiguration.exclude ?? [];
    const include = fullConfiguration.include;
    const excludeContent = fullConfiguration.excludeContent;
    const includeContent = fullConfiguration.includeContent;

    // always exclude default configuration file
    fullConfiguration.exclude = [
      globToRegexp(path.join(this.absoluteSourceDirectory, DEFAULT_FILE)),
    ];

    exclude.forEach((globPattern) => {
      fullConfiguration.exclude.push(convertGlob(globPattern));
    });

    if (include) {
      fullConfiguration.include = [];
      include.forEach((globPattern) => {
        fullConfiguration.include.push(convertGlob(globPattern));
      });
    }

    if (excludeContent) {
      fullConfiguration.excludeContent = [];
      excludeContent.forEach((globPattern) => {
        fullConfiguration.excludeContent.push(convertGlob(globPattern));
      });
    }

    if (includeContent) {
      includeContent = [];
      includeContent.forEach((globPattern) => {
        fullConfiguration.includeContent.push(convertGlob(globPattern));
      });
    }

    return fullConfiguration;
  }
}

module.exports = {Config};
