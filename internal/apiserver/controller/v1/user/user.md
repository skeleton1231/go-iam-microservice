```shell
curl -s -XPOST -H'Content-Type: application/json' -d'{"username":"admin","password":"Admin@2021"}' http://127.0.0.1:8883/login | jq -r .token
```

# 创建用户
```shell
curl -s -XPOST -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2MTc5MjI4OTQsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2MTc4MzY0OTQsInN1YiI6ImFkbWluIn0.9qztVJseQ9XwqOFVUHNOtG96-KUovndz0SSr_QBsxAA' -d'{"password":"User@2021","metadata":{"name":"colin"},"nickname":"colin","email":"colin@foxmail.com","phone":"1812884xxxx"}' http://127.0.0.1:8080/v1/users
```

# 列出用户
``` shell
curl -s -XGET -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2MTc5MjI4OTQsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2MTc4MzY0OTQsInN1YiI6ImFkbWluIn0.9qztVJseQ9XwqOFVUHNOtG96-KUovndz0SSr_QBsxAA' 'http://127.0.0.1:8080/v1/users?offset=0&limit=10'
```

# 获取 colin 用户的详细信息
``` shell
curl -s -XGET -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2ODM5NTc4ODksImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2ODM4NzE0ODksInN1YiI6ImFkbWluIn0.2B5hcyWln3OqvSd54dlvFKiLhMTApaa4CRoJdX7Ob48' http://127.0.0.1:8883/v1/users/admin
```

# 修改 colin 用户
``` shell
$ curl -s -XPUT -H'Content-Type: application/json' -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2MTc5MjI4OTQsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2MTc4MzY0OTQsInN1YiI6ImFkbWluIn0.9qztVJseQ9XwqOFVUHNOtG96-KUovndz0SSr_QBsxAA' -d'{"nickname":"colin","email":"colin_modified@foxmail.com","phone":"1812884xxxx"}' http://127.0.0.1:8080/v1/users/colin
```

# 删除 colin 用户
``` shell
$ curl -s -XDELETE -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2MTc5MjI4OTQsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2MTc4MzY0OTQsInN1YiI6ImFkbWluIn0.9qztVJseQ9XwqOFVUHNOtG96-KUovndz0SSr_QBsxAA' http://127.0.0.1:8080/v1/users/colin
```

# 批量删除用户
``` shell
$ curl -s -XDELETE -H'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJpYW0uYXBpLm1hcm1vdGVkdS5jb20iLCJleHAiOjE2MTc5MjI4OTQsImlkZW50aXR5IjoiYWRtaW4iLCJpc3MiOiJpYW0tYXBpc2VydmVyIiwib3JpZ19pYXQiOjE2MTc4MzY0OTQsInN1YiI6ImFkbWluIn0.9qztVJseQ9XwqOFVUHNOtG96-KUovndz0SSr_QBsxAA' 'http://127.0.0.1:8080/v1/users?name=colin&name=mark&name=john'
```