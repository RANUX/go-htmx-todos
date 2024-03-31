// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"todo/data"
	"todo/types"
	"todo/view/alert"
)

func GetUser(ctx context.Context) *data.User {
	if user, ok := ctx.Value(string(types.ContextEnum.User)).(*data.User); ok {
		return user
	}
	return &data.User{Username: "anonymous"}
}

func Base() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Todos Application</title><script src=\"/public/js/htmx.min.js\" defer></script><script src=\"/public/js/tailwind.min.js\" defer></script><link rel=\"stylesheet\" href=\"/public/css/app.css\"><script src=\"/public/js/app.js\"></script></head><body><div id=\"load\" class=\"load\"></div><nav class=\"bg-white border-b border-gray-200 z-30 w-full\"><div class=\"px-6 py-3\"><div class=\"flex items-center justify-between\"><div class=\"flex items-center justify-between w-full\"><div class=\"text-xl font-semibold text-gray-700\">Todos App ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(GetUser(ctx).Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/layout/base.templ`, Line: 38, Col: 88}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"hidden md:flex items-center ml-10\"><a href=\"/todos\" class=\"text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium\">Todos</a> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if GetUser(ctx).Username != "anonymous" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"/profile\" class=\"text-gray-700 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium\">Profile</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></div></nav><div id=\"htmx-error-alert\" hidden>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = alert.Alert("Oops, something went wrong with the server...", types.AlertEnum.Error).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div id=\"htmx-error-alert-unknown\" hidden>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = alert.Alert("Unexpected error, check your connection and refresh the page", types.AlertEnum.Error).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div id=\"htmx-alerts\" hx-get=\"/alerts\" hx-trigger=\"htmx:load from:body\" hx-swap=\"beforeend\"></div><main id=\"main\" class=\"pt-24 pb-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main></body><script>\n\t\t\tdocument.body.addEventListener('htmx:afterRequest', function(event) {\n\t\t\t\tconsole.log('htmx:afterRequest');\n\n\t\t\t\tvar htmxAlert = document.querySelector('#htmx-error-alert');\n\t\t\t\tvar htmxAlertUnknown = document.querySelector('#htmx-error-alert-unknown');\n\n\t\t\t\tif (event.detail.successful) {\n\t\t\t\t\t// If the event was successful and the #htmx-error-alert element exists\n\t\t\t\t\tif (htmxAlert) {\n\t\t\t\t\t\thtmxAlert.hidden = true;\n\t\t\t\t\t}\n\t\t\t\t} else if (event.detail.failed && event.detail.xhr) {\n\t\t\t\t\t// If the request failed and there is an xhr object with the error details\n\t\t\t\t\tconsole.error(`server error: ${event.detail.xhr.status} - ${event.detail.xhr.statusText}`);\n\t\t\t\t\tif (htmxAlert) {\n\t\t\t\t\thtmxAlert.hidden = false;\n\t\t\t\t\t}\n\t\t\t\t} else {\n\t\t\t\t\t// For all other cases, assume an unknown error\n\t\t\t\t\tconsole.error('htmx:afterRequest unknown error');\n\t\t\t\t\tif (htmxAlertUnknown) {\n\t\t\t\t\t\thtmxAlertUnknown.hidden = false;\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t});\n\n\t\t\tdocument.addEventListener('htmx:afterSwap', function(event) {\n\t\t\t\tif(event.target.id === 'htmx-alerts') {\n\t\t\t\t\t// remove alerts after 5 seconds\n\t\t\t\t\tsetTimeout(function() {\n\t\t\t\t\t\tdocument.getElementById('htmx-alerts').innerHTML = '';\n\t\t\t\t\t}, 5000);\n\t\t\t\t}\n\t\t\t});\n\n\t\t</script></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
