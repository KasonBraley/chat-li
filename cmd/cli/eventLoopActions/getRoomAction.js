const { question } = require('../utils/readLine.js');
const chalk = require('chalk');

const getRoomType = async () => {
  console.log(
    chalk.red('\n====='),
    chalk.cyan('Choose an Option'),
    chalk.red('=====\n')
  );

  console.log(
    chalk.yellow('1]'),
    chalk.red('---'),
    chalk.cyan('Join a public room'),
    chalk.red('---\n')
  );
  console.log(
    chalk.yellow('2]'),
    chalk.red('---'),
    chalk.cyan('Join a private room'),
    chalk.red('---\n')
  );
  console.log(
    chalk.yellow('3]'),
    chalk.red('---'),
    chalk.cyan('Create a private room'),
    chalk.red('---\n')
  );
  let choice = await question('');

  return choice;
};

module.exports = getRoomType;
