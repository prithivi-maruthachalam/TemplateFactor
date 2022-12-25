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

    // Extract all options except config
    const optionsObject = (({config, ...obj}) => obj)(options);

    // Overwrite configuration with command line options
    Object.keys(optionsObject).forEach((optionKey) => {
      this.fullConfiguration[optionKey] = optionsObject[optionKey];
    });

    // resolve name
    if (!this.fullConfiguration.name) {
      this.fullConfiguration.name = path.basename(this.absoluteSrcDir);
    }

    // todo: check if name already exists

    this.#getDirTree();
  }

  /**
   * Parses the configuration and returns a dir tree of the source directory
   * 
   * @return {Object} Tree representing the dir structure
   */
  #getDirTree() {
    const excludes = this.fullConfiguration.exclude ?? []; // Copy existing excludes

    // Create new excludes list with reqex for default config file
    this.fullConfiguration.exclude = [
      globToRegexp(path.join(this.absoluteSrcDir, DEFAULT_FILE)),
    ];

    // Convert all glob expressions to regex
    excludes.forEach((globPattern) => {
      this.fullConfiguration.exclude.push(
          globToRegexp(path.join(this.absoluteSrcDir, globPattern.replace(/\/$/, ''))),
      );
    });

    // Get fs tree of source folder
    const fsTree = dirTree(this.absoluteSrcDir,
        {exclude: this.fullConfiguration.exclude},
    );
    
    return fsTree;
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
