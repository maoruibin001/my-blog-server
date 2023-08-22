# 前端必备技能之flex布局 

css3 flex布局极大的方便了页面的结构布局，目前大多数主流浏览器已经支持这个属性，特别是移动端已经支持得非常好，掌握flex布局已经势在必行了。

flex的历史我这就不讨论了，本篇文章主要分享flex的语法，用简短的文字讲述flex各种语法的用处。

想要容器使用flex布局，只需要在容器上加上一个样式，display:flex。

flex语法分为两种，一种是用在父容器上，一种是用在子元素上。

用在父元素上的属性有：justify-content、align-content、align-items、flex-direction、flex-wrap、flex-flow。

justify-content属性定义了项目在主轴上的对齐方式。

```javascript.box {
  justify-content: flex-start | flex-end | center | space-between | space-around;
}```

flex-start：与主轴的起点对齐。
flex-end：与主轴的终点对齐。
center：与主轴的中点对齐。
space-between：与主轴两端对齐，中间子元素间隔平均分布。
space-around：每个子元素间隔相等的距离。所以，中间的子元素间距比两端的子元素间距大一倍。

align-content属性定义了多根轴线的对齐方式。如果项目只有一根轴线，该属性不起作用。

```javascript.box {
  align-content: flex-start | flex-end | center | space-between | space-around | stretch;
}```

flex-start：与交叉轴的起点对齐。
flex-end：与交叉轴的终点对齐。
center：与交叉轴的中点对齐。
space-between：与交叉轴两端对齐，轴线之间的间隔平均分布。
space-around：每根轴线两侧的间隔都相等。所以，轴线之间的间隔比轴线与边框的间隔大一倍。
stretch（默认值）：轴线占满整个交叉轴。

align-items属性定义项目在交叉轴上如何对齐。

```javascript.box {
  align-items: flex-start | flex-end | center | stretch | base-line 
  }```

flex-start：交叉轴的起点对齐。
flex-end：交叉轴的终点对齐。
center：交叉轴的中点对齐。
baseline: 项目的第一行文字的基线对齐。
stretch（默认值）：如果项目未设置高度或设为auto，将占满整个容器的高度。

flex-direction属性决定主轴的方向（即项目的排列方向）。

```javascript.box {
  flex-direction: row | row-reverse | column | column-reverse
}```

row（默认值）：主轴为水平方向，起点在左端。
row-reverse：主轴为水平方向，起点在右端。
column：主轴为垂直方向，起点在上沿。
column-reverse：主轴为垂直方向，起点在下沿。

flex-wrap属性定义，如果一条轴线排不下，如何换行。

```javascript.box {
    flex-wrap: wrap | no-wrap | wrap-reverse
}```

nowrap（默认）：不换行。
wrap：换行，第一行在上方。
wrap-reverse：换行，第一行在下方。

fex-flow属性是flex-direction属性和flex-wrap属性的简写形式，默认值为row nowrap。

```javascriptflex-flow: <flex-direction> || <flex-wrap>```

用在子元素上的属性有：order、flex-grow、flex-shrink、flex-basis、flex、align-self

order属性定义项目的排列顺序。数值越小，排列越靠前，默认为0。

```javascript.item {
    order: <number> /*default 0*/
}```

flex-grow属性定义项目的放大比例，默认为0，即如果存在剩余空间，也不放大。

```javascript.item {
    flex-grow: <number> /*default 0*/
}```

flex-shrink属性定义了项目的缩小比例，默认为1，即如果空间不足，该项目将缩小。

```javascript.item {
    flex-shrink: <number> /*default 1*/
}```

flex-basis属性定义了在分配多余空间之前，项目占据的主轴空间（main size）。浏览器根据这个属性，计算主轴是否有多余空间。它的默认值为auto，即项目的本来大小。

```javascript.item {
    flex-basis: <length> /*default auto*/
}```

它可以设为跟width或height属性一样的值（比如350px），则项目将占据固定空间。

flex属性是flex-grow、flex-shrink和flex-basis的简写。

```javascript.item {
    flex: none | [ <'flex-grow'> <'flex-shrink'>? || <'flex-basis'> ]
}```

该属性有两个快捷值：auto (1 1 auto) 和 none (0 0 auto)。
建议优先使用这个属性，而不是单独写三个分离的属性，因为浏览器会推算相关值。

align-self属性允许单个项目有与其他项目不一样的对齐方式，可覆盖align-items属性。默认值为auto，表示继承父元素的align-items属性，如果没有父元素，则等同于stretch。

```javascript.item {
  align-self: auto | flex-start | flex-end | center | baseline | stretch;
}```

该属性可能取6个值，除了auto，其他都与align-items属性完全一致。

以上就是flex语法的全部内容，比较少也比较简单，并且很多都有相似之处，只要统一记忆，把相似的东西归类，就能熟练掌握了。thx~

