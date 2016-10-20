package foo

//comment
//comment
//comment
//comment
//comment
//comment

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

func (f *Foo) FooMethod() error {
	return nil
}
