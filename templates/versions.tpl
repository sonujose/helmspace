<table class="table table-dark">
<thead>
    <tr>
    <th scope="col">Version</th>
    <th scope="col">AppVersion</th>
    <th scope="col">Created On</th>
    </tr>
</thead>
<tbody>
{{range $v:= .chartItem}}
    <tr>
    <td>{{ $v.Version }}</td>
    <td>{{ $v.AppVersion }}</td>
    <td>{{ $v.Created.Format "Jan 02, 2006 15:04:05 UTC" }}</td>
    <td><a href="{{$.repoDetails.URL}}/{{ (index $v.Urls 0) }}">Download</a></td>
    </tr>
    {{end}}
</tbody>
</table>