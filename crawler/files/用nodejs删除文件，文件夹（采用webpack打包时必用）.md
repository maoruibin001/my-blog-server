# 用nodejs删除文件，文件夹（采用webpack打包时必用） 

```javascript使用webpack打包，如果文件内容有修改，那么会出现些重复多余的文件。```

遇到这样的事，很多时候我们会采用gulp +webpack的方式处理，但是如果只是使用gulp的删除文件功能就采用gulp那真的是大炮打蚊子。前两个有朋友问我怎么做，我就把自己用node删除文件的方法共享给了他，稍微改造了下，觉得还行，比较好用，所以分享一下，有需要的朋友可以下载。

//del.js

```javascript/**
 * Created by lenovo on 2017/5/12.
 */
var fs = require('fs');

//获取从命令行传入的参数列表
function getParamList(val, config) {
    var valList = val.split('=');
    if (valList[0] === config) {
        return valList[1].split(',');
    } else {
        return [];
    }
}

//获取从命令行传入的参数列表（去除默认传入的两个）
var agrv = process.argv.slice(2);
if (agrv.length > 0) {
    agrv.forEach(function(val, index, array) {
        var list = getParamList(val, '--targets');
        console.log('list', list);
        list.forEach(function(ele, ind) {
            deleteTarget(ele || './dist');
        })
    });
} else {
    //如果从命令中没有传入参数，则直接默认删除顶级目录下的dist目录。
    deleteTarget('./dist');
}

// 删除目标文件夹或文件
function deleteTarget(fileUrl) {
    // 如果当前url不存在，则退出
    if (!fs.existsSync(fileUrl)) return;
    // 当前文件为文件夹时
    if (fs.statSync(fileUrl).isDirectory()) {
        var files = fs.readdirSync(fileUrl);
        var len = files.length,
            removeNumber = 0;
        if (len > 0) {
            files.forEach(function(file) {
                removeNumber ++;
                var stats = fs.statSync(fileUrl+'/'+file);
                var url = fileUrl + '/' + file;
                if (fs.statSync(url).isDirectory()) {
                    deleteTarget(url);
                } else {
                    fs.unlinkSync(url);
                }

            });
            if (removeNumber === len) {
                // 删除当前文件夹下的所有文件后，删除当前空文件夹（注：所有的都采用同步删除）
                fs.rmdirSync(fileUrl);
                console.log('删除文件夹' + fileUrl + '成功');
            }
        } else {
            fs.rmdirSync(fileUrl)
        }
    } else {
        // 当前文件为文件时
        fs.unlinkSync(fileUrl);
        console.log('删除文件' + fileUrl + '成功');
    }
}
```

//package.json

```javascript{
  "name": "webpack-maoruibin2.0",
  "version": "1.0.1",
  "description": "",
  "author": {
    "name": "maoruibin",
    "email": "595123108@qq.com"
  },
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "start": "webpack-dev-server --hot",
    **"webpack": "node server/delete.js && webpack",
    "webpack:p": "node server/delete.js --targets=./dist,test2 && webpack",**
    "del": "node server/delete.js",
    "build": "babel-node ./node_modules/webpack/bin/webpack"
  },
  "dependencies": {
    "vue": "^2.3.3"
  },
  "devDependencies": {
    "babel-core": "^6.24.1",
    "babel-loader": "^7.0.0",
    "babel-preset-es2015": "^6.24.1",
    "babel-preset-stage-0": "^6.24.1",
    "babel-preset-stage-1": "^6.24.1",
    "babel-preset-stage-2": "^6.24.1",
    "babel-preset-stage-3":"^6.24.1",
    "babel-plugin-transform-class-properties": "^6.24.1",
    "html-webpack-plugin": "^2.28.0",
    "webpack-dev-server": "^2.4.5",
    "webpack": "^2.5.1",
    "glob": "^7.1.1",
    "string-loader": "^0.0.1",
    "style-loader": "^0.17.0",
    "css-loader": "^0.28.1",
    "url-loader": "^0.5.8",
    "file-loader":"^0.11.1",
    "extract-text-webpack-plugin":"^2.1.0",
    "vue-loader":"^12.0.3",
    "vue-template-compiler":"^2.3.3",
    "webpack-chunk-hash": "^0.4.0",
    "inline-manifest-webpack-plugin":"^3.0.1"
  },
  "author": "",
  "license": "ISC"
}
```

