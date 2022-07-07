package templating

import (
	"html/template"
	"errors"
	"net/http"

	"github.com/unfavorablenode/thin_node/dto"
)

var ViewDirectoryPath string = "/view/"
var ViewExtension string = "_view.html"
var HasLayout bool = true

var cache map[string]*template.Template

func CreateViewCache()	{

}

func RenderView(writer http.ResponseWriter, data dto.Dto, viewName string) error   {
    if hasCache := isCacheRegistered(); hasCache    {
	tmpl, ok := cache[viewName]
	if !ok	{
	    return errors.New("viewName not in cache")
	}
	
	tmpl.Execute(writer, data)
    } else  {
	viewPath := CreateViewPath(viewName)

	template, err := template.ParseFiles(viewPath)	

	if err != nil	{
	    return err
	}

	template.Execute(writer, data)
    }
    return nil
}

func isCacheRegistered() bool	{
    return len(cache) > 0
}

func CreateViewPath(viewName string) string	{
    return ViewDirectoryPath + viewName + ViewExtension
}
