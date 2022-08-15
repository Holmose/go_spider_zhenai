package config

const (
	// ParseCity ... Parser Names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// ItemSaverPort Service Ports
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	// ElasticIndex ElasticSearch
	ElasticIndex = "dating_profile"

	// ItemSaverRpc RPC Endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
)
