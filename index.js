const { createLambda } = require('@now/build-utils/lambda.js');
const glob = require('@now/build-utils/fs/glob.js');
const path = require('path');
const rename = require('@now/build-utils/fs/rename.js');

exports.config = {
  maxLambdaSize: '25mb',
};

exports.build = async ({ files, entrypoint, config }) => {
  // move all user code to 'user' subdirectory
  const userFiles = rename(files, name => path.join('user', name));
  const launcherFiles = await glob('**', path.join(__dirname, 'dist'));
  const zipFiles = { ...userFiles, ...launcherFiles };

  const { port, timeout } = Object.assign({}, config);

  const lambda = await createLambda({
    files: zipFiles,
    handler: 'launcher',
    runtime: 'go1.x',
    environment: {
      NOW_STATIC_BIN_LOCATION: path.join('user', entrypoint),
      // TODO: default or error?
      NOW_STATIC_BIN_PORT: '' + port || '8080',
      NOW_STATIC_BIN_TIMEOUT: '' + timeout || '50',
    },
  });

  return { [entrypoint]: lambda };
};
