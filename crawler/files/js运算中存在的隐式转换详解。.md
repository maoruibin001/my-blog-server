# js运算中存在的隐式转换详解。 

今天在猿问中看到了两个相似的问题，一个相对简单点，一个相对复杂，我想是不是很多人都存在这样的疑惑呢？所以我把我的回答分享出来，希望能对存在疑惑的的朋友一些帮助。

首先得感谢下@徐锦杰，没有他的问题，也不会有今天这篇文章，谢谢！下面是问题的原文：

```javascript<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title></title>
    <script type="text/javascript">
    function inherit(p){
        if(p==null) throw TypeError();
        if(Object.create) return Object.create(p);
        var t=typeof p;
        if(t!=="object"&& t!=="function") throw TypeError();
        function f(){};
        f.prototype=p;
        return new f();
    }
    function enumeration(namesToValues){
        var enumeration=function(){throw "can't instantiate enumeration"};
        var proto=enumeration.prototype={
            constructor:enumeration,
            toString:function(){return this.name;},
            valueOf:function(){return this.value;},
            toJSON:function(){return this.name;}
        };
        enumeration.values=[];
        for(name in namesToValues){
            var e=inherit(proto); 
            e.name=name;     
            e.value=namesToValues[name]; 
            enumeration[name]=e;
            enumeration.values.push(e);
        }
        enumeration.foreach=function(f,c){
            for (var i=0;i<this.values.length;i++)
                f.call(c,this.values[i]);
        };
        return enumeration;//
    }
    var Coin=enumeration({Penny:1,Nickel:5,Dime:10,Quarter:25});
    var c=Coin.Dime;
    console.log(c instanceof Coin);
    console.log(c.constructor== Coin);
    console.log(c);
    console.log(Coin.Nickel);
    console.log(Coin.values);
    console.log(Coin.Nickel==5); 
    console.log(Coin.Nickel+Coin.Penny);     //==6
    </script>
</head>
<body>
</body>
</html>```

这里Coin.Nickel，Coin.Penny都是类属型，值是一个对象，为什么可以相加呢？

以下是我的回答：

首先要知道Coin.Nickel的值为{name:"Nickel"value:5}，Coin.Penny的值为{name:"Penny"value:1}，其次请看这个对象
var proto=enumeration.prototype={
constructor:enumeration,
toString:function(){return this.name;},
valueOf:function(){return this.value;},
toJSON:function(){return this.name;}
};
标识的是什么意思呢？表示的是对enumeration的原型进行重写。 js中任意两个值进行运算都会存在“隐式转换”，这是什么意思呢，就是运算的时候js解释器会自动给你转化下数据类型。例如： 1 + ‘2’结果是 ‘12’，这就是把数字1转化成了字符串‘1’，然后进行字符串拼接。 那么对象也会转化，对象怎么转化的呢？对象和一个其他类型进行运算时，要先自动调用对象的valueOf方法，如果valueOf方法不能把参与运算的元素转化为相同类型时，再自动调用对象的toStirng方法，这个方法是干什么的呢，正常情况下是把对象转化为字符串，即是转化为"[object Object]",然后再运算。而这里对valueOf方法和toString方法进行了重写，所以，按照上面的规则Coin.Nickel + Coin.Penny实际上是Coin.Nickel.valueOf() +  Coin.Penny.valueOf() 等价于5+1，结果为6.这样答案就出来了！

最后我想补充的是除了字符串和数字想加，采用的是字符串拼接之外，其他都符合基本的js运算规则。

