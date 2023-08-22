# webpack打包vue2.0项目时必现问题。 

[Vue warn]: You are using the runtime-only build of Vue where the template compiler is not available. Either pre-compile the templates into render functions, or use the compiler-included build.

(found in <Root>)

这个问题是怎么造成的呢，找了很久找不到处理方法，上网查了也没找到一个好的处理方案。后来去看官方文档，找到了类似的答案。

这是什么意思呢?
运行时构建不包含模板编译器，因此不支持 template 选项，只能用 render 选项，但即使使用运行时构建，在单文件组件中也依然可以写模板，因为单文件组件的模板会在构建时预编译为 render 函数。运行时构建比独立构建要轻量30%，只有 17.14 Kb min+gzip大小。
上面一段是官方api中的解释。就是说，如果我们想使用template，我们不能直接在客户端使用npm install之后的vue。此时，再去看查vue模块，添加几行
resolve: {
alias: {
'vue': 'vue/dist/vue.js'
}
}
再运行，没错ok了。

以下是我的完成的代码
webpack.config.babel.js

```javascript/**
 * Created by lenovo on 2017/5/8.
 */
import path from 'path';
import HtmlWebpackPlugin from 'html-webpack-plugin';
const config = {
    entry: './src/index.js',
    output: {
        filename: 'bundle.js',
        path: path.join(__dirname, 'dist')
    },
    module: {
        loaders:[
            {
                test: /\.js$/,
                loader: 'babel'
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            }
        ]
    },
    plugins: [
        new HtmlWebpackPlugin({
            filename: 'index.html',
            template: './index.html',
            title: 'hello App'
        })
    ],
    resolve: {
        alias: {
            'vue': 'vue/dist/vue.js'
        }
    }
}
export default config;```

package.json

```javascript{
  "name": "demo",
  "version": "1.0.0",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "dependencies": {
    "vue": "^2.3.2"
  },
  "devDependencies": {
    "babel-core": "^6.3.26",
    "babel-loader": "^6.2.0",
    "babel-preset-es2015": "^6.24.1",
    "html-webpack-plugin": "^2.28.0",
    "webpack": "^1.12.9",
    "vue-loader": "^12.0.3",
    "vue-template-compiler":"^2.3.2"
  }
}
```

不知道有没有朋友遇到过这样的问题，如果遇到了而你正好不知道怎么解决，我想这篇文章会帮到你。

