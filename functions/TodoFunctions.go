package functions

import "udit/api-padhai/models"

func GetAllFunctionsToCreate() []models.Function {
	var allFunctions []models.Function
	allFunctions = append(allFunctions, function_GetSessionFunction())
	return allFunctions
}

func function_GetSessionFunction() models.Function {
	return models.Function{
		FunctionName: "GetSessionId",
		FunctionCreationQuery: "create or alter function dbo.GetSessionId( 							" +
			"	@user_id            int,										" +
			"	@login_platform     varchar(7),									" +
			"	@login_dttime       datetime									" +
			") 																	" +
			"	returns varchar(50)												" +
			"begin 																" +
			"	declare @return_value   varchar(50) = null 						" +
			"	if(@login_platform not in ('ANDROID', 'IOS')) 					" +
			"		return @return_value 										" +
			"	select @return_value = '[' + cast(@user_id as varchar) + '][' + @login_platform + '][' +  replace(CONVERT(varchar, @login_dttime, 21), ' ', 'T') + ']' " +
			"	return @return_value											" +
			"end",
	}
}
