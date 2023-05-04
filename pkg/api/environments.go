package api

type Environment string

func (e Environment) String() string {
	return string(e)
}

const (
	CustomEnvironment     Environment = "custom"
	DevelopEnvironment    Environment = "dev"
	ProductionEnvironment Environment = ""
	StagingEnvironment    Environment = "stg"
)
