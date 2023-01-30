

# 问题备注：

## 基础相关

### 值范围（隐式代码块、值传递）

### map 的随机性



## 实战相关

### web应用协程奔溃问题

### 值传递问题

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




