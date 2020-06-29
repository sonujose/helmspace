{{ define "charts/readme.tpl" }}
<div id="readme">

</div>

<script>

$(document).ready(function(){
  
  console.log("Ready - Fetching readme content...");
  var readmeApiEndpoint = "/api/v1/readme/core-microservice/1.1.0" 

  $.get(readmeApiEndpoint, function(data, status){
    console.log("readme api call" + readmeApiEndpoint + "->" + status)
    var converter   = new showdown.Converter(),
        htmlReadme  = converter.makeHtml(data.readme);
    $("#readme").html(htmlReadme);
  });

});

</script>
{{ end }}