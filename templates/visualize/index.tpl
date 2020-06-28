{{ define "visualize/index.tpl" }}

{{ template "visualize/header.tpl" .}}

<div class="visualize-dash">

  <!-- Left panel-->
  <div class="viz-left-panel col-md-2">
    <div>
      <a href="/chart/{{ (index .chartItem 0).Name }}">Chart Detail Page</a>
    </div>
  </div>

  <!-- Main section -->
  <div class="viz-node-panel col-md-10">
    {{ template "visualize/nodegraph.tpl" .}}
  </div>

</div>

<!-- Footer section -->
{{ template "visualize/footer.tpl" .}}

{{ end }}