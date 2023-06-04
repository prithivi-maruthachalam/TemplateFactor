const fs = require('fs');
const path = require('path');
const buildFsTree = require('directory-tree');
const {DirectoryDoesNotExist, TemplateExists} = require('./errors');
const {Config} = require('./config');
const {getStorageDirPath, getTemplateFileName} = require('./utils');

/**
 * A dir template class
 * @class
 */
class Template {
  /**
     * @param {string} sourceDirectory path to the the template source
     * @param {Object} options commandline options
     * @constructor
     */
  constructor(sourceDirectory, options) {
    // get absolute path to the source directory
    this.absoluteSourceDirectory = path.resolve(sourceDirectory);

    // throw an error if the directory doesn't exist
    if (!fs.existsSync(this.absoluteSourceDirectory)) {
      throw new DirectoryDoesNotExist(this.absoluteSourceDirectory);
    }

    // checks, validates and creates a new configuration
    this.configuration = new Config(options, this.absoluteSourceDirectory);

    // Check if a template already exists with this name
    const templateFilePath = path.join(
        getStorageDirPath(),
        getTemplateFileName(this.configuration.jsonConfig.name));
    if (fs.existsSync(templateFilePath)) {
      throw new TemplateExists(this.configuration.jsonConfig.name);
    }

    // get a directory tree based on this configuration
    const fsTree = this.#getFsTree(this.configuration.exclude);
    // console.debug('Dir Tree', JSON.stringify(fsTree, null, 2));
    // console.debug('Template Configuration', this.configuration);

    // recurse through the dir tree and configuration and generate the template document
    this.templateJsonObject = this.#makeTemplateJsonObject(fsTree);
    console.debug('Template', JSON.stringify(this.templateJsonObject, null, 2));

    // TODO: Save the template object
    fs.writeFileSync(templateFilePath, JSON.stringify(this.templateJsonObject));
  }

  /* PRIVATE MEMBERS */

  /**
   * Creates and returns a tree of the filestructure of the source
   * folder with the specified patterns excluded
   *
   * @param {string[]} excludes
   * @return {Object} the directory and file structure as a tree
   */
  #getFsTree(excludes) {
    return buildFsTree(this.absoluteSourceDirectory, {
      exclude: excludes,
    });
  }

  /**
   *
   * @param {Object} dTree The directory tree
   * @param {Object} config Template configuration
   * @return {Object} the json representation of the template
   */
  #makeTemplateJsonObject(dTree) {
    const templateObject = {};

    templateObject.name = this.configuration.jsonConfig.name;
    templateObject.fileContentIncluded =
      this.configuration.jsonConfig?.saveFileContent ?? false; // should file content be included in the template
    templateObject.isContentLinked =
      this.configuration.jsonConfig?.optimizeStorage ?? false; // if included, should the content be linked or included in the template itself

    templateObject.filesAndFolders =
      this.#parseAndDeflateTree(dTree.children ?? [], []);

    return templateObject;
  }

  /**
   * Returns a flat list of files and folders to include in the
   * template with the specific files included (through glob patterns)
   *
   * @param {Object} dTreeChildArray
   * @param {Object} resultArray
   * @return {Array} A list of files and folders to include in the template
   */
  #parseAndDeflateTree(dTreeChildArray, resultArray) {
    dTreeChildArray.forEach((child) => {
      const obj = {
        entity: child.path.replace(this.absoluteSourceDirectory + '/', ''),
        isFile: (child.children) ? false : true,
      };

      if (obj.isFile == false) {
        // if this is dir, definitely push
        resultArray.push(obj);
      } else {
        // this is a file

        // check if the file matches and include glob expr from the config
        const isIncludeMatch = this.configuration.jsonConfig.include.some(
            (rx) => rx.test(child.path),
        );

        if (this.configuration.jsonConfig.saveFiles == true || isIncludeMatch) {
          // if the file is to be saved, push

          if (this.configuration.jsonConfig.saveFileContent) {
            // if file content is to be saved
            if (this.configuration.jsonConfig.optimizeStorage == true) {
              // if storage is to be optimised, just add a link
              obj.fileContent = {
                isLink: true,
                data: child.path,
              };
            } else {
              // storage does not have to be optimised
              obj.fileContent = {
                isLink: false,
                data: fs.readFileSync(child.path),
              };
            }
          }

          resultArray.push(obj);
        }
      }

      if (obj.isFile == false) {
        // if this is a folder, recurse for children
        resultArray = this.#parseAndDeflateTree(child.children, resultArray);
      }
    });

    return resultArray;
  }
}

module.exports = {Template};
