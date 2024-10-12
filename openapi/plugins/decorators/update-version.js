module.exports = UpdateVersion

function UpdateVersion() {
    return {
        Info: {
            leave(info, ctx) {
                try {
                    const packageJson = require('../../package.json');
                    info.version = packageJson.version;
                    console.log(`Updating the version to: ${packageJson.version}`)
                }
                catch (err) {
                    ctx.report({
                        message: `Failed to update version. Error was: ${err}`
                    })
                }
            },
        }
    }
}