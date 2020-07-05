# Simple HTML generator with Golang

Generate a simple HTML from template.

**Template file**

index.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Code Practice with Golang</title>
</head>
<body>
    <p>Hello {{.Name}}</p>
    <p>Coupon: {{.Coupon}}</p>
    <p>Amount: {{.Amount}}</p>
    
</body>
</html>
```

**Result of generated file**

index_gen.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Code Practice with Golang</title>
</head>
<body>
    <p>Hello Quan Vu</p>
    <p>Coupon: AIX1000</p>
    <p>Amount: 5000</p>
    
</body>
</html>
```

**Run it**

```shell
go run main.go
```

Enjoy!