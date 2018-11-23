const crypto = require('crypto');
const path = require('path');
const { createLambda } = require('@now/build-utils/lambda.js');
const glob = require('@now/build-utils/fs/glob.js');
const rename = require('@now/build-utils/fs/rename.js');
const objectHash = require('object-hash');

exports.config = {
  maxLambdaSize: '25mb',
};

exports.analyze = ({ files, entrypoint, config }) => {
  const entrypointHash = files[entrypoint].digest;
  const objHash = objectHash(config, { algorithm: 'sha256' });
  const combinedHashes = [entrypointHash, objHash].join('');

  return crypto
    .createHash('sha256')
    .update(combinedHashes)
    .digest('hex');
};

exports.build = async ({ files, entrypoint, config }) => {
  // move all user code to 'user' subdirectory
  const userFiles = rename(files, name => path.join('user', name));
  const launcherFiles = await glob('**', path.join(__dirname, 'dist'));
  const zipFiles = { ...userFiles, ...launcherFiles };

  const { port, timeout } = Object.assign(
    { port: 8080, timeout: 50 },
    config || {},
  );

  const lambda = await createLambda({
    files: zipFiles,
    handler: 'launcher',
    runtime: 'go1.x',
    environment: {
      NOW_STATIC_BIN_LOCATION: path.join('user', entrypoint),
      NOW_STATIC_BIN_PORT: '' + port,
      NOW_STATIC_BIN_TIMEOUT: '' + timeout,
    },
  });

  return { [entrypoint]: lambda };
};
