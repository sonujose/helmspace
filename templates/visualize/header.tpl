{{ define "visualize/header.tpl" }}
<!doctype html>
<html>

  <head>
    <!--Use the title variable to set the title of the page-->
    <title>Helmspace</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta charset="UTF-8">

    <!--Site Manifest and Icons-->
    <link rel="apple-touch-icon" sizes="180x180" href="/static/img/icons/apple-touch-icon.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/static/img/icons/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/static/img/icons/favicon-16x16.png">
    <link rel="manifest" href="/static/img/icons/site.webmanifest">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
    
    <!--Use bootstrap to make the application look nice-->
    <link rel="stylesheet" href="/static/css/visualize.css">
    <link rel="stylesheet" href="/static/css/tiles.css">
    <link rel="stylesheet" href="/static/css/nodegraph.css">
    
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/cytoscape/3.2.7/cytoscape.js"></script>
    <script src="https://unpkg.com/dagre@0.7.4/dist/dagre.js"></script>
    <script src="https://cytoscape.org/cytoscape.js-dagre/cytoscape-dagre.js"></script>
    <script src="/static/js/nodegraph.js"></script>
  </head>

  <body>
    <!--Embed the menu.html template at this location-->
    {{ template "visualize/menu.tpl" . }}

{{ end }}