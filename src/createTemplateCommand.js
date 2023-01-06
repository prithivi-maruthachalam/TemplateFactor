const fs = require('fs');
const path = require('path');
const globToRegexp = require('glob-to-regexp');
const {DirectoryDoesNotExist, InvalidJsonfile, UnknownError} = require('./errors');
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

    const directoryTree = this.#getDirTree();
    const template = this.#makeTemplateObject(directoryTree);

    console.debug(template);
  }

  #makeTemplateObject(dTree) {
    const templateObject = {};
    templateObject.name = this.fullConfiguration.name;
    templateObject.filesIncluded = this.fullConfiguration.saveFiles ?? false;
    templateObject.fileContentsIncluded = (this.fullConfiguration.saveFiles && this.fullConfiguration.saveFileContent) ?? false;
    templateObject.fileContentType = (!templateObject.fileContentsIncluded) ? 'NONE' : ((this.fullConfiguration.optimizeStorage) ? 'LINK' : 'CONTENT');

    // parse the directory tree and populate the template object
    templateObject.filesAndFolders = this.#parseAndDeflateTree(templateObject, dTree.children ?? [], []);
    return templateObject;
  }

  #parseAndDeflateTree(templateObject, dTreeChildArray, resultArray) {
    dTreeChildArray.forEach((child) => {
      const obj = {
        entity: child.path.replace(this.absoluteSrcDir + '/', ''),
        isFile: (child.children) ? false : true,
      };

      if (obj.isFile == false) {
        resultArray.push(obj);
      } else if (obj.isFile == true && templateObject.filesIncluded == true) {
        if (templateObject.fileContentsIncluded == true) {
          if (templateObject.fileContentType == 'LINK') {
            obj.fileContent = {
              link: child.path,
            };
          } else if (templateObject.fileContentType == 'CONTENT') {
            obj.fileContent = {
              content: fs.readFileSync(child.path),
            };
          }
        }

        resultArray.push(obj);
      }


      if (obj.isFile == false) {
        // if this is a folder, recurse for children
        resultArray = this.#parseAndDeflateTree(templateObject, child.children, resultArray);
      }
    });

    return resultArray;
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
      // if the config file is not specific, recurse with the default file as the parameter
      console.warn(`No config file specified, looking for '${DEFAULT_FILE}' in template srouce directory`);
      return this.#getConfig(path.join(this.absoluteSrcDir, DEFAULT_FILE));
    }

    configFile = path.resolve(configFile); // resolve absolute path for config file

    // Check if file exists
    if (!fs.existsSync(configFile)) {
      if (configFile != path.join(this.absoluteSrcDir, DEFAULT_FILE)) {
        // if the file we're looking for is not the default file, throw an error
        throw new FileDoesNotExist(configFile);
      } else {
        // if it is the default file that we were looking for, issue a warning and continue with default configuration
        console.log('No configuration file found. Using standard configuration');

        // Return empty object if no configuration is available
        return {};
      }
    } else {
      console.log(`Configuration file found at '${configFile}'`);
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
    validateConfig(configObject);

    return configObject;
  }
}

module.exports = {CreateTemplateCommand};
