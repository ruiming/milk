<section>
  <ul class="article">
  {{ range .Articles.List }}
    <li> {{.Title}} </li>
  {{ end }}
  </ul>
</section>
<footer>
  <div class="author">
    Official website:
    <a href="http://{{.Website}}">{{.Website}}</a> /
    Contact me:
    <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
  </div>
</footer>
