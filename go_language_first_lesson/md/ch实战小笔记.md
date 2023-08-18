

# 问题备注：

## 基础相关



### 值范围（隐式代码块、值传递）

### 逻辑操作符的顺序

### 零值不可用与零值可用

### 循环变量重用

### map 的随机性

在循环过程中如果对map进行了修改，那么结果和遍历map一样具有随机性

（这种随机仅仅是在创建初始iterator时随机选择一个bucket，key存储在哪里是根据hash值来定的）

### switch 对比





## 实战相关

### web应用协程奔溃问题

panic 无法跨 goroutine 捕获

### 指针相关问题

#### 结构体的值传递问题

```go
// (u LoginUser) 是值传递，方法内只收到了副本，所以外面的对象没收到值
func (u LoginUser) GetOriginHeaderByHttp(req *http.Request) (err error) {
    u.OriginHeader, err = GetOriginHeaderByHttp(req)
    return
}
// (u *LoginUser) 是指针传递，可以同步修改到外面的对象
func (u *LoginUser) GetOriginHeaderByHttp(req *http.Request) (err error) {
    u.OriginHeader, err = GetOriginHeaderByHttp(req)
    return
}
```

#### 值拷贝

go是值传递，包括 range


#### for指针

```go
var all []*Item
for _, item := range items {
    all = append(all, &item) // 会取最后一个值
}

// 正确写法：
for _, item := range items {
    item := item
    all = append(all, &item)
}
```

```go
var prints []func()
for _, v := range []int{1, 2, 3} {
    prints = append(prints, func() { fmt.Println(v) })
}
for _, print := range prints {
    print()
}

// 正确写法：
for _, v := range []int{1, 2, 3} {
    v := v
    prints = append(prints, func() { fmt.Println(v) })
}
```


#### 切片指针

扩容时会分配新的数组，切片会与原数组解除“绑定”

切片排序会影响到原有切片的问题

#### map指针

map 的自动扩容会导致 value 地址变化，所以 Go 不允许获取 map 中 value 的地址


#### 空指针

nil error 值 != nil

### 并发问题

map并发：无并发写保护，1.9引入并发写安全的 [sync.Map](





### defer 的优先级问题）

例如 https://www.cnblogs.com/ricklz/p/9574645.html





# 代码技巧

## 目录结构

## 静态扫描

## 循环依赖

## 错误链处理





# 其他文章分享

go语言最全优化技巧总结，值得收藏！-赵柯 -腾讯大讲堂
https://mp.weixin.qq.com/s/Ux7io_C1ghVLICuDPExHYg



Commit Message：Go 语言第一课-xxx-xxx



