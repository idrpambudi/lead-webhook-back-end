package cfg

type (
	Configurations struct {
		Port     int
		Database Database
	}

	Database struct {
		MySql MySql
	}

	MySql struct {
		URL string
	}
)
