package courseimport

type course struct {
	ID          string `yaml:"id"`
	Category    string `yaml:"category"`
	Difficulty  string `yaml:"difficulty"`
	Description string `yaml:"description"`
	Publish     bool   `yaml:"publish"`
	Slug        string `yaml:"slug"`
	Title       string `yaml:"title"`

	CoruseItems []*courseItem `yaml:"course_items"`
}

type courseItem struct {
	ID          string `yaml:"id"`
	ChallengeID string `yaml:"challenge_id"`
}
