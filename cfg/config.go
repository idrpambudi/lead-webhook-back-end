package cfg

type (
	Configurations struct {
		Server   Server
		Database Database
	}

	Server struct {
		Port int
	}

	Database struct {
		MySql MySql
	}

	MySql struct {
		URL string
	}
)
