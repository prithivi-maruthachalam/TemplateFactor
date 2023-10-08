const path = require('path');

const getStorageDirPath = () => {
  return path.join(process.cwd(), '.templates');
};

const getTemplateFileName = (templateName) => {
  return `${templateName}.json`;
};

module.exports = {getStorageDirPath, getTemplateFileName};
