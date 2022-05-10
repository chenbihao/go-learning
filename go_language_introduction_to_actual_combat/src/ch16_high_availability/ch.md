# 高可用

## 字符串连接


`fmt.Sprintf` > `string +=` > `bytes.Buffer.WriteString` > `strings.Builder.WriteString`


## 面向错误的设计

首先接受系统会错误的事实

- 隔离

![img.png](img.png)

![img_1.png](img_1.png)

![img_2.png](img_2.png)

- 冗余

![img_3.png](img_3.png)

![img_4.png](img_4.png)

- 限流

![img_5.png](img_5.png)

![img_6.png](img_6.png)

![img_7.png](img_7.png)

- 断路器

![img_8.png](img_8.png)

![img_9.png](img_9.png)




## 面向恢复的设计

预知所有失败是不可能的

![img_10.png](img_10.png)

![img_11.png](img_11.png)

![img_12.png](img_12.png)

![img_13.png](img_13.png)



## Chaos Engineering

![img_14.png](img_14.png)

![img_15.png](img_15.png)

```md
- 稳定系统的表现行为
- 尝试真实的问题
- 尝试真实的产品
- 自动化、持续
- 范围可控
```

![img_16.png](img_16.png)


## 结束

![img_17.png](img_17.png)




