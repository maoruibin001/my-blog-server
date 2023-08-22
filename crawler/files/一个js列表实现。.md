# 一个js列表实现。 

列表简介：下面分享的这个列表是我精心设计的，里面融和了很多js的设计模式，有需要的朋友可以看一看。
/**

- Created by MAORUIBIN on 2016-03-16.
*/
(function(window) {
var win = window,
_restore = {
listSize: function() {
return this.restore.length;
},
length: function() {
console.log(this.restore);
return this.restore.length;
},
clear: function() {
delete this.restore;
this.restore = [];
this.listIndex = 0;
return this;
},
toSting: function() {
return this.restore;
},
getElement: function(ele) {
var index = this.find(ele);
if (index > -1) {
return this.restore[index];
}
return null;
},
insert: function(ele, pos) {
this.listIndex++;
this.restore.splice(pos, 0, ele);
return this;
},
append: function(ele) {
this.listIndex++;
this.restore.push(ele);
return this;
},
remove: function(ele) {
this.listIndex--;
var index = this.find(ele);
this.restore.splice(index, 1);
return this;
},
front: function(ele) {
var index = this.find(ele);
this.restore.splice(index, 1);
this.restore.splice(0, 0, ele);
return this;
},
end: function(ele) {
var index = this.find(ele);
this.restore.splice(index, 1);
this.restore.splice(this.restore.length, 0, ele);
return this;
},
prev: function(ele) {
var index = this.find(ele);
this.restore.splice(index, 1);
this.restore.splice(index-1, 0, ele);
return this;
},
next: function(ele) {
var index = this.find(ele);
console.log(index);
this.restore.splice(index,1);
this.restore.splice(index + 1, 0, ele);
return this;
},
currPos: function(ele) {
return this.find(ele);
},
moveTo: function(ele, pos) {
var index = this.find(ele);
this.restore.splice(index,1);
this.restore.splice(pos, 0, ele)
return this;
},
find: function(ele) {
for (var i = 0, len = this.restore.length; i < len; ++i) {
if (ele === this.restore[i]) {
return i;
}
}
return -1;
}
}
function list() {
return new list.prototype.init();
}
list.fn = list.prototype ={
init: function() {
this.pos = 0;
this.listIndex = 0;
this.restore = [];
this.listSize = _restore.listSize;
this.length = _restore.length;
this.clear = _restore.clear;
this.toString = _restore.toSting;
this.getElement = _restore.getElement;
this.insert = _restore.insert;
this.append = _restore.append;
this.remove = _restore.remove;
this.front = _restore.front;
this.end = _restore.end;
this.prev = _restore.prev;
this.next = _restore.next;
this.currPos = _restore.currPos;
this.moveTo = _restore.moveTo;
this.find = _restore.find;
this._sort = _restore.sort;
}
};
list.fn.init.prototype = list.fn;
win.$ = win.list = list;
win.$$ = list();
})(window)

