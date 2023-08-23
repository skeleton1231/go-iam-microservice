这是一个基于"Functional Options"的设计模式，该模式是一种在Go中构建和配置复杂对象的方法。它利用了Go中函数和闭包的特性，允许在不使用复杂的构建器或配置结构的情况下，以可读和可扩展的方式创建对象。

**Functional Options模式的优点**：

1. 可扩展性：当需要更多的配置选项时，只需添加更多的函数，而不需要修改现有代码。
2. 可读性：在创建对象时，每个选项都清晰地标明了其用途。
3. 避免冗余：没有必要使用完整的构建器或配置对象。

以下是一个完整示例：

```go
package main

import (
	"fmt"
)

// 1. 定义Person结构体
type Person struct {
	Name       string
	Age        int
	Email      string
	JobFunction func()
}

// 2. 定义Functional Option的类型
type Option func(*Person)

// 3. 为每个属性提供一个WithXXX函数
func WithName(name string) Option {
	return func(p *Person) {
		p.Name = name
	}
}

func WithAge(age int) Option {
	return func(p *Person) {
		p.Age = age
	}
}

func WithEmail(email string) Option {
	return func(p *Person) {
		p.Email = email
	}
}

func WithJobFunction(jobFunc func()) Option {
	return func(p *Person) {
		p.JobFunction = jobFunc
	}
}

// 4. 提供一个创建Person对象的函数，接受任意数量的Option作为参数
func NewPerson(opts ...Option) *Person {
	p := &Person{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func main() {
	// 定义职业功能
	programmerWork := func() {
		fmt.Println("Writing code...")
	}

	teacherWork := func() {
		fmt.Println("Teaching students...")
	}

	// 使用Functional Options创建和配置Person对象
	john := NewPerson(WithName("John"), WithAge(30), WithEmail("john@example.com"), WithJobFunction(programmerWork))
	fmt.Printf("%s's job: ", john.Name)
	john.JobFunction()

	alice := NewPerson(WithName("Alice"), WithAge(28), WithJobFunction(teacherWork))
	fmt.Printf("%s's job: ", alice.Name)
	alice.JobFunction()
}
```

当运行此代码时，你会得到如下输出：

```
John's job: Writing code...
Alice's job: Teaching students...
```

这种方法非常适合于创建和配置具有多种可选设置的对象，并且当对象的设置或配置可能发生变化时，它仍然非常有用，因为不需要对现有的代码进行大量修改。

Go中函数是一等公民 可以结合闭包做非常灵活的设计