## CURD
```shell

curl -X POST -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODM5NTc4ODksImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODM4NzE0ODksInN1YiI6ImFkbWluIn0.2B5hcyWln3OqvSd54dlvFKiLhMTApaa4CRoJdX7Ob48' -d '{"asin":"B0B3HXK000","title":"Kuggini Bone Conduction Headphones Bluetooth, Open-Ear Sports Headphones Wireless with Mic, Bluetooth 5.3, IPX6 Waterproof, Hi-Fi Sound Quality, Only 23g, Perfect for Workout, Running, Cycling 00000","sku":"LU-US-000","brand":"Kuggini","product_group":"","product_type":""}' http://127.0.0.1:8883/v2/items

curl -X GET -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODM5NTc4ODksImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODM4NzE0ODksInN1YiI6ImFkbWluIn0.2B5hcyWln3OqvSd54dlvFKiLhMTApaa4CRoJdX7Ob48'  http://127.0.0.1:8883/v2/items

curl -X GET -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODM5NTc4ODksImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODM4NzE0ODksInN1YiI6ImFkbWluIn0.2B5hcyWln3OqvSd54dlvFKiLhMTApaa4CRoJdX7Ob48'  http://127.0.0.1:8883/v2/items/10000

curl -X PUT -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODM5NTc4ODksImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODM4NzE0ODksInN1YiI6ImFkbWluIn0.2B5hcyWln3OqvSd54dlvFKiLhMTApaa4CRoJdX7Ob48' -d '{"asin":"B0B3HXK000","title":"Kuggini Bone Conduction Headphones Bluetooth, Open-Ear Sports Headphones Wireless with Mic, Bluetooth 5.3, IPX7 Waterproof, Hi-Fi Sound Quality, Only 23g, Perfect for Workout, Running, Cycling 000-0001","sku":"LU-US-10000","brand":"Kuggini","product_group":"electronics","product_type":"headphones"}' http://127.0.0.1:8883/v2/items/10000


```

## Instruction
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