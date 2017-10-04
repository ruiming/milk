<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  {{.HtmlHead}}
  <link rel="stylesheet" href="/static/css/index.css"></link>
</head>

<body>
  <div class="container">
      {{.LayoutContent}}
  </div>
  <div>
      {{.SideBar}}
  </div>
  
  {{.Scripts}}
  <script src="/static/js/reload.min.js"></script>
</body>
</html>

