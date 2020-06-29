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
    
    <!--Scripts JQuery and bootstrap-->
    <script src="https://code.jquery.com/jquery-2.2.4.min.js" integrity="sha256-BbhdlvQf/xTY9gja0Dq3HiwQF8LaCRTXxZKRutelT44=" crossorigin="anonymous"></script>
    <script async src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
    
    <!-- Visualize library scripts -->
    <script src="/static/js/springy.min.js"></script>
    <script src="/static/js/springyui.min.js"></script>
  </head>

  <body>
    <!--Embed the menu.html template at this location-->
    {{ template "visualize/menu.tpl" . }}

{{ end }}