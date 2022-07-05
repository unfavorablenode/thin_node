package env

import (
    "os"
    "regexp"
)

func RegisterEnv(filePaths ...string) error    {
    filePaths = validateAmountOfPathsReturnDefaultIfNone(filePaths)

    for _, path := range filePaths  {
	if err := checkIfFileExists(path); err != nil	{
		return err
	}

	fileContent, err := getFileContent(path)
	if err != nil	{
	    return err
	}

	validLines, err := retrieveValidLinesFromContent(fileContent)

	parseFileContentAndRegisterVars(validLines)
    }

    return nil
}

func validateAmountOfPathsReturnDefaultIfNone(filePaths []string) []string   {
    if len(filePaths) <= 0  {
	return []string { ".env" }
    }
    return filePaths
}

func checkIfFileExists(filePath string) error	{
    // Attempt to get file description. Returns error if it cannot get description.
    if _, err := os.Stat(filePath); err != nil  {
	return err
    }

    return nil
}

func getFileContent(filePath string) (string, error)	{
    data, err := os.ReadFile(filePath)
    // Convert byte array to string
    return string(data), err
}

func retrieveValidLinesFromContent(content string) ([]string, error)	{
    validLines := []string{}
    var bufferString string = ""
    
    for _, rune := range content	{
	// Check for an enter or an comment
	if rune == '\n' || rune == '#'	{
	    // Match bufferString to env syntax of "[key]=[value]"
	    matched, err := regexp.Match("/^\\w+=\\w+/", []byte(bufferString)) 

	    if matched == false	{
		continue
	    }

	    if err != nil	{
		return nil, err
	    }

	    validLines = append(validLines, bufferString)
	    bufferString = ""
	} else	{
	    bufferString = bufferString + string(rune)
	}
    }

    return validLines, nil
}

func parseFileContentAndRegisterVars(content []string)	{

}
