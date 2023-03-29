const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  lintOnSave: process.env.NODE_ENV === "development",
  transpileDependencies: true,
  outputDir: "dist",
  assetsDir: "assets",
  publicPath: process.env.PUBLIC_PATH ? process.env.PUBLIC_PATH : "./",
  productionSourceMap: false,
  configureWebpack: (config) => {
    config.devtool = false;
  },
  devServer: {
    host: "0.0.0.0",
    port: 8080,
    open: true,
    proxy: {
      ["/"]: {
        target: `http://121.5.145.152:8080`,
        changeOrigin: true,
        pathRewrite: {
          ["^"]: "",
        },
        ws: false,
      },
    },
    // disableHostCheck: true,
  },
});
