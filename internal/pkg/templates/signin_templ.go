// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/maxzhovtyj/live/internal/pkg/templates/layout"

func SignIn() templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"min-h-screen bg-gray-100 flex flex-col justify-center py-12 sm:px-6 lg:px-8\"><div class=\"sm:mx-auto sm:w-full sm:max-w-md\"><h2 class=\"mt-6 text-center text-3xl font-extrabold text-gray-900\">Sign in to your account</h2><p class=\"mt-2 text-center text-sm text-gray-600 max-w\">Or <a href=\"/sign-up\" class=\"font-medium text-primary\">create an account</a></p></div><div class=\"mt-8 sm:mx-auto sm:w-full sm:max-w-md\"><div class=\"bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10\"><form hx-ext=\"response-targets\" hx-post=\"/sign-in\" hx-target-*=\"#serious-errors\" class=\"space-y-6\"><div><label for=\"email\" class=\"block text-sm font-medium text-gray-700\">Email address</label><div class=\"mt-1\"><input id=\"email\" name=\"email\" type=\"email\" autocomplete=\"email\" required class=\"appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm\" placeholder=\"Enter your email address\"></div></div><div><label for=\"password\" class=\"block text-sm font-medium text-gray-700\">Password</label><div class=\"mt-1\"><input id=\"password\" name=\"password\" type=\"password\" autocomplete=\"current-password\" required class=\"appearance-none rounded-md relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm\" placeholder=\"Enter your password\"></div></div><div id=\"serious-errors\" class=\"font-medium text-red-700\"></div><div><button type=\"submit\" class=\"relative w-full btn btn-primary\">Sign in</button></div></form></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Index().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
