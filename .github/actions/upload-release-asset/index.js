const core = require('@actions/core')
const { GitHub } = require('@actions/github')
const fs = require('fs')
const targets = require('../targets.json')

const github = new GitHub(process.env.GITHUB_TOKEN)

;(async function () {
  try {
    const uploadUrl = core.getInput('upload_url', { required: true })
    const assetContentType = 'application/octet-stream'

    const contentLength = (filePath) => fs.statSync(filePath).size

    for (let i = 0; i < targets.length; i++) {
      const assetPath = targets[i].out
      const assetName = assetPath
      const headers = {
        'content-type': assetContentType,
        'content-length': contentLength(assetPath),
      }
      await github.repos.uploadReleaseAsset({
        url: uploadUrl,
        headers,
        name: assetName,
        file: fs.readFileSync(assetPath),
      })
    }
  } catch (error) {
    core.setFailed(error.message)
  }
})()
