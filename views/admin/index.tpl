<table>
  <colgroup>
    <col width="50%">
    <col width="22%">
    <col width="20">
    <col width="20">
    <col width="28%">
    <col width="80">
  </colgroup>
  <thead>
    <tr>
      <th>Title</th>
      <th>Tags</th>
      <th>View</th>
      <th>Comment</th>
      <th>CreatedAt</th>
      <th>Action</th>
    </tr>
  </thead>
  <tbody>
    {{ range .Articles.List }}
    <tr id={{.Id}}>
      <td>{{.Title}}{{.Id}}</td>
      <td>
        {{ range .Tags }}
        <span>{{.Name}}
        </span>
        {{ end }}
      </td>
      <td>{{.View}}</td>
      <td>0</td>
      <td>{{.CreatedAt}}</td>
      <td class="operation">
        <div class="delete-icon"></div>
        <div class="edit-icon"></div>
      </td>
    </tr>
    {{ end }}
  </tbody>
</table>

<section class="page">
  {{if eq .Articles.PageNo 1}}
    <a href="/admin/post/{{.Articles.PageNo}}" class="pre hidden">Pre</a>
  {{else}}
    <a href="/admin/post/{{.Articles.PageNo | decrease }}" class="pre">Pre</a>
  {{end}}

  <span> {{.Articles.PageNo}} / {{.Articles.TotalPage}} </span>

  {{if eq .Articles.PageNo .Articles.TotalPage}}
    <a href="/admin/post/{{.Articles.PageNo}}" class="next hidden">Next</a>
  {{else}}
    <a href="/admin/post/{{.Articles.PageNo | increase }}" class="next">Next</a>
  {{end}}
</section>

<footer>
  <div class="author">
    Official website:
    <a href="http://{{.Website}}">{{.Website}}</a> /
    Contact me:
    <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
  </div>
</footer>
