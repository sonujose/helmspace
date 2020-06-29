{{ define "charts/readme.tpl" }}
<div id="readme">

</div>

<script>

$(document).ready(function(){
  
  console.log("Ready - Fetching chart metedata from ui..");

  chartName = $("#chartName").text()
  version = $("#chartVersion").text()

  var readmeApiEndpoint = "/api/v1/readme/" + chartName + "/" + version

  console.log("API Call readme" + readmeApiEndpoint)
  
  $.get(readmeApiEndpoint, function(data, status){
    console.log("readme api call" + readmeApiEndpoint + "->" + status)
    var converter   = new showdown.Converter(),
        htmlReadme  = converter.makeHtml(data.readme);
    $("#readme").html(htmlReadme);
  });

});

</script>
{{ end }}