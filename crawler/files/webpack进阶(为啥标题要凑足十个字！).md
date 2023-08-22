# webpack进阶(为啥标题要凑足十个字！) 

当我们学一个东西时，总不甘心只是学点表面皮毛，了解得越多就越想把他弄清楚。今天我要分享的内容是webpack进阶的内容，相对于前面讲的会更加深入，分为以下几点：

1、webpack热更新

2、公共模块打包

3、css预处理

暂时就分享这三个方面的内容，感兴趣的小伙伴可以看下去。

webpack热更新
热更新是什么意思呢？就是在不刷新页面的情况下，页面内容自动变更。webpack要怎么实现呢？这里有三点关键的地方要注意：
1、必须是单页面里面的某个模块（主页面index.html就不行，为什么？因为没刷新页面）。
2、必须在每个模块入口最前面加上一个热部署模块。以webpack-hot-middleware为例，入口文件必须在前面加上webpack-hot-middleware/client?reload=true，否则不行，很关键！

```javascriptconst webpackBaseConfig = require('./webpack.base.config');

Object.keys(webpackBaseConfig.entry).forEach(function (name) {
  if (name !== 'vonder') {
    webpackBaseConfig.entry[name] = ['webpack-hot-middleware/client?reload=true'].concat(webpackBaseConfig.entry[name])
  }
});```

当然此处也能自己建一个文件，里面require('webpack-hot-middleware/client?reload=true'),  （vue-cli生成的文件就是这么干的），不过看得莫名其妙，不如直接写来的优雅~

3、必须在plugins中加上webpack.HotModuleReplacementPlugin

```javascript plugins: [
    new webpack.HotModuleReplacementPlugin(),
  ]```

做到以上三点，就能做到热模块替换，三点中缺一不可，切记切记！！

公共模块打包

我们使用webpack通常是为了模块管理、压缩打包等功能。假如我们有几个公共js或者第三方js文件，很明显我们不希望由于更改了点页面逻辑的时候让用户从新下载（一般第三方库还不小，自己页面逻辑还比较小）。此时就要用到公共模块打包这个东西了。wbpack提供了一个打包公共模块的插件，webpack.optimize.CommonsChunkPlugin，使用的时候在plugins里面加上就行了：

```javascript plugins: [
     new webpack.optimize.CommonsChunkPlugin({
     name: 'commons/commons',
     }),
 ]```

当然，最好是把要提取的公共模块在入口文件里注明，像这样：

```javascriptentry: {
    app: [resolve('../src/index.js')],
    vendor: ['vue', 'jquery']
  },```

当然，如果你只是这样用，会发现一个问题，每次改动页面逻辑，你的公用common.js文件的hash值也会变，怎么办？两步走，首先打包用的chunkhash，具体问什么可以自行百度。第二在plugins里面再加一个公共模块打包，像这样：

```javascript new webpack.optimize.CommonsChunkPlugin({
     name: 'commons/commons',
   }),

    new webpack.optimize.CommonsChunkPlugin({
      name: 'commons/manifest',
      chunks: ['vendor']
    }),```

这样会生成两个文件，commons.js和vendor.js，你会发现common.js内容就是你引用的第三方包或者公共文件，每次改变页面逻辑commons.js不会变化。但是vendor.js会变化，不过vendor.js比较小，所以相对来说不是太影响，并且vendor.js包含的内容只是打包相关的内容，内容比较少，且直接删除都不影响。当然提取公共模块的时候也可以设置一些策略，例如minChunks、chunks等配置，具体我就不展开讲了。

css预处理
css预处理，可以说是前端自己动手写css必备的工具了，可以省去我们很多功夫。配置也比较简单，有几种方式可以配置，我这里就将两种比较简单的：
1、直接在loader里面配置：

```javascript{
        test: /\.css$/,
        use: ExtraTextWebpackPlugin.extract({
          fallback: "style-loader",
          use: ["css-loader", {
            loader: "postcss-loader",
            options: {
              ident: 'postcss',
              plugins: (loader) => [
                require('autoprefixer')({
                  browsers:["ie >= 8",
                    "Firefox >= 20",
                    "Android > 4.4"]
                }),
              ]
            }
          }]
        }),
        include: resolve('../src')
      },```

第二种，配置一个post.css.config文件
这里要分两步走，第一步现在loader里面配置下：

```javascript{
  test:/\.css$/,
  use:[
    'style-loader','css-loader?importLoaders=1','postcss-loader' 
  ] 
}```

第二步，在你放源文件的目录（我这里是src），下面创建一个postcss.config.js文件。里面加上下面的内容：

```javascriptmodule.exports = {
    plugins: [
        require('autoprefixer')({ /* ...options */ })
    ]
}```

我觉第一种方式比较优雅，第二种方式你要找到源文件目录，不能随便放，也不能配置目录，如果有其他人看你代码，看见这个postcss.config.js简直不明觉厉，一个字，很坑！所以推荐使用第一种方式。

以上就是全部，如果有需要，以后会持续更新，谢谢。

