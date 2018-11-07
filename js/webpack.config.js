
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
                { test: /\.jsx?$/, use: 'babel-loader', exclude: path.join(__dirname, 'node_modules/') },
                { test: /\.s?css$/, use: [ 'style-loader', 'css-loader', 'sass-loader' ]},
            ]
        },
        devServer: {
            contentBase: path.join(__dirname, "dist"),
            compress: true,
            port: 8081,
            proxy: {
                "/graphql": {
                    target: "http://localhost:5000"
                }
            }
        },
        devtool: isProduction ? 'source-map' : 'cheap-module-eval-source-map'
    }
};
