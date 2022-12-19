#! /usr/bin/env node

const {Command} = require('commander');
const {createTemplate} = require('./templates');


/**
 * Sets up and parses command line arguments using commander
 *
 * @return {Command}
 */
const setupCmdArgs = () => {
  const cliProgram = new Command();

  cliProgram
      .name('templatefactory')
      .description('CLI tool for creating and using templates')
      .summary('Create, use and manage templates')
      .version('1.0.0');

  cliProgram
      .showHelpAfterError();

  // create-template command
  cliProgram
      .command('create')
      .description('create template from folder')
      .argument('<targetDirectory>', 'source directory for template')
      .option('-n, --name <templateName>', 'name for the new template')
      .option('-f, --saveFiles', 'to include files in the template or not')
      .option('-x, --saveFileContent',
          'to include file content in the template or not')
      .option('-cf, --config <file>', 'path to configuration file')
      .action((targetDirectory, options, command) => {
        return createTemplate(targetDirectory, options);
      });

  cliProgram.parse();

  return cliProgram;
};

/**
 * main
 */
console.log('\n');
setupCmdArgs();
console.log('\n');
