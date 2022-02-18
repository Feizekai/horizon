package grafana

type Mapper map[string]*Grafana

type Grafana struct {
	BasicDashboard      string `yaml:"basicDashboard"`
	ServerlessDashboard string `yaml:"serverlessDashboard"`
	MemcachedDashboard  string `yaml:"memcachedDashboard"`
	QuerySeries         string `yaml:"querySeries"`
}

type SLO struct {
	OverviewDashboard string `yaml:"overviewDashboard"`
	APIDashboard      string `yaml:"apiDashboard"`
	PipelineDashboard string `yaml:"pipelineDashboard"`
}
