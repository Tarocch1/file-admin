const core = require('@actions/core');

const ref = core.getInput('ref', { required: true });
const tag = ref.split('/')[2];
core.setOutput('tag', tag);
