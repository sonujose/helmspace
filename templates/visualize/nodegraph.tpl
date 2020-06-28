{{ define "visualize/nodegraph.tpl" }}

<script>
var graph = new Springy.Graph();

var dennis = graph.newNode({
  label: 'dev-microservice',
  ondoubleclick: function() { console.log("Node clicked"); }
});
var michael = graph.newNode({label: 'Deployment'});
var jessica = graph.newNode({label: 'Service'});
var timothy = graph.newNode({label: 'Configmap'});
var barbara = graph.newNode({label: 'Secret'});
var franklin = graph.newNode({label: 'ServiceAccount'});
var monty = graph.newNode({label: 'Statefulset'});
var james = graph.newNode({label: 'Daemonset'});
var bianca = graph.newNode({label: 'Ingress'});

var dennis2 = graph.newNode({
  label: 'new-microservice',
  ondoubleclick: function() { console.log("Node clicked"); }
});
var michael2 = graph.newNode({label: 'Deployment'});
var jessica2 = graph.newNode({label: 'Service'});

graph.newEdge(dennis2, michael2, {color: '#00A0B0'});
graph.newEdge(dennis2, jessica2, {color: '#00A0B0'});

graph.newEdge(dennis, michael, {color: '#00A0B0'});
graph.newEdge(michael, dennis, {color: '#6A4A3C'});
graph.newEdge(michael, jessica, {color: '#CC333F'});
graph.newEdge(jessica, barbara, {color: '#EB6841'});
graph.newEdge(michael, timothy, {color: '#EDC951'});
graph.newEdge(franklin, monty, {color: '#7DBE3C'});
graph.newEdge(dennis, monty, {color: '#000000'});
graph.newEdge(monty, james, {color: '#00A0B0'});
graph.newEdge(barbara, timothy, {color: '#6A4A3C'});
graph.newEdge(dennis, bianca, {color: '#CC333F'});
graph.newEdge(bianca, monty, {color: '#EB6841'});

jQuery(function(){
  var springy = window.springy = jQuery('#springydemo').springy({
    graph: graph,
    nodeSelected: function(node){
      console.log('Node selected: ' + JSON.stringify(node.data));
    }
  });
});
</script>

<canvas id="springydemo" width="740" height="580" />

{{ end }}









