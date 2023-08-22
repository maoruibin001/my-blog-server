# webpack.config.js中使用ES6语法 

```javascript es5向es6迁移```

前段时间使用es6习惯了，但是忽然回过头来发现自己的webpack.config.js依旧还在使用require，module.exports，觉得特别别扭，就去网上查阅相关资料。很明显，答案一大片，总结起来就是三点。

```javascript 第一、把webpack.config.js改名为webpack.config.babel.js。

 第二、把增加一个.babelrc的文件，里面写上{ "presets": ["es2015"]}。

 第三、在package.json文件中加上  ```

"devDependencies": {

```javascript        "babel-core": "^6.3.26",
        "babel-loader": "^6.2.0",
        "babel-preset-es2015": "^6.24.1"```

}

然后运行下npm install 一切都搞定。但是，出问题了！
import path from 'path';
^^^^^^
SyntaxError: Unexpected token import

报错了，诶，别人都好好的，为啥我就不行了呢？？很明显，这不科学啊。找资料，看文档，半天也没发现个所以然。然后在一个回答中，发现了这么一行
"webpack": "^1.12.9"；一看自己webpack版本2.0以上，可能是这一个问题。果断加上一试，真行了。这是什么原因？到目前为止我还在研究中，不过到底是把问题先解决了。有记过我会在评论中给出，写这篇文章的目的是给遇到同样问题的朋友一个解决方案。目前的结论是，webpack2.0以上，不能这样实现在webpack.config.js中书写es6语法。

感兴趣的同学可以自己动手去试试~

