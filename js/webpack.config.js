
const HtmlWebpackPlugin = require('html-webpack-plugin');

const path = require('path');

module.exports = (env) => {
    const isProduction = env === "production";
    return {
        entry: './src/index.js',
        mode: isProduction ? 'production': 'development',
        output: {
            path: path.join(__dirname, 'dist'),
            filename: 'js/bundle.js',
            chunkFilename: 'js/[name].js'
        },
        module: {
            rules: [
                { test: /\.js$/, use: 'babel-loader', exclude: path.join(__dirname, 'node_modules/') },
                { test: /\.s?css$/, use: [ 'style-loader', 'css-loader', 'sass-loader' ]},
            ]
        },
        plugins: [
            new HtmlWebpackPlugin({
                template: 'src/index.ejs',
                filename: 'index.html',
                inject: true,
            }),
        ],
        devServer: {
            contentBase: path.join(__dirname, "dist"),
            compress: true,
            port: 8080,
        },
        devtool: isProduction ? 'source-map' : 'cheap-module-eval-source-map'
    }
};
