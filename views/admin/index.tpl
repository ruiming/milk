<table>
  <colgroup>
    <col width="20">
    <col width="50%">
    <col width="25%">
    <col width="20">
    <col width="20">
    <col width="25%">
  </colgroup>
  <thead>
    <tr>
      <th></th>
      <th>Title</th>
      <th>Tags</th>
      <th>View</th>
      <th>Comment</th>
      <th>CreatedAt</th>
    </tr>
  </thead>
  <tbody>
    {{ range .Articles.List }}
    <tr>
      <td>
        <input type="checkbox" />
      </td>
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
    </tr>
    {{ end }}
  </tbody>
</table>

<footer>
  <div class="author">
    Official website:
    <a href="http://{{.Website}}">{{.Website}}</a> /
    Contact me:
    <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
  </div>
</footer>
