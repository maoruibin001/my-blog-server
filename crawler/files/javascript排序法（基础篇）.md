# javascript排序法（基础篇） 

javascript中实现的排序方法，分为两部分：基础排序方法和高级排序方法，这篇文章讲的是基础排序方法。
基础排序方法有三种：
1.冒泡排序
2.选择排序
3.插入排序
我会一次讲解每一个排序，并给出相应的实现代码，最后还有三个排序的的性能比较。写作这篇文章主要是希望大家能明白js中排序的原理，一遍实现适合自己的排序方法。

冒泡排序：数据会像气泡一样从数组的一端漂到另一端，这是javascript中最容易实现的排序方法。

选择排序：从数组的开头开始，将第一个元素与其他元素进行比较。检查完所有元素后，最小元素会被放到数组的第一个位置，然后算法会从第二个位置开始，当进行到倒数第二个位置时，排序结束。

插入排序：类似于人类按数字或字母顺序对数据进行排序。插入排序有两个循环，外循环将数组元素挨个移动，而内循环对外循环选中的元素及他后面的那个元素进行比较。如果外循环选中元素比内循环中选中元素小，那么数组元素会向右移动，为内循环中这个元素腾出位置。

以下是各种排序方法实现的代码：

/**

- 
Created by MAORUIBIN on 2016-03-17.
*///冒泡排序
(function(window) {
var win = window;
function bubbleSort(arr) {
for(var j = arr.length - 1; j > 0; j--){
for (var i = 0; i < j; ++i) {
if (arr[i] > arr[i + 1]) {
var temp = arr[i];
arr[i] = arr[i + 1];
arr[i + 1] = temp;
}
}
}
return arr;
}
//选择排序
function selectionSort(arr) {
var min, temp,mark;
for (var i = 0, len = arr.length - 1; i < len; ++i) {
min = i;
for (var j = i + 1; j <= len; ++j) {
if (arr[j] < arr[min]) {
min = j;
}
}
temp = arr[i];
arr[i] = arr[min];
arr[min] = temp;
}
return arr;
}
//插入排序
function insertSort(arr) {
var temp, i, j;
for (i = 1; i < arr.length; ++i) {
j = i;
temp = arr[i];
while (j > 0 && arr[j - 1] >= temp) {
arr[j] = arr[j - 1];
--j;
}
arr[j] = temp;
}
return arr;
}
win.bubbleSort = bubbleSort;
win.selectionSort = selectionSort;
win.insertSort = insertSort;
})(window)


Created by MAORUIBIN on 2016-03-17.
*///冒泡排序
(function(window) {
var win = window;

function bubbleSort(arr) {
for(var j = arr.length - 1; j > 0; j--){
for (var i = 0; i < j; ++i) {
if (arr[i] > arr[i + 1]) {
var temp = arr[i];
arr[i] = arr[i + 1];
arr[i + 1] = temp;
}
}
}
return arr;
}

//选择排序
function selectionSort(arr) {
var min, temp,mark;
for (var i = 0, len = arr.length - 1; i < len; ++i) {
min = i;
for (var j = i + 1; j <= len; ++j) {
if (arr[j] < arr[min]) {
min = j;
}
}
temp = arr[i];
arr[i] = arr[min];
arr[min] = temp;
}
return arr;
}

//插入排序
function insertSort(arr) {
var temp, i, j;
for (i = 1; i < arr.length; ++i) {
j = i;
temp = arr[i];
while (j > 0 && arr[j - 1] >= temp) {
arr[j] = arr[j - 1];
--j;
}
arr[j] = temp;
}
return arr;
}

win.bubbleSort = bubbleSort;
win.selectionSort = selectionSort;
win.insertSort = insertSort;
})(window)

经过测试，冒泡排序最慢，其次选择排序，插入排序最快。
如果你有什么问题，可以留言我会及时回复。下一篇javascript排序法（高级篇）。

