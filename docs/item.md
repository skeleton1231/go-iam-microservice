## CURD
```shell

curl -X POST -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODQwNTY5OTcsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODM5NzA1OTcsInN1YiI6ImFkbWluIn0.14Yb0ay5Cnsd96UXPu_AodX9E-9NLNVd6j8t-cnYjkI' -d '{"asin":"B0B3HXK001","title":"Kuggini Bone Conduction Headphones Bluetooth, Open-Ear Sports Headphones Wireless with Mic, Bluetooth 5.3, IPX6 Waterproof, Hi-Fi Sound Quality, Only 23g, Perfect for Workout, Running, Cycling 000000001","brand":"Kuggini","product_group":"","product_type":""}' http://127.0.0.1:8883/v2/items

curl -X GET -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODQzODczMDksImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODQzMDA5MDksInN1YiI6ImFkbWluIn0.i0ai2LpHzjJvM0lF9Ld8783LdF4Uilxlix1iKUoJdc0'  http://127.0.0.1:8883/v2/items

curl -X GET -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODM5NTc4ODksImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODM4NzE0ODksInN1YiI6ImFkbWluIn0.2B5hcyWln3OqvSd54dlvFKiLhMTApaa4CRoJdX7Ob48'  http://127.0.0.1:8883/v2/items/10000

curl -X PUT -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODM5NTc4ODksImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODM4NzE0ODksInN1YiI6ImFkbWluIn0.2B5hcyWln3OqvSd54dlvFKiLhMTApaa4CRoJdX7Ob48' -d '{"asin":"B0B3HXK000","title":"Kuggini Bone Conduction Headphones Bluetooth, Open-Ear Sports Headphones Wireless with Mic, Bluetooth 5.3, IPX7 Waterproof, Hi-Fi Sound Quality, Only 23g, Perfect for Workout, Running, Cycling 000-0001","sku":"LU-US-10000","brand":"Kuggini","product_group":"electronics","product_type":"headphones"}' http://127.0.0.1:8883/v2/items/10000

curl -X GET -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODQzODczMDksImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODQzMDA5MDksInN1YiI6ImFkbWluIn0.i0ai2LpHzjJvM0lF9Ld8783LdF4Uilxlix1iKUoJdc0' "http://127.0.0.1:8883/v2/items?fieldSelector=brand=Puggini,title=headphones&offset=0&limit=10"


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

## ItemAttributes

在这种情况下，你的curl命令应该像这样：

1. 创建一个新的项目属性：
   ```bash
   curl -X POST -H "Content-Type: application/json" -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODQ2NTk3ODgsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODQ1NzMzODgsInN1YiI6ImFkbWluIn0.9JQQYLFuhuIm1RG-OzsqJukU46U9vel6u0B96pHydYo' -d '{"item_id": 1657330406107648000, "binding": "Hardcover", "item_height": 23.5, "item_length": 15.0, "item_width": 1.5, "item_weight": 0.8, "item_dimensions_unit": "cm", "package_height": 25.0, "package_length": 16.5, "package_width": 2.0, "package_weight": 1.0, "package_dimensions_unit": "cm", "release_date": "2023-01-01T00:00:00Z"}' http://localhost:8080/api/v2/itemAttris
   ```

2. 更新一个已存在的项目属性：
   ```bash
   curl -X PUT -H "Content-Type: application/json" -d '{"item_id": 1, "binding": "Paperback", "item_height": 23.0, "item_length": 15.0, "item_width": 1.2, "item_weight": 0.7, "item_dimensions_unit": "cm", "package_height": 24.5, "package_length": 16.0, "package_width": 1.7, "package_weight": 0.9, "package_dimensions_unit": "cm", "release_date": "2023-01-02T00:00:00Z"}' http://localhost:8080/api/v2/itemAttris/1
   ```

3. 获取一个已存在的项目属性：
   ```bash
   curl -X GET -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODQwNTY5OTcsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODM5NzA1OTcsInN1YiI6ImFkbWluIn0.14Yb0ay5Cnsd96UXPu_AodX9E-9NLNVd6j8t-cnYjkI' http://localhost:8080/api/v2/itemAttris/1
   ```

4. 删除一个已存在的项目属性：
   ```bash
   curl -X DELETE http://localhost:8080/api/v2/itemAttris/1
   ```

模拟数据可能看起来像这样：

```go
// Mock Data
mockItemAttris := []*v1.ItemAttributes{
	{
		ID:                    1,
		ItemID:                1,
		Binding:               "Hardcover",
		ItemHeight:            23.5,
		ItemLength:            15.0,
		ItemWidth:             1.5,
		ItemWeight:            0.8,
		ItemDimensionsUnit:    "cm",
		PackageHeight:         25.0,
		PackageLength:         16.5,
		PackageWidth:          2.0,
		PackageWeight:         1.0,
		PackageDimensionsUnit: "cm",
		ReleaseDate:           time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	},
	// Add more items as necessary...
}
```