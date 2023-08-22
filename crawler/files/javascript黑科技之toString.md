# javascript黑科技之toString 

今天看co源码时，看到了yield*黑科技。一谈到黑科技，我想起了自己之前用到过两个黑科技：获取数组的最大值和对数组进行排序。

相信很多童鞋在在获取到后台数据，并且要让自己排序的时候，会出现想骂娘的冲动，但是当我们不得不去满足的时候，我们会想到用sort去处理，很明显这很有效。我下面的两个实现采用了一条不同的路，特别是获取最大值时，会让你耳目一新，这两个实现归根到底是对Object.toString的运用。

first： 对数组进行排序（此数组可能是一个基本类型数组，也可能是一个array-object类型）

```javascript/**
 * 数组排序
 * @param  {Array}   arr    待排序数组
 * @param  {string}  key    排序使用的key值，默认为'value';
 * @param  {string}  direct 排序的方向，默认从小到大（可传入任意类型数值，或不传）;
 */
function sort(arr, key, direct) {
    if (!(arr instanceof Array)) return '';

    if (!arr[0]) return arr;

    var tostring = Object.prototype.toString,
        _key = key || 'value';
    if (typeof arr[0] === 'object') {
        //修改Object.prototype.toString返回排序的元素的key值。
        Object.prototype.toString = function() {
            return Number(this[_key]);
        }
    }

    var direct = Number(direct) <= 0 ? -1 : 1;
    arr.sort(function(v1, v2) {
        return direct * (v1 - v2);
    });
    //重置Object.prototype.toString为初始值（必须）。
    Object.prototype.toString = tostring;
    return arr;
}```

second： 获取数组最值（此数组可能是一个基本类型数组，也可能是一个array-object类型）

```javascript/**
 * 获取数组最值
 * @param  {Array}   arr    待获取最值数组
 * @param  {string}  key    待获取最值数组key值，默认为'value';
 * @param  {any}  flag      获取最大值还是最小值，默认获取最大值（可传入任意类型数值，或不传）;
 */

function getMaxOrMin(arr, key, flag) {
    if (!arr) return '';

    if (!arr[0]) return arr;
    var tostring = Object.prototype.toString,
        _key = key || 'value';
        ret = undefined;
    if (typeof arr[0] === 'object') {
        Object.prototype.toString = function() {
            return this[_key];
        }
    }
    if (flag === 'min') {
        flag = 0;
    }
    if (Number(flag) <= 0) {
        ret = Math.min.apply(null, arr);
    } else {
        ret = Math.max.apply(null, arr);
    }

    Object.prototype.toString = tostring;
    return ret;
}```

本文旨在提供一种处理问题的方向，黑科技永远在我们手中，最重要的我们要怎么用。
end： 提供一个更黑的科技。曾经看了一本js书，John Resig写的 好像叫javascript尖刀，具体不确定，全英文的~。里面有讲了一个获取函数形参个数的方法，例如getMaxOrMin。很简单，直接来getMaxOrMin.length。没看到过吧，我在之前也没看到过，所以记得比较清楚，在此分享给大家。

