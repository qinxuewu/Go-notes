<!DOCTYPE html>

<html>
<head>
  <title>Beego文件上传</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>
<body>
  <header>
    <h1 class="logo">{{.Website}}</h1>
  </header>
  <div>
     <form enctype="multipart/form-data" method="post" action="/file">
         <input type="file" name="uploadname" />
         <input type="submit">
     </form>
  </div>

</body>
</html>
