/**
 * An unknown error occurred
 */
class UnknownError extends Error {
  /**
   *
   */
  constructor() {
    super(`An unknown error occurred`);

    // use the name of the class as the name of the error
    this.name = this.constructor.name;

    this.applicationErrorCode = 0;
  }
}

/**
 * Folder does not exist
 */
class DirectoryDoesNotExist extends Error {
  /**
     * @param {String} folder path to the folder
     */
  constructor(folder) {
    super(`The folder \'${folder}\' does not exist`);

    // use the name of the class as the name of the error
    this.name = this.constructor.name;

    this.applicationErrorCode = 1;
    this.path = folder;
  }
}

/**
 * File does not exist
 */
class FileDoesNotExist extends Error {
  /**
     * @param {String} file path to the file
     */
  constructor(file) {
    super(`The file \'${file}\' does not exist`);

    // use the name of the class as the name of the error
    this.name = this.constructor.name;

    this.applicationErrorCode = 2;
    this.path = file;
  }
}

/**
 * Invalid json in file
 */
class InvalidJsonfile extends Error {
  /**
     * @param {String} file path to the file
     */
  constructor(file) {
    super(`The file \'${file}\' does not contain valid json data`);

    // use the name of the class as the name of the error
    this.name = this.constructor.name;

    this.applicationErrorCode = 3;
    this.path = file;
  }
}

/**
 * Invalid configuration schema
 */
class InvalidConfigSchema extends Error {
  /**
   * @param {String} validationMessage error message from the validator
   */
  constructor(validationMessage) {
    super(`The configuration is invalid | ${validationMessage}`);

    // use the name of the class as the name of the error
    this.name = this.constructor.name;

    this.applicationErrorCode = 4;
    this.details = validationMessage;
  }
}

module.exports = {
  DirectoryDoesNotExist,
  FileDoesNotExist,
  InvalidJsonfile,
  UnknownError,
  InvalidConfigSchema,
};
