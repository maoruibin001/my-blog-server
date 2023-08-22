# vuepress中使用vue-apollo 

## 直接上代码：

```javascriptconst httpLink = createHttpLink({
  // 你需要在这里使用绝对路径
  uri,
})

// 缓存实现
const cache = new InMemoryCache()

// 创建 apollo 客户端
const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
})

const apolloProvider = new VueApollo({
    defaultClient: apolloClient,
})

export default ({ Vue, options }) => {
    Vue.use(VueApollo)
    // xxx
 }
```

此时会提示错误："TypeError: Cannot read property ‘defaultClient’ of undefined"

解决方法非常简单

```javascriptexport default ({ Vue, options }) => {
    Vue.use(VueApollo)
    // 加上下面这行代码
    Vue.prototype.$apolloProvider = apolloProvider
    // xxx
 }
```

