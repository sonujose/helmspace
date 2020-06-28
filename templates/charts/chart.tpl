{{ define "charts/chart.tpl" }}

{{ template "charts/header.tpl" .}}

<div class="main-page">

  <!-- Dashboard Banner Section -->
  <div class="chart-banner">
    <div class='container'>
      <!--Chart box-icon & name-->
      <div class="chart-box">

        <!-- Chart Main section-->
        <div class="chart-box-left">
          <div>
            {{ if (index .chartItem 0).Icon}}  
            <img class="card-img-top-detail" src={{ (index .chartItem 0).Icon }}>
            {{else}}
            <img class="card-img-top-detail" src="/static/img/default_icon.png">
            {{end}}
          </div>
          <div class="chart-metadata">
            <div class="chart-name">{{ (index .chartItem 0).Name }}</div>
            <div>Version: {{ (index .chartItem 0).Version }}</div>
            <div style="margin-top: 5px">Published on {{ (index .chartItem 0).Created.Format "Jan 02, 2006 15:04:05 UTC" }} </div>
          </div>
        </div>

        <!-- Chart bottons section-->
        <div class="chart-box-right">
          <a class="btn banner-btn" href="/visualize/{{ (index .chartItem 0).Name }}"><i class="fa fa-dashboard"></i> Visualize</a>
        </div>
      </div>

      <!--Chart Description-->
      <div class="chart-description">{{ (index .chartItem 0).Description }}</div>

    </div>
  </div>


  <!-- Detail Section-->
  <div class="container">

    <!-- The chart page tab section -->
    <div class="tabset col-md-9">
      <!--readme tab-->
      <input type="radio" name="tabset" id="tab1" aria-controls="readme" checked>
      <label for="tab1"><i class="fa fa-file-text-o tab-icon" ></i> Readme</label>
      
      <!--versions tab-->
      <input type="radio" name="tabset" id="tab4" aria-controls="versions">
      <label for="tab4"><i class="fa fa-tags tab-icon" ></i>Versions</label>

      <!--Dependency tab-->
      <input type="radio" name="tabset" id="tab5" aria-controls="dependency">
      <label for="tab5"><i class="fa fa fa-plug tab-icon"></i>Dependencies</label>
      
      <div class="tab-panels">

        <!--readme section-->
        <section id="readme" class="tab-panel">
          {{ template "charts/readme.tpl" .}}
        </section>

        <!-- versions section -->
        <section id="versions" class="tab-panel">
          {{ template "charts/versions.tpl" .}}
        </section>

        <!--Dependency section-->
        <section id="dependency" class="tab-panel">
          {{ template "charts/dependency.tpl" .}}
        </section>
      </div>
      
    </div>

    <!-- Add repo and Install chart-->
    <div class="chart-install col-md-3">
      <div>Add Repo</div>
      {{ $repoAddCommand := printf "helm repo add %s %s" .repoDetails.Name .repoDetails.URL }}
      <input type="text" class="chart-install-input" name="firstname" value ="{{ $repoAddCommand }}" readonly>
      <div>Install Chart</div>
      <input type="text" class="chart-install-input" name="firstname" value ="helm install release-1 {{.repoDetails.Name}}/{{ (index .chartItem 0).Name }}" readonly>
    </div> 
    
  </div>

</div>
<!-- Footer section -->
{{ template "charts/footer.tpl" .}}

{{ end }}