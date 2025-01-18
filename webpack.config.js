module.exports = {
  mode: "development",

  entry: "./ts/main.ts",
  output: {
    path: `${__dirname}/static`,
    filename: "script.js",
  },
  module: {
    rules: [{
      test: /\.ts$/,
      use: "ts-loader",
      exclude: /node_modules/
    }]
  },
  resolve: {
    extensions: [".ts", ".js"]
  }
};