#! /usr/bin/env node

const {Command} = require('commander'); // to parse command line arguments
const {CreateTemplateCommand} = require('./createTemplateCommand');

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
      .option('-f', 'to include files in the template')
      .option('-x', 'to include file content in the template')
      .option('-s', 'to optimize storage for storing file content')
      .action((srcDir, options) => {
        const createTemplateCommand = new CreateTemplateCommand(
            srcDir,
            options,
        );
        createTemplateCommand.run();
      });

  cliProgram.parse();
};

/**
 * main
 */
setupCmdArgs();
