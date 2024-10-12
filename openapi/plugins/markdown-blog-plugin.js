const UpdateVersion = require('./decorators/update-version')

module.exports = function markdownBlogPlugin() {
    return {
        id: 'markdown-blog-plugin',
        decorators: {
            oas3: {
                "update-version": UpdateVersion
            },
        },
    };
};