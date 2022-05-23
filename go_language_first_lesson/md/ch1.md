[TOC]

# 基础



## 变量声明



### 通用变量声明方法

<img src="ch1.assets/image-20220521154148174.png" style="zoom:50%;" />

### 不赋初值的话默认零值

<img src="ch1.assets/image-20220521154355884.png" />

### 变量声明支持

- 变量声明块（block）

  ``` go
  var ( 
  	a int = 128 
  	...
  )
  ```

- 一行声明多个变量

  ``` go
  var a, b, c int = 5, 6, 7
  ```

- 以上混合

  ``` go
  var ( 
  	a, b, c int = 5, 6, 7
  	c, d, e rune = 'C', 'D', 'E'
  )
  ```



### 语法糖支持

- 省略类型信息声明

  ``` go
  var b = 13
  // 或者不接受默认类型，进行显示类型转型
  var b = int32(13)
  ```

- 短变量声明

  ```
  a := 12
  ```



### 包级变量的声明形式

- 推荐的方式（声明聚类与就近原则，声明一致性）

  ``` go
  // 声明但延迟初始化
  var ( 
  	netGo bool 
  	netCgo bool 
  )
  // 声明并同时显式初始化（声明一致性）
  var (
  	a = 13
  	b = int32(17)
  	f = float32(3.14)
  )
  ```

### 局部变量的声明形式

- 推荐的方式

  ``` go
  // 延迟初始化（采用通用声明）
  var err error
  // 显式初始化（采用短变量）
  a := 17
  f := 3.14
  // 不接受默认类型（采用短变量）
  a := int32(17)
  f := float32(3.14)
  // 分支控制（采用短变量）
  for _, c := range chars {...}
  ```

### 总结

![image-20220521160147066](ch1.assets/image-20220521160147066.png)



## 代码块与作用域



### 代码块与隐式代码块

![image-20220521160323627](ch1.assets/image-20220521160323627.png)



### 预定义标识符（宇宙隐式代码块的标识符）

![image-20220521163148128](ch1.assets/image-20220521163148128.png)

### 包级标识符

- 包顶层声明中的常量、类型、变量或函数（不包括方法）

### 文件代码块标识符

- 导入的包名

### 函数 / 方法体中

- 大括号作为范围界定



tips：控制语句隐式代码块：

```go
// 隐式
if a := 1; false { 
} else if b := 2; false {
}

// 转换成显式

{ // 等价于第一个if的隐式代码块
    a := 1 // 变量a作用域始于此
    if false {

    } else {
        { // 等价于第一个else if的隐式代码块
            b := 2 // 变量b的作用域始于此
            if false {

            }
            // 变量b的作用域终止于此
        }
    }
    // 变量a作用域终止于此
}
```



### 避免变量遮蔽的原则



可能出现的问题：

- 遮蔽预定义标识符
- 遮蔽包代码块中的变量
- 遮蔽外层显式代码块中的变量



tips：短变量声明与控制语句的结合十分容易导致变量遮蔽问题，并且很不容易识别

可以利用工具**辅助**检测变量遮蔽问题，例如`go vet`





## 基本数据类型



### 整型



#### 平台无关整形

![image-20220521164520736](ch1.assets/image-20220521164520736.png)

> go 采用2的补码作为整形的比特位编码方法
>
> ![image-20220521164620263](ch1.assets/image-20220521164620263.png)



#### 平台相关整形

![image-20220521164633640](ch1.assets/image-20220521164633640.png)



#### 整型的溢出问题

```go
var s int8 = 127
s += 1 // 预期128，实际结果-128

var u uint8 = 1
u -= 2 // 预期-1，实际结果255
```

tips：容易发生在循环语句的结束条件判定



#### 字面值与格式化输出

早期版本

```go
a := 53        // 十进制
b := 0700      // 八进制，以"0"为前缀
c1 := 0xaabbcc // 十六进制，以"0x"为前缀
c2 := 0Xddeeff // 十六进制，以"0X"为前缀
```

1.13 增加

```go
d1 := 0b10000001 // 二进制，以"0b"为前缀
d2 := 0B10000001 // 二进制，以"0B"为前缀
e1 := 0o700      // 八进制，以"0o"为前缀
e2 := 0O700      // 八进制，以"0O"为前缀

// 数字分隔符
a := 5_3_7		// 十进制: 537
```

标准库 fmt 包

```go
var a int8 = 59
fmt.Printf("%b\n", a) //输出二进制：111011
fmt.Printf("%d\n", a) //输出十进制：59
fmt.Printf("%o\n", a) //输出八进制：73
fmt.Printf("%O\n", a) //输出八进制(带0o前缀)：0o73
fmt.Printf("%x\n", a) //输出十六进制(小写)：3b
fmt.Printf("%X\n", a) //输出十六进制(大写)：3B
```



### 浮点型

#### 平台无关

- float32
- float64

#### 二进制表示

![image-20220521165909164](ch1.assets/image-20220521165909164.png)

![image-20220521165926887](ch1.assets/image-20220521165926887.png)

#### 转换成二进制例子

```go
有一个整数：139.8125

1.分别把整数部分和小数部分转换成二进制形式：
	整数部分：139d => 10001011b
	小数部分：0.8125d => 0.1101b（乘 2 取整）
	139.8125d=10001011.1101b
2.移动小数点直到只有一个1
	10001011.1101b => 1.00010111101b 
	移动了7位（指数7），尾数为00010111101b
3.计算阶码
	转换过程：阶码 = 指数 + 偏移值
	(偏移值计算：2^(e-1)-1，e=阶码部分的 bit 位数)
	阶码 = 7 + 127 = 134d = 10000110b
4.各自归位（见下图），得到最终二进制表示
	（尾数位数不足 23 位，可在后面补 0）
	最终浮点数 139.8125d 的二进制表示就为：
	0b_0_10000110_00010111101_000000000000
```

![image-20220521170609728](ch1.assets/image-20220521170609728.png)



#### 字面值与格式化输出

字面值

```go
3.1415
.15  // 整数部分如果为0，整数部分可以省略不写
81.80
82. // 小数部分如果为0，小数点后的0可以省略不写

// 十进制科学计数法（e底数为10）
6674.28e-2 // 6674.28 * 10^(-2) = 66.742800
.12345E+5  // 0.12345 * 10^5 = 12345.000000

// 十六进制科学计数法（p底数为2）
0x2.p10  // 2.0 * 2^10 = 2048.000000
0x1.Fp+0 // 1.9375 * 2^0 = 1.937500
```

格式化输出

```go
fmt.Printf("%e\n", f) // 1.234568e+02		十进制科学计数法
fmt.Printf("%x\n", f) // 0x1.edd3be22e5de1p+06	十六进制科学计数法
```



### 复数类型

#### 平台无关

- complex64（实虚都为float32）
- complex128（实虚都为float64）

#### 字面值与格式化输出

```go
// 通过复数字面值初始化
var c = 5 + 6i
var d = 0o123 + .12345E+5i // 83+12345i

// complex 函数
var c = complex(5, 6) // 5 + 6i
var d = complex(0o123, .12345E+5) // 83+12345i

// 预定义函数 real 和 imag
var c = complex(5, 6) // 5 + 6i
r := real(c) // 5.000000
i := imag(c) // 6.000000
```

格式化输出：参考 float



### 自定义数值类型

#### type 关键字

基于原生数值类型

```go
// 例子：
type MyInt int32

// 类型安全规则
var m int = 5
var n int32 = 6
var a MyInt = m // 错误：在赋值中不能将m（int类型）作为MyInt类型使用
var a MyInt = n // 错误：在赋值中不能将n（int32类型）作为MyInt类型使用

// 显示转型解决上面问题
var m int = 5
var n int32 = 6
var a MyInt = MyInt(m) // ok
var a MyInt = MyInt(n) // ok
```

#### 类型别名（初衷是重构）

与type关键字可以互相替换

```go
type MyInt = int32

var n int32 = 6
var a MyInt = n // ok
```



tips：

开发生产中尽量不用浮点

tips2：

```go
// 容易混淆：
type MyInt int32	// 自定义新类型 MyInt
type MyInt = int32	// 与 int32 完全等价，可以直接相互赋值和运算
```



### 字符串类型



> 非原生问题：需要注意类型安全、防止缓冲区溢出、同步问题、获取长度代价大、非ASCII字符支持



自带原生字符串类型：string



优点：

- 数据不可变，提升并发安全性与存储利用率
- 获取长度时间复杂度为常数级
- 所见即所得（不会进行转义）
- 默认采用 Unicode 字符集，UTF-8 编码



字节序列

```go
var s = "中国人"
fmt.Printf("the length of s = %d\n", len(s)) // 9

for i := 0; i < len(s); i++ {
  fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
}
fmt.Printf("\n")
```

字符序列   （Unicode 字符集的码点）

```go
var s = "中国人"
fmt.Println("the character count in s is", utf8.RuneCountInString(s)) // 3

for _, c := range s {
  fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba
}
fmt.Printf("\n")
```



#### rune 类型与字符字面值

一个 rune 示例相当于一个 Unicode 字符，等价于 int32 类型

```go
// 字面值
'a'  // ASCII字符
'中' // Unicode字符集中的中文字符
'\n' // 换行字符
'\'' // 单引号字符

// Unicode 专用转义字符\u 或\U 作为前缀
'\u4e2d'     // 字符：中
'\U00004e2d' // 字符：中
'\u0027'     // 单引号字符

// 本质是整型数
'\x27'  // 使用十六进制表示的单引号字符
'\047'  // 使用八进制表示的单引号字符
```



#### 字符串字面值

```go
"abc\n"
"中国人"
"\u4e2d\u56fd\u4eba" // 中国人
"\U00004e2d\U000056fd\U00004eba" // 中国人
"中\u56fd\u4eba" // 中国人，不同字符字面值形式混合在一起
"\xe4\xb8\xad\xe5\x9b\xbd\xe4\xba\xba" // 十六进制表示的字符串字面值：中国人
```



实验

```go
// rune -> []byte
func encodeRune() { 
    var r rune = 0x4E2D
    fmt.Printf("the unicode charactor is %c\n", r) // 中
    buf := make([]byte, 3)
    _ = utf8.EncodeRune(buf, r) // 对rune进行utf-8编码
    fmt.Printf("utf-8 representation is 0x%X\n", buf) // 0xE4B8AD
} 
// []byte -> rune 
func decodeRune() {
    var buf = []byte{0xE4, 0xB8, 0xAD}         
    r, _ := utf8.DecodeRune(buf) // 对buf进行utf-8解码
    fmt.Printf("the unicode charactor after decoding [0xE4, 0xB8, 0xAD] is %s\n", string(r)) // 中
}
```



内部表示

```go
// $GOROOT/src/reflect/value.go

// StringHeader是一个string的运行时表示
type StringHeader struct {
    Data uintptr
    Len  int
}
```



#### 常见操作

下标操作

```go
var s = "中国人"
fmt.Printf("0x%x\n", s[0]) // 0xe4：字符“中” utf-8编码的第一个字节
```

字符迭代 - 常规 for 迭代

```go
var s = "中国人"
for i := 0; i < len(s); i++ {
  fmt.Printf("index: %d, value: 0x%x\n", i, s[i]) // index: 0, value: 0xe4
}
```

字符迭代 - for range 迭代

```go
var s = "中国人"
for i, v := range s {
    fmt.Printf("index: %d, value: 0x%x\n", i, v) // index: 0, value: 0x4e2d
}
```

字符串连接

```go
s := "Rob Pike, "
s = s + "Robert Griesemer, "
s += " Ken Thompson"
fmt.Println(s) // Rob Pike, Robert Griesemer, Ken Thompson
```

字符串比较 （= =、!= 、>=、<=、> 和 <）

```go
func main() {
        s1 := "世界和平"
        s2 := "世界" + "和平"
        fmt.Println(s1 == s2) // true

        s1 = "Go"
        s2 = "C"
        fmt.Println(s1 != s2) // true

        s1 = "12345"
        s2 = "23456"
        fmt.Println(s1 < s2)  // true  只判断第一个就短路返回
        fmt.Println(s1 <= s2) // true  只判断第一个就短路返回

        s1 = "12345"
        s2 = "123"
        fmt.Println(s1 > s2)  // true  判断到第四个才返回
        fmt.Println(s1 >= s2) // true  判断到第四个才返回
}
```

字符串转换（字符串转换有开销，因为 string 是不可变的，需要分配新内存）

```go
var s string = "中国人"
                      
// string -> []rune
rs := []rune(s) 
fmt.Printf("%x\n", rs) // [4e2d 56fd 4eba]
                
// string -> []byte
bs := []byte(s) 
fmt.Printf("%x\n", bs) // e4b8ade59bbde4baba
                
// []rune -> string
s1 := string(rs)
fmt.Println(s1) // 中国人
                
// []byte -> string
s2 := string(bs)
fmt.Println(s2) // 中国人
```

























































