# golang包
## 1. 包的导入
要在代码中引用其他包的内容，需要使用 import 关键字导入使用的包。具体语法如下：
```golang
package main
import "包的路径"
```
>- 注意事项：
>>- import 导入语句通常放在源码文件开头包声明语句的下面；
>>- 导入的包名需要使用双引号包裹起来；
>>- 包名是从GOPATH/src/ 后开始计算的，使用/ 进行路径分隔。

包的导入有两种写法，分别是**单行导入**和**多行导入**。
### 单行导入
```golang
package main
import "包 1 的路径"
import "包 2 的路径"
```

### 多行导入
```golang
package main
import (
    "包 1 的路径"
    "包 2 的路径"
)

```
## 2.包的导入路径
包的引用路径有两种写法，分别是**全路径导入**和**相对路径导入**。
### 全路径导入
包的绝对路径就是**GOROOT/src/**或**GOPATH/src/**后面包的存放路径，如下所示：
```golang
package main
import "lab/test"
import "database/sql/driver"
import "database/sql"
```
>- 上面代码的含义如下：
>>- test 包是自定义的包，其源码位于GOPATH/src/lab/test 目录下；
>>- driver 包的源码位于GOROOT/src/database/sql/driver 目录下；
>>- sql 包的源码位于GOROOT/src/database/sql 目录下。


### 相对路径导入
相对路径只能用于导入GOPATH 下的包，标准包的导入只能使用全路径导入。

例如包 a 的所在路径是GOPATH/src/lab/a，包 b 的所在路径为GOPATH/src/lab/b，如果在包 b 中导入包 a ，则可以使用相对路径导入方式。示例如下：
```golang
package main
import "../a"
...
```
## 3. 包的引用格式
包的引用有四种格式，下面以 **fmt** 包为例来分别演示一下这四种格式。
1. 标准引用格式
```golang
package main
import "fmt"
```
此时可以用**fmt.**作为前缀来使用 **fmt** 包中的方法，这是常用的一种方式。
```golang
package main
import "fmt"
func main() {
    fmt.Println("Hello GO语言")
}
```
2. 重命名引用格式
   在导入包的时候，我们还可以为导入的包设置别名，如下所示：
```golang
package main
import F "fmt"
```
其中 **F** 就是 **fmt** 包的别名，使用时我们可以使用F.来代替标准引用格式的fmt.来作为前缀使用 **fmt** 包中的方法。
```golang
package main
import F "fmt"
func main() {
    F.Println("Hello GO语言")
}
```
3. 省略引用格式
```golang
package main
import . "fmt"
```
这种格式相当于把 fmt 包直接合并到当前程序中，在使用 fmt 包内的方法是可以不用加前缀fmt.，直接引用。
```golang
package main
import . "fmt"
func main() {
    //不需要加前缀 fmt.
    Println("Hello GO语言")
}
```
4. 匿名引用格式
   在引用某个包时，如果只是希望执行包初始化的 init 函数，而不使用包内部的数据时，可以使用匿名引用格式，如下所示：
```golang
package main
import _ "fmt"
```
匿名导入的包与其他方式导入的包一样都会被编译到可执行文件中。

使用标准格式引用包，但是代码中却没有使用包，编译器会报错。如果包中有 init 初始化函数，则通过import _ "包的路径" 这种方式引用包，仅执行包的初始化函数，即使包没有 init 初始化函数，也不会引发编译器报错。
```golang
package main
import (
    _ "database/sql"
    "fmt"
)
func main() {
    fmt.Println("Hello GO语言")
}
```
>- 注意：
>>- 一个包可以有多个 init 函数，包加载时会执行全部的 init 函数，但并不能保证执行顺序，所以不建议在一个包中放入多个 init 函数，将需要初始化的逻辑放到一个 init 函数里面。
>>- 包不能出现环形引用的情况，比如包 a 引用了包 b，包 b 引用了包 c，如果包 c 又引用了包 a，则编译不能通过。
>>- 包的重复引用是允许的，比如包 a 引用了包 b 和包 c，包 b 和包 c 都引用了包 d。这种场景相当于重复引用了 d，这种情况是允许的，并且 Go 编译器保证包 d 的 init 函数只会执行一次。

## GO包的特点
>- Go语言包的初始化有如下特点：
>>- 包初始化程序从 main 函数引用的包开始，逐级查找包的引用，直到找到没有引用其他包的包，最终生成一个包引用的有向无环图。
>>- Go 编译器会将有向无环图转换为一棵树，然后从树的叶子节点开始逐层向上对包进行初始化。
>>- 单个包的初始化过程如上图所示，先初始化常量，然后是全局变量，最后执行包的 init 函数。

# golang问题解决办法
>- 1. vscode安装工具那里失败可以在cmd输入下面两句
>>- go env -w GO111MODULE=on
>>- go env -w GOPROXY=https://goproxy.io,direct
>>- 然后重启vscode再安装就可以了.
>>- (如果你的go是自定义路径安装需要先修改环境变量)