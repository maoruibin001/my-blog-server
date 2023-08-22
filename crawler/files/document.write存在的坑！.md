# document.write存在的坑！ 

今天看猿问中看见一个很有意思的问题，由于篇幅有点长，因此决定将我回答的结果分享出来，希望大家遇到这样问题后知道为什么。下面是对应的代码：

```javascript<!DOCTYPE html>
<html>
<head lang="en">
    <meta charset="UTF-8">
    <title></title>
</head>
<body>
<input type="text" name="i1" value="1"/>
<input type="text" name="i1" value="2"/>
<input type="text" name="i1" value="3"/>
<input type="text" name="i1" value="4"/>
<br/>
<button onclick="total()">kankan</button>
<script>
    function total() {
        var inputs = document.getElementsByName('i1');
            for (var i = 0; i < inputs.length; i ++) {
               document.write(inputs[i].value + "<br>");
            }
    }
</script>
</body>
</html>```

咋一看，这是一个完全正确的代码，可是输出的结果却出人意料。我们想要输出的结果是三排三个value，可实际的输出结果只有一个value，并且是第一个。为什么呢？下面是我的回答：

这里存在两个问题。第一，getElementsByName()返回的不是一个数组，是一个类数组对象。你用Array.isArray()这个方法可以检验。第二，为什么你这里只输出了一个结果，因为document.write(),每输出一次就会刷新一次页面，所以你第一次输出之后，刷新页面。由于getElementsByName()获取的是一个动态集合（即每用一次必须查找一次），但是此时页面已经刷新了，所以没有对应的DOM结构，也就不能再获取到相应的mynode对象，当然就不能输出东西了。综上，你只能输出第一次的结果。此时改进方法有两种。第一，不用document.write输出，改用console.log输出。第二，将类数组对象改为真正的数组，即在mynode下面加一行，mynode = [].slice.call(mynode);后面的不变，也可以输出想要的结果。

不知道有没有朋友也遇到过这样的问题，我之所以将他提出来，最主要的原因是我很少用document.write这个东西，当时觉得是刷新问题，可具体原因还是说不清，百度搜索也找不到相关答案，只能自己想。希望给遇到过相同问题的朋友一个提示，这也算也这篇文章的初衷吧。

