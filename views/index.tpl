<header>
    <h1 class="logo">Welcome to Beego</h1>
    <div class="description">
      Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
    </div>
</header>
<section>
  {{ range .Articles.List }}
  <p> {{.Title}} </p>
  {{ end }}
</section>
<footer>
  <div class="author">
    Official website:
    <a href="http://{{.Website}}">{{.Website}}</a> /
    Contact me:
    <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
  </div>
</footer>
