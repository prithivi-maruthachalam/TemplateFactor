const fs = require('fs');
const path = require('path');
const {Config} = require('./config');
const buildFsTree = require('directory-tree');
const {DirectoryDoesNotExist} = require('./errors');


/**
 * A template class
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

    // create a new configuration object
    this.configuration = new Config(options, this.absoluteSourceDirectory);

    // TODO: Check if a template already exists with this name

    // TODO: Get a directory tree of the source directory based on configuration
    const fsTree = this.#getFsTree(this.configuration.exclude);
    console.debug(fsTree);

    // TODO: From directory tree and configuration generate the template json object
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
}

module.exports = {Template};
