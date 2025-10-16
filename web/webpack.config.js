const path = require('path');

module.exports = {
  mode: process.env.NODE_ENV || 'development',
  entry: './react/index.tsx',
  output: {
    filename: 'react-bundle.js',
    path: path.resolve(__dirname, 'static'),
    clean: false,
  },
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
      {
        test: /\.css$/i,
        use: ['style-loader', 'css-loader'],
      },
    ],
  },
  resolve: {
    extensions: ['.tsx', '.ts', '.js'],
  },
  devtool: process.env.NODE_ENV === 'production' ? false : 'source-map',
  optimization: {
    minimize: process.env.NODE_ENV === 'production',
  },
  target: ['web', 'es5'],
};
