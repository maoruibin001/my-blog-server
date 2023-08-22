# 前端体验优化tree-shaking 

## tree-shaking介绍

Tree-shaking 字面意思就是 摇晃树， 其实就是去除那些引用的但却没有使用的代码

## 前提

想要代码配置tree-shaking,必须采用es6的模块语法，因为es6的模块采用的是静态分析，也就是从字面量对代码进行分析。之前的require是动态分析，必须代码执行到才知道引用的什么模块。

## 设置方式

- 一、.babelrc 中添加

```javascript "presets": [
    [
      "es2015", {
        "modules": false,
      }
    ],
    "stage-2"
  ],
```

或者在babel loader中的options里面添加同样的代码。
这个设置的作用是表示不对es6进行处理。

一定要注意是配置es2015的选项，而不是env之类的，否则没有用。

- 二、使用uglifyjs-webpack-plugin，使用时非常简单

```javascriptplugins: [
    new UglifyJsPlugin(),
...
]
```

## 效果

最大代码文件app.js代码大小需要由82k压缩到74k，压缩比例为12%。
gzip 之后由28k减少为27k

注意:

- 网上那些乱七八糟的插件一点用没有，例如webpack-deep-scope-plugin。
- 尽量不要在自己项目的package.json中添加"sideEffects": false

sideEffects: false 的意思并不是我这个模块真的没有副作用（一个函数会、或者可能会对函数外部变量产生影响的行为)，而只是为了在摇树时告诉 webpack：我这个包在设计的时候就是期望没有副作用的，即使他打完包后是有副作用的，webpack 同学你摇树时放心的当成无副作用包摇就好啦！显然在不能确保的情况下，很容易就把项目摇挂了，所以要慎用！

