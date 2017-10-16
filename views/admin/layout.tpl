<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  {{.HtmlHead}}
  <link rel="stylesheet" href="/static/css/index.css"></link>
</head>

<body>
  <header>
      <h1 class="logo">Milk System</h1>
  </header>
  <nav>
    <a href="#">文章</a>
    <a href="#">标签</a>
    <a href="#">评论</a>
    <a href="#">友链</a>
    <a href="#">关于</a>
    <a href="#">配置</a>
  </nav>
  <div class="container">
      {{.LayoutContent}}
  </div>
  <div>
      {{.SideBar}}
  </div>
  
  {{.Scripts}}
  <script src="/static/js/reload.min.js"></script>
  <script src="/static/js/index.js"></script>
</body>
</html>

