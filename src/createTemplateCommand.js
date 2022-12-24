const fs = require('fs');
const path = require('path');
const {DirectoryDoesNotExist, FileDoesNotExist, InvalidJsonfile, UnknownError} = require('./errors');
const {validateConfig} = require('./validation');

const DEFAULT_FILE = 'tf.config.json';

/**
 * Command to create a new template
 */
class CreateTemplateCommand {
  /**
     * CreateTemplateCommand constructor
     * @param {String} srcDir source directory for the template
     * @param {Object} options command line options for the create command
     */
  constructor(srcDir, options) {
    this.absoluteSrcDir = path.resolve(srcDir);
    if (!fs.existsSync(this.absoluteSrcDir)) {
      throw new DirectoryDoesNotExist(this.absoluteSrcDir);
    }

    // get and validate config
    this.configuration = this.#getConfig(options.config);
    console.log(this.configuration);
  }


  /**
   * Fetches and Validates configuration
   *
   * @param {String} configFile path to the configuration file
   * @return {Object} Object representing the configuration
   */
  #getConfig(configFile) {
    if (!configFile) {
      console.warn(`No config file specified, looking for '${DEFAULT_FILE}'`);
      return this.#getConfig(DEFAULT_FILE);
    }

    // full path to the file to look for
    const absolutePath = path.join(this.absoluteSrcDir, configFile);

    // check if file exists
    if (!fs.existsSync(absolutePath)) {
      if (configFile != DEFAULT_FILE) {
        throw new FileDoesNotExist(absolutePath);
      } else {
        console.log('No configuration file found. Using standard configuration');

        // return empty object if no configuration is available
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

    // validate configuration
    validateConfig(configObject);

    return configObject;
  }

  /**
   * Creates the template
   */
  run() {
  }
}

module.exports = {CreateTemplateCommand};
