You can use `mock.AnythingOfType("*model.Item")` in the `Create` and `Update` tests because the method signatures for those functions expect a `*model.Item` as one of the arguments.

Here are the method signatures for `Create` and `Update` in the `MockItemStore`:

```go
Create(ctx context.Context, item *model.Item, opts v1.CreateOptions) error
Update(ctx context.Context, item *model.Item, opts v1.UpdateOptions) error
```

As you can see, both methods expect a `*model.Item` as the second argument. Therefore, when setting up the mocks for these methods, you can use `mock.AnythingOfType("*model.Item")` to match the expected type for that argument.

The reason you can't use this specific setup for mocking the `Delete` method is that the method signature in the `MockItemStore` is:

```go
Delete(ctx context.Context, id int, opts v1.DeleteOptions) error
```

It expects an `int` as the second argument (id), not a `*model.Item`. Therefore, you should use the correct type when setting up the mock:

```go
mockItemStore.On("Delete", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("v1.DeleteOptions")).Return(nil)
```

Alternatively, since you know the expected item ID in the test, you can use the specific value:

```go
mockItemStore.On("Delete", mock.Anything, itemID, v1.DeleteOptions{}).Return(nil)
```

Both options should work, but the latter is more precise as it checks that the correct item ID is being passed to the `Delete` method.