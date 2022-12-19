const chalk = require('chalk');
const path = require('path');
const fs = require('fs');
const {hasOverlappingKeys} = require('./utils');
const cloneDeep = require('lodash.clonedeep');

let ERR_STR = '';
const DEFAULT_FILE = 'tf.config.json';

/**
 * Reads the first available config file and returns the config as an object
 *
 * @param {String} configFilePath
 * @param {String} targetDirectory
 * @return {Object}
 */
const getConfig = (configFilePath, targetDirectory) => {
  console.log('\n');
  let configContent = {};

  if (!configFilePath) {
    console.warn(
        `No config file specified.` +
        ` Looking for ${chalk.bold.blue(DEFAULT_FILE)} config file.`,
    );

    const defaultFilepath = path.join(targetDirectory, DEFAULT_FILE);
    if (
      !fs.existsSync(defaultFilepath) ||
        !fs.lstatSync(defaultFilepath).isFile()) {
      console.warn(
          'No default config file found in template source directory.' +
            ' Using default configuration.',
      );
    } else {
      console.log(`Found config file at ${chalk.bold.blue(DEFAULT_FILE)}`);
      // get the content of the configuration file
      configContent = JSON.parse(fs.readFileSync(defaultFilepath));
    }
  } else {
    if (!fs.existsSync(configFilePath)) {
      // config file does not exist
      ERR_STR = `Config file ${path.resolve(configFilePath)} does not exist`;
      throw new Error(ERR_STR);
    }

    if (!fs.lstatSync(configFilePath).isFile()) {
      ERR_STR = `${path.resolve(configFilePath)} is not a file`;
      throw new Error(ERR_STR);
    }

    // get the content of the configuration file
    configContent = JSON.parse(fs.readFileSync(configFilePath));
  }

  return configContent;
};

const createTemplate = (targetDirectoryPath, options) => {
  // get target directory and validate
  const targetDirectory = path.resolve(targetDirectoryPath);

  if (!fs.existsSync(targetDirectory)) {
    ERR_STR = `Directory ${targetDirectory} does not exist`;
    throw new Error(ERR_STR);
  }

  console.log(
      `Creating template from directory ${chalk.bold.blue(targetDirectory)}`,
  );

  // get config
  const configContent = getConfig(options.config ?? '', targetDirectory);

  // create object without config option
  const cmdOptions = (({config, ...o}) => o)(options);

  if (hasOverlappingKeys(configContent, cmdOptions)) {
    console.log(
        'Command line options will overwrite options from the config file',
    );
  }

  const functionalOptions = cloneDeep(configContent);
  Object.keys(cmdOptions).forEach((optionKey) => {
    functionalOptions[optionKey] = cmdOptions[optionKey];
  });

  console.debug(functionalOptions);
};

module.exports = {createTemplate};
