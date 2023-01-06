#! /usr/bin/env node

const {Command} = require('commander'); // to parse command line arguments
const {Template} = require('./template');

/**
 * Sets up and parses command line arguments using commander
 */
const setupCmdArgs = () => {
  const cliProgram = new Command(); // set up new command line parser

  cliProgram
      .name('templatefactory')
      .description('CLI tool for creating and using templates')
      .summary('Create, use and manage templates')
      .version('1.0.0');

  cliProgram
      .showHelpAfterError();

  // command : create
  cliProgram
      .command('create')
      .description('create template from folder')
      .argument('<src>', 'source directory for template')
      .option('-n, --name <templateName>', 'name for the new template')
      .option('-cf, --config <file>', 'path to configuration file')
      // flags
      .option('-f, --saveFiles', 'to include files in the template')
      .option('-x, --saveFileContent', 'to include file content in the template')
      .option('-s, --optimizeStorage', 'to optimize storage for storing file content')
      .action((srcDir, options) => {
        new Template(srcDir, options);
      });

  cliProgram.parse();
};

/**
 * main
 */
setupCmdArgs();
