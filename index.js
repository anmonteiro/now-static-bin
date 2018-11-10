const { createLambda } = require('@now/build-utils/lambda.js');
const FileBlob = require('@now/build-utils/file-blob.js');
const FileFsRef = require('@now/build-utils/file-fs-ref.js');
const fs = require('fs');
const path = require('path');
const { promisify } = require('util');

const fsp = {
  readFile: promisify(fs.readFile)
};

exports.build = async ({ files, entrypoint, config }) => {
  console.log('preparing lambda files...');
  const launcherPath = path.join(__dirname, 'launcher.js');
  let launcherData = await fsp.readFile(launcherPath, 'utf8');

  if (config.port != null) {
    launcherData = launcherData.replace(
      '// BRIDGE_PORT_PLACEHOLDER',
      `bridge.port = ${config.port};`);
  }

  launcherData = launcherData.replace(
    '// PLACEHOLDER',
    `const proc = initProc({command: "./${entrypoint}"});`);

  const launcherFiles = {
    'bridge.js': new FileFsRef({ fsPath: require('@now/node-bridge') }),
    'launcher.js': new FileBlob({
      data: launcherData
    }),
  };

  const lambda = await createLambda({
    files: { ...files, ...launcherFiles },
    handler: 'launcher.launcher',
    runtime: 'nodejs8.10'
  });

  return { [entrypoint]: lambda };
};
