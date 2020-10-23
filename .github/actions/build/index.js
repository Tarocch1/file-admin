const exec = require('@actions/exec');
const core = require('@actions/core');
const targets = require('../../../targets.json');

(async function () {
  let GO_VERSION = '';
  let BUILD_TIME = '';
  await exec.exec('go', ['version'], {
    listeners: {
      stdout: data => {
        GO_VERSION += data.toString();
      },
    },
  });
  await exec.exec('date', [], {
    listeners: {
      stdout: data => {
        BUILD_TIME += data.toString();
      },
    },
  });
  GO_VERSION = GO_VERSION.replace('go version ', '');
  GO_VERSION = GO_VERSION.replace(/\n/g, '');
  BUILD_TIME = BUILD_TIME.replace(/\n/g, '');
  const ldflags = `-w -s -X 'main.version=${process.env['VERSION']}' -X 'main.goVersion=${GO_VERSION}' -X 'main.buildTime=${BUILD_TIME}' -X 'main.commitID=${process.env['COMMIT_ID']}'`;
  for (let i = 0; i < targets.length; i++) {
    const target = targets[i];
    const env = {
      CGO_ENABLED: '0',
      GOOS: target.os,
      GOARCH: target.arch,
      ...process.env,
    };
    if (target.arm) {
      env['GOARM'] = target.arm;
    }
    if (target.mips) {
      env['GOMIPS'] = target.mips;
      env['GOMIPS64'] = target.mips;
      env['GOMIPSLE'] = target.mips;
      env['GOMIPS64LE'] = target.mips;
    }
    const options = { env };
    core.info(`Starting build ${target.out}`);
    await exec.exec(
      'go',
      ['build', '-o', target.out, '-ldflags', ldflags],
      options,
    );
  }
})();
