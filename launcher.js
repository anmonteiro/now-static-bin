const child_process = require('child_process');
const { Bridge } = require('./bridge.js');
const bridge = new Bridge();
// BRIDGE_PORT_PLACEHOLDER

function initProc({ command, args, options }) {
  const proc = child_process.spawn(command, args, options);

  proc.stdout.on('data', (data) => {
    console.log(data.toString('utf8'));
  });

  proc.stderr.on('data', (data) => {
    console.log(data.toString('utf8'));
  });

  proc.on('close', (code) => {
    const error = `child process exited with code ${code}`;
    console.log(error);
    bridge.userError = error;
  });

  proc.on('error', (error) => {
    console.log(error);
    bridge.userError = error;
  });

  return new Promise(function cb(resolve, reject) {
    function ensureRunning() {
      if (proc.pid !== undefined) {
        return resolve();
      }
      return setTimeout(ensureRunning, 50);
    }

    setTimeout(ensureRunning, 50)
  })
}

// PLACEHOLDER

async function launcher(event) {
  await proc;
  return bridge.launcher(event);
}

exports.launcher = launcher;