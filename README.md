# Notes
## Refs
- https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
- https://gobyexample.com/
- https://awesome-go.com/
- https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/13.5.html
- https://medium.com/a-journey-with-go

#### Testing

```
go test -v -c cover
```

#### Quotes

> Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer

> Interfaces are a great tool for hiding complexity away from other parts of the system. In our case our test helper code did not need to know the exact shape it was asserting on, only how to "ask" for its area.

> Pointers, Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.

> Pointers can be nil. When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.

