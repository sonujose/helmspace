<!--Embed the header.html template at this location-->
{{ template "header.tpl" .}}

<!-- Dashboard Card Tiles -->
<div class='container chart-detail'>

  <!-- Add repo and Install chart-->
  <div class="chart-install">
    <h3>Add Repo</h3>
    {{ $repoAddCommand := printf "helm repo add %s %s" .repoDetails.Name .repoDetails.URL }}
    <input type="text" class="chart-install-input" name="firstname" value ="{{ $repoAddCommand }}" readonly>
    <h3>Install Chart</h3>
    <input type="text" class="chart-install-input" name="firstname" value ="helm install release-1 {{.repoDetails.Name}}/{{ (index .chartItem 0).Name }}" readonly>
  </div>

  <!--Chart box-icon & name-->
  <div class="chart-box">
    <div class="chart-icon-container">
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

  <!--Chart Description-->
  <div class="chart-description">{{ (index .chartItem 0).Description }}</div>

  <!-- The chart page tab section -->
  <div class="tabset">
    <!--readme tab-->
    <input type="radio" name="tabset" id="tab1" aria-controls="keriko" checked>
    <label for="tab1">Readme</label>
    
    <!--Visualize tab-->
    <input type="radio" name="tabset" id="tab2" aria-controls="maxoni">
    <label for="tab2">Visualize</label>

    <!--Info tab-->
    <input type="radio" name="tabset" id="tab3" aria-controls="larboni">
    <label for="tab3">Info</label>
    
    <!--versions tab-->
    <input type="radio" name="tabset" id="tab4" aria-controls="bingo">
    <label for="tab4">Versions</label>

     <!--versions tab-->
    <input type="radio" name="tabset" id="tab5" aria-controls="plekora">
    <label for="tab5">Dependencies</label>
    
    <div class="tab-panels">
      <!--readme section-->
      <section id="keriko" class="tab-panel">
        <div>
        Introduction
        This chart bootstraps an Apache Airflow deployment on a Kubernetes cluster using the Helm package manager.

        Bitnami charts can be used with Kubeapps for deployment and management of Helm Charts in clusters. 
        This Helm chart has been tested on top of Bitnami Kubernetes Production Runtime (BKPR). 
        Deploy BKPR to get automated TLS certificates, logging and monitoring for your applications.
        </div>
      </section>

      <!--visualize section-->
      <section id="maxoni" class="tab-panel">
        <div>
          Feature Coming Soon!!
        </div>
      </section>

      <!--info section-->
      <section id="larboni" class="tab-panel">
        <div>Version: <span>{{ (index .chartItem 0).Version }}</span></div>
        <div>AppVersion: <span>{{ (index .chartItem 0).AppVersion }}</span></div>
        {{ if (index .chartItem 0).Home }}
        <div>Home: <a target="_blank" href="{{ (index .chartItem 0).Home }}">{{ (index .chartItem 0).Home }}</a></div>
        {{end}}
      </section>

      <!-- versions section -->
      <section id="bingo" class="tab-panel">
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
      </section>

      <!--Dependency section-->
      <section id="plekora" class="tab-panel">
        <div>
          Primary Dependency of the charts
          <ul>
            <li>Postgres</li>
          </ul>
        </div>
      </section>

    </div>
  </div>
  
</div>

<!-- Footer section -->
{{ template "footer.tpl" .}}