const fs = require('fs');
const path = require('path');
const globToRegexp = require('glob-to-regexp');
const {DirectoryDoesNotExist, FileDoesNotExist, InvalidJsonfile, UnknownError} = require('./errors');
const {validateConfig} = require('./validation');
const dirTree = require('directory-tree');

const DEFAULT_FILE = 'tf.config.json';

/**
 * Command to create a new template
 * @class
 */
class CreateTemplateCommand {
  /**
     * @constructor
     * @param {String} srcDir Source directory for the template
     * @param {Object} options command line options for the create command
     */
  constructor(srcDir, options) {
    this.absoluteSrcDir = path.resolve(srcDir);
    if (!fs.existsSync(this.absoluteSrcDir)) {
      throw new DirectoryDoesNotExist(this.absoluteSrcDir);
    }

    // Get validated configuration
    this.fullConfiguration = this.#getConfig(options.config);

    // Convert all excluded glob patterns to regexp
    const excludes = this.fullConfiguration.exclude ?? [];
    this.fullConfiguration.exclude = [];
    excludes.forEach((globPattern) => {
      this.fullConfiguration.exclude.push(globToRegexp(globPattern));
    });

    // Extract all options except config
    const optionsObject = (({config, ...obj}) => obj)(options);

    // Overwrite configuration with command line options
    Object.keys(optionsObject).forEach((optionKey) => {
      this.fullConfiguration[optionKey] = optionsObject[optionKey];
    });

    console.debug(this.fullConfiguration);
    // Get fs tree
    const fsTree = dirTree(this.absoluteSrcDir, {
      exclude: this.fullConfiguration.exclude ?? [],
    });
    console.debug(JSON.stringify(fsTree, null, 2));
  }


  /**
   * Fetches and validates configuration
   *
   * @param {String} configFile Path to the configuration file
   * @return {Object} Object representing the configuration
   */
  #getConfig(configFile) {
    if (!configFile) {
      console.warn(`No config file specified, looking for '${DEFAULT_FILE}'`);
      return this.#getConfig(DEFAULT_FILE);
    }

    // Full path to the file to look for
    const absolutePath = path.join(this.absoluteSrcDir, configFile);

    // Check if file exists
    if (!fs.existsSync(absolutePath)) {
      if (configFile != DEFAULT_FILE) {
        throw new FileDoesNotExist(absolutePath);
      } else {
        console.log('No configuration file found. Using standard configuration');

        // Return empty object if no configuration is available
        return {};
      }
    } else {
      console.log(`Configuration file found at '${absolutePath}'`);
    }

    let configObject = {};
    try {
      const configContent = fs.readFileSync(absolutePath);
      configObject = JSON.parse(configContent);
    } catch (error) {
      if (error instanceof SyntaxError) {
        throw new InvalidJsonfile(absolutePath);
      } else {
        throw new UnknownError();
      }
    }

    // Validate configuration
    validateConfig(configObject);

    return configObject;
  }
}

module.exports = {CreateTemplateCommand};
