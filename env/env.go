package env

import (
    "os"
    "regexp"
)

func RegisterEnv(filePaths ...string) error    {
    filePaths = returnDefaultRouteIfNonePassed(filePaths)

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

func returnDefaultRouteIfNonePassed(filePaths []string) []string   {
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
    // If '#' rune is encountered it should skip to the next end of line
    var skipping bool = false
    
    for _, rune := range content	{
	// Check for an enter or an comment
	if rune == '\n'	{
	    skipping = false
	    matched, err := stringIsValidEnvSyntax(bufferString)

	    if err != nil   {
		return nil, err
	    }

	    if !matched	{
		bufferString = ""
		continue
	    }

	    validLines = append(validLines, bufferString)
	    bufferString = ""
	} else if rune == '#'	{
	    skipping = true
	} else if !skipping	{
	    bufferString = bufferString + string(rune)
	}
    }

    // Add what is left in the buffer it it matches the regex
    if matched , err := stringIsValidEnvSyntax(bufferString); err == nil && matched {
	validLines = append(validLines, bufferString)
    }

    return validLines, nil
}

func stringIsValidEnvSyntax(stringToTest string) (bool, error)	{
    // Match string to env syntax of "[key]=[value]"
    return regexp.Match("^\\w+=\\w+", []byte(stringToTest))
}

func parseFileContentAndRegisterVars(content []string)	{

}
