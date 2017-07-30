
### 介绍

在终端中显示时，有时需要根据当前终端的宽度（列数，columns）或高度（行数,rows）来进行相应
的显示控制。

### 实现

`Go`中获取终端的行数和列数，在本`package`中，通过两种方式来进行相应的实现：

#### 方法一

借助`stty size` 终端命令，然后在`Go`获取它的输出，再解析出对应的行数与列数

```
$ stty size
output: 34 78 
```

其中的输出格式为：(rows, columns)，然后在相关的接口中进行解析获取即可


#### 方法二

通过`syscall.Syscall`方法，传递系统调用的类型为`syscall.SYS_IOCTL`来访问获取，具体实现看包中相关代码。


### 应用

`import "github.com/knewbie/go-kit/terminal"`

调用相关接口：


```
// Method 1
terminal.Width()    // coloumns
terminal.Height()   // rows

// Method 2
terminal.GetHeight()    // rows
terminal.GetWidth()     // coloumns
```

