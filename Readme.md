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

## 关于benchmark 基准测试

在 Go 中编写[基准测试](https://golang.org/pkg/testing/#hdr-Benchmarks)（benchmarks）是该语言的另一个一级特性，它与编写测试非常相似。



```
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```

你会看到上面的代码和写测试差不多。

`testing.B` 可使你访问隐性命名（cryptically named）`b.N`。

基准测试运行时，代码会运行 `b.N` 次，并测量需要多长时间。

代码运行的次数不会对你产生影响，测试框架会选择一个它所认为的最佳值，以便让你获得更合理的结果。

用 `go test -bench=.` 来运行基准测试。 (如果在 Windows Powershell 环境下使用 `go test -bench="."`)

![](https://gitee.com/gavinyjb/images/raw/master/img/20210409142822.png)

以上结果说明运行一次这个函数需要 123.5 纳秒（在我的电脑上）。这挺不错的，为了测试它运行了 9310923次

注意：基准测试默认是顺序运行的。

## 关于子测试 分组测试

```
func TestSumAllTails(t *testing.T) {
	t.Run("测试一", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("测试二", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9, 11})
		want := []int{2, 20}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
```

> 为了方便理解 此处贴上测试函数SumAllTails()
>
> ```
> func Sum(numbers []int) int {
> 	sum := 0
> 	//for i := 0; i < 5; i++ {
> 	//	sum += number[i]
> 	//}
> 
> 	for _, number := range numbers {
> 		sum += number
> 	}
> 	return sum
> }
> func SumAllTails(numbersToSum ...[]int) (sums []int) {
> 	for _, numbers := range numbersToSum {
> 		sums = append(sums, Sum(numbers[1:]))
> 	}
> 	return sums
> }
> 
> ```

### 与此相关的另一个示例的样例

```
func ExampleSumAllTails() {
	sums := SumAllTails([]int{1, 2}, []int{0, 9, 11})
	fmt.Println(sums)
	// Output: [2 20]
}
```

## 关于测试套件-辅助函数

```
func TestHello(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%q' want '%q'", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'world'", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

}
```

> 我们在这里做了什么？
>
> 我们将断言重构为函数。这减少了重复并且提高了测试的可读性。在 Go 中，你可以在其他函数中声明函数并将它们分配给变量。你可以像调用普通函数一样调用它们。我们需要传入 `t *testing.T`，这样我们就可以在需要的时候令测试代码失败。
>
> `t.Helper()` 需要告诉测试套件这个方法是辅助函数（helper）。通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。这将帮助其他开发人员更容易地跟踪问题。如果你仍然不理解，请注释掉它，使测试失败并观察测试输出。