package constants

const (
	LogInfoGetAll    = "Get list of %s with Offset %d and Limit %d."
	LogInfoTotalRows = "Get total rows %s"
	LogInfoSummary   = "Get summary of %s"

	LogErrGetAll                  = "failed to get list of %s"
	LogErrGetByID                 = "failed to get %s by ID %d"
	LogErrGetRows                 = "failed to get total rows of %s"
	LogErrGetSummary              = "failed to get summary of %s"
	LogErrGenerateQueryWithFilter = "failed generate query with filter"
)
