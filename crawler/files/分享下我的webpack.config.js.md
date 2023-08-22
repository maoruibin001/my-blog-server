# 分享下我的webpack.config.js 

```javascriptvar path = require('path');
var glob = require('glob');
var webpack = require('webpack');
/*提取样式插件，可以将你的样式提取到单独的css文件里*/
var ExtractTextPlugin = require('extract-text-webpack-plugin');
/*
生成HTML的插件，
具体可以去这里查看https://www.npmjs.com/package/html-webpack-plugin
 */
var HtmlWebpackPlugin = require('html-webpack-plugin');
//提取公共的js库插件
var CommonsChunkPlugin = webpack.optimize.CommonsChunkPlugin;
//压缩混淆插件
var UglifyJsPlugin = webpack.optimize.UglifyJsPlugin;

// const debug = process.env.NODE_ENV !== 'production';
const debug = false;

//获取入口模块对象集合
var getEntry = function (globPath, folderLevel) {
  var files = glob.sync(globPath);
  var entries = {}, entry, moduleName;
  for (var i = 0; i < files.length; i++) {
    entry = files[i];
    moduleName = entry.split('/')[folderLevel];
    entries[moduleName] = './' + entry;
  }
  return entries;
}

var globPath = 'src/**/**/js/index.js';
var entries = getEntry(globPath, 2);
var chunks = Object.keys(entries);

var config = {
  //入口文件输出配置
  entry: entries,
  //入口文件输出配置
  output: {
    path: path.join(__dirname, 'assets'), //输出目录的配置
    // publicPath: './',
    filename: 'js/[name].js',
    chunkFilename: '[id].chunk.js?[chunkhash]',
  },
  module: {
    loaders: [
      {
        test: /\.css$/,
        loader: ExtractTextPlugin.extract('style', 'css')
      },
      {
        test: /\.less$/,
        loader: ExtractTextPlugin.extract('css!less')
      },
      {
        test: /\.html$/,
        loader: "html?-minimize" //避免压缩html,https://github.com/webpack/html-loader/issues/50
      },
      {
        test: /\.(woffwoff2ttfeotsvg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
        loader: 'file-loader?name=fonts/[name].[ext]'
      },
      {
        test: /\.(pngjpe?ggif)$/,
        loader: 'url-loader?limit=8192&name=img/[name]-[hash].[ext]'
      }
    ]
  },
  plugins: [
    new webpack.ProvidePlugin({ //加载zepto
      $: 'zepto'
    }),
    new CommonsChunkPlugin({
      name: 'vendors', // 将公共模块提取，生成名为`vendors`的chunk
      chunks: chunks,
      minChunks: chunks.length // 提取所有entry共同依赖的模块
    }),
    new ExtractTextPlugin('styles/[name].css'), //单独使用link标签加载css并设置路径，相对于output配置中的publickPath
    new UglifyJsPlugin({ //压缩代码
      compress: {
        warnings: false
      },
      except: ['$super', '$', 'exports', 'require', 'App', 'Vue'] //排除关键字
    }),
  ]
}
console.log(__dirname);
var pages = glob.sync('src/**/**/index.html');
pages.forEach(function(page, index) {
  var moduleName = Object.keys(entries)[index];
  var pathname = page.replace(/index.html/g, '');

  var conf = {
    filename: path.resolve(__dirname, 'assets/' + moduleName +'.html'), //生成的html存放路径，相对于path
    template: path.resolve(__dirname, pathname + 'index.html'), //html模板路径
    inject: false, //js插入的位置，true/'head'/'body'/false
    /*
     * 压缩这块，调用了html-minify，会导致压缩时候的很多html语法检查问题，
     * 如在html标签属性上使用{{...}}表达式，很多情况下并不需要在此配置压缩项，
     * 另外，UglifyJsPlugin会在压缩代码的时候连同html一起压缩。
     * 为避免压缩html，需要在html-loader上配置'html?-minimize'，见loaders中html-loader的配置。
     */
    minify: { //压缩HTML文件
     removeComments: true, //移除HTML中的注释
     collapseWhitespace: false //删除空白符与换行符
    }
  };
  if (moduleName in config.entry) {
    // conf.favicon = path.resolve(__dirname, 'src/images/favicon.ico');
    conf.inject = 'body';
    conf.chunks = ['vendors', moduleName];
    conf.hash = true;
  }
  config.plugins.push(new HtmlWebpackPlugin(conf));
});

module.exports = config;```

