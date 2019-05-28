<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  <header>
    <h1 class="logo">Welcome to Beego</h1>
    <div class="description">
      Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
    </div>
  </header>
  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop">
      <form action="/" method="post">
            {{ .xsrfdata }}
            <input type="text" name="message"  placeholder="请输入字符串"/>
            <input type="number" name="nums"   placeholder="请输入数字" />
            <input type="submit" value="Post"/>
      </form>
  </div>


</body>
</html>
