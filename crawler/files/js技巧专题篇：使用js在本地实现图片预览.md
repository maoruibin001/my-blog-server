# js技巧专题篇：使用js在本地实现图片预览 

本篇文章主要讨论如何使用js在本地实现图片的预览。这种技巧同样是本应该由后台配合实现的功能，此时采用js的实现方式，在用户网络掉线或本地运行时依然有效，可以减轻服务器压力增强用户体验。当然实现方法有很多，此处只讨论html5的FileReader API，以下是实现方法：

相关html代码如下：

```javascript<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title></title>
    <style>
        div {
            height: 200px;
        }
        img {
            height: 200px;
            display: none;
        }
    </style>
</head>
<body>
<form action="" enctype="multipart/form-data">
    <div id="previewBox">
        <img class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="" id="previewSrc"/>
    </div>
    <input type="file" id="uploadImg"/>
</form>
<script class="lazyload" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXYzh8+PB/AAffA0nNPuCLAAAAAElFTkSuQmCC" data-original="js/previewPic.js"></script>
<script>
    var uploadImg = document.getElementById('uploadImg'),
            previewSrc = document.getElementById('previewSrc');
    previewPic({
        uploadImg: uploadImg,
        previewSrc: previewSrc
    })
</script>
</body>
</html>```

相关js代码如下： 

```javascript/**
 * Created by MAORUIBIN on 2016-03-30.
 */
;(function(window){
    var win = window;
    var previewPic = function(options) {
        var _upload = options.uploadImg,
            _src = options. previewSrc;

        _upload.onchange = function() {
            var self = this;
            var reg = /(.JEPG|.jpeg|.JPG|.jpg|.GIF|.gif|.PNG|.png|.BMP|.bmp){1}/
            var _value = this.value;
            if (!reg.test(_value)) {
                alert('请上传图片格式的文件');
                return false
            }else {
                if (win.FileReader) {
                    var fileReader = new FileReader();
                    _file = this.files[0];
                    fileReader.onload = function(e) {
                        _src.setAttribute('src', this.result);
                        _src.style.display = 'block';
                        return true;
                    }
                    fileReader.onerror = function() {
                        alert('读取文件失败');
                        return false;
                    }
                    //这一步必须，如果没有这一步前面的onload，onerror都不会触发，就像ajax的发送请求一样
                   fileReader.readAsDataURL(_file);
                }else {
                    alert('您的浏览器不支持本地预览');
                    return false;
                }
            }
        }
    }

    win.previewPic = previewPic;
})(window)```

如果有什么疏漏，可以提出来讨论。

