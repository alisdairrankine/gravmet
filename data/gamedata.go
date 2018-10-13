package data

type Team struct {
	Name      string     `yaml:"name"`
	Buildings []Building `yaml:"buildings"`
	Units     []Unit     `yaml:"units"`
}

type Building struct {
	Name     string   `yaml:"name"`
	Cost     int      `yaml:"cost"`
	Power    int      `yaml:"power"`
	Requires []string `yaml:"requires"`
}

type Unit struct {
	Name     string   `yaml:"name"`
	Cost     int      `yaml:"cost"`
	Requires []string `yaml:"requires"`
	Home     string   `yaml:"home"`
}
