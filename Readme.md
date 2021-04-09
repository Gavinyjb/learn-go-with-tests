# 记性太差，做个记事本记一下哈～

## 关于打开文档：

```
运行 
godoc -http=:6060 
并访问 
http://localhost:6060/pkg/。
在这里你能看到 $GOPATH 下所有包的列表，
假如你是在 $GOPATH/src/github.com/{your_id} 下编写的这些代码，你就能在文档中找到它。
```

## 关于示例

你可以通过编写 [示例](https://blog.golang.org/examples) 更深入地了解测试，标准库的文档中能够找到许多这样的例子。

通常示例代码与实际代码所做的工作相比是过时的，就像人们经常忘记更新 readme 文件一样，代码也是如此。

Go 示例执行起来就像测试一样，所以你可以相信示例反映出的是代码的实际功能。

作为包的测试套件的一部分，示例会被编译（并可选择性地执行）。

与典型的测试一样，示例是存在于一个包的 _test.go 文件中的函数。向 `adder_test.go` 文件添加以下 ExampleAdd 函数。

```
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
```

> 为了便于理解，此处粘贴一下Add函数：
>
> ```
> // Add takes two integers and returns the sum of them
> func Add(x, y int) int {
>     return x + y
> }
> ```

## 关于mod初始化

```
go mod init //我一般都默认了
```

> 为了便于理解，我此处插入一个mod样例：
>
> ```
> module github.com/gavin/learn-go-with-tests/iteration
> 
> go 1.16
> 
> ```
>
> 省去的根目录是 ～/go/src 也就是 $GOPATH/src
>
> 因为：
>
> ```
> export GOPATH=$HOME/go
> export PATH=$PATH:$GOPATH/bin
> ```
>
> 在环境变量～/ .bash_profile 中