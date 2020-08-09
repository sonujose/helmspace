
    window.addEventListener("DOMContentLoaded", function() {

        var markupNodes = [];
        var markupEdges = [];
        var n1="n1",n2="n2",n3="n3",n4="n4",n5="n5",n6="n6",n7="n7",n8="n8"

        markupNodes.push({data: { id: n1, label: "deployment" }},
                         {data: {id: n2, label: "service" }},
                         {data: {id: n3, label: "configmap" }},
                         {data: {id: n4, label: "statefulset" }},
                         {data: {id: n5, label: "ingress" }},
                         {data: {id: n6, label: "secret" }},
                         {data: {id: n7, label: "hpa" }},
                         {data: {id: n8, label: "poddisruption" }})
        
        markupEdges.push({data: { source: n1, target: n2 }},
          {data: { source: n3, target: n1 }},
          {data: { source: n4, target: n5 }},
          {data: { source: n6, target: n1 }},
          {data: { source: n7, target: n1 }},
          {data: { source: n8, target: n1 }},
          {data: { source: n8, target: n1 }},
          {data: { source: n2, target: n5 }})

    
        var cy = (window.cy = cytoscape({
          container: document.getElementById("cy"),
      
          layout: {
            name: "dagre"
          },
          style: cytoscape.stylesheet()
          .selector('node').css({
            "content": "data(label)",
            "text-valign": "bottom",
            "background-color": "#3ea2f8",
            "border-width": 17,
            "border-color": "#4a5158",
            "border-opacity": 0.3,
            "height": "17px",
            "width":"17px",
            "font-size": "12px",
            "shape": "ellipse",
            "color": "white",
            "transition": "width 2s, height 4s"
          })
          .selector('node.highlight').css({
              "background-color": "#3ea2f8",
              "border-color": "#fdfdfd",
              "height": "20px",
              "width":"20px",
              "border-width": 20
          })
          .selector('edge.highlight').css({
            "line-color": "#f00",
          })
          .selector('edge').css({
            "width": 1,
            "line-color": "#fff",
            "target-arrow-color": "#9dbaea",
            "curve-style": "bezier"
          }),
          elements: {
            nodes: markupNodes,
            edges: markupEdges
          }
        }));
        
        cy.elements()
          .layout({
            name: "dagre",
            fit: false,
            ready: () => {
              cy.zoom(1);
              cy.center(cy.getElementById("n1"));
            }
          })
          .run();

        cy.on("tap", "node", function tapNode(e) {
          const node = e.target;
          let msg = node.data
          console.log(msg)
        //   node
        //     .connectedEdges()
        //     .targets()
        //     .style("visibility", "hidden");
        });
        
        cy.on("zoom", e => {
          console.log(cy.zoom());
        });

      cy.on('mouseover', 'node', function mouseover(e){
          var sel = e.target;
          sel.addClass('highlight');
          //cy.elements().difference(sel.outgoers()).not(sel).addClass('semitransp');
          //sel.addClass('highlight').outgoers().union(sel.incomers()).addClass('highlight');
          sel.connectedEdges().style({ 'line-color': '#3ea2f8' });
      });
      cy.on('mouseout', 'node', function mouseout(e){
          var sel = e.target;
          //cy.elements().removeClass('semitransp');
          sel.removeClass('highlight').outgoers().union(sel.incomers()).removeClass('highlight');
          sel.connectedEdges().style({ 'line-color': '#fff' });
      });

      });