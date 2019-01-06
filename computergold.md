# 计算机黄金--- 计算机里的小技巧

- 如何判断系统是32位或者是64位
```go
const IS64  = 32 << (^uint(0) >> 63) // 判断是否是64位平台。// 这样就可以照顾是32和64位系统了。
//使用这么个常量，可以很好的解决问题。
```