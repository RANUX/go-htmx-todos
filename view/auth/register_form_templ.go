// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package auth

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func RegisterForm() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"space-y-6\" action=\"#\"><div><label for=\"username\" class=\"block mb-2 text-sm font-medium text-gray-900\">Username:</label> <input type=\"text\" id=\"username\" name=\"username\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5\" placeholder=\"Your name\" required></div><div><label for=\"password\" class=\"block mb-2 text-sm font-medium text-gray-900\">Password:</label><div class=\"relative\"><input type=\"password\" id=\"password\" name=\"password\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 pr-10\" placeholder=\"**********\" required><div class=\"absolute inset-y-0 right-0 pr-3 flex items-center text-sm leading-5\"><img src=\"https://cdn.jsdelivr.net/npm/bootstrap-icons/icons/eye-fill.svg\" alt=\"Show password\" class=\"cursor-pointer\"></div></div></div><div class=\"flex justify-between\"><div class=\"text-sm\"><a href=\"#\" class=\"text-blue-600 hover:text-blue-500\">Forgot password?</a></div><button type=\"submit\" class=\"text-white bg-green-600 hover:bg-green-700 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center\">Sign up</button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
