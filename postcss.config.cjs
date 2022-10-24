const autoprefixer = require('autoprefixer');
const postcssJitProps = require('postcss-jit-props');
const OpenProps = require('open-props');
const nesting = require('postcss-nesting');
const customMedia = require('postcss-custom-media');

const config = {
	plugins: [autoprefixer, postcssJitProps(OpenProps), nesting, customMedia]
};

module.exports = config;
