// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Header() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header class=\"bg-white relative\"><nav class=\"hidden mx-auto lg:flex lg:flex-1 max-w-7xl lg-items-center lg-justify-between p-6 lg:px-8\" aria-label=\"Global\"><div class=\"flex -m-1.5 p-1.5 items-center justify-between\"><a href=\"/\" class=\"px-8 -m-1.5 p-1.5\"><div class=\"text-4xl font-bold tracking-wide text-gray-900\">LIVE</div></a> <a href=\"/chat\" class=\"px-8 text-sm font-semibold leading-6 text-gray-900\">Chats</a> <a href=\"/meeting\" class=\"px-8 text-sm font-semibold leading-6 text-gray-900\">Meetings</a></div><div class=\"hidden lg:flex lg:flex-1 lg:justify-end\"><a hx-post=\"/sign-out\" type=\"button\" class=\"flex items-center -mx-3 cursor-pointer block rounded-lg px-3 py-2.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Log out <svg class=\"ml-2 w-6 h-6 text-gray-800 dark:text-white\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 16 16\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 8h11m0 0-4-4m4 4-4 4m-5 3H3a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h3\"></path></svg></a></div></nav><div class=\"lg:hidden flex items-end absolute top-0 right-0 m-8\"><button class=\"relative group\" hx-on:click=\"document.getElementById(&#39;nav-content&#39;).classList.toggle(&#39;hidden&#39;);\"><div class=\"relative flex overflow-hidden items-center justify-center rounded-full w-[50px] h-[50px] transform transition-all bg-slate-700 ring-0 ring-gray-300 hover:ring-8 group-focus:ring-4 ring-opacity-30 duration-200 shadow-md\"><div class=\"flex flex-col justify-between w-[20px] h-[20px] transform transition-all duration-300 origin-center overflow-hidden group-focus:translate-x-1.5\"><div class=\"bg-white h-[2px] w-7 transform transition-all duration-300 origin-left group-focus:rotate-[42deg] group-focus:w-2/3 delay-150\"></div><div class=\"bg-white h-[2px] w-7 rounded transform transition-all duration-300 group-focus:translate-x-10\"></div><div class=\"bg-white h-[2px] w-7 transform transition-all duration-300 origin-left group-focus:-rotate-[42deg] group-focus:w-2/3 delay-150\"></div></div></div></button></div><!-- Mobile menu, show/hide based on menu open state. --><div id=\"nav-content\" class=\"hidden\" role=\"dialog\" aria-modal=\"true\"><!-- Background backdrop, show/hide based on slide-over state. --><div class=\"fixed inset-0 z-10\"></div><div class=\"fixed inset-y-0 right-0 z-10 w-full overflow-y-auto bg-white px-6 py-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10\"><div class=\"flex items-center justify-between\"><button type=\"button\" class=\"-m-2.5 rounded-md p-2.5 text-gray-700\" hx-on:click=\"document.getElementById(&#39;nav-content&#39;).classList.toggle(&#39;hidden&#39;);\"><span class=\"sr-only\">Close menu</span> <svg class=\"h-6 w-6\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" aria-hidden=\"true\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M6 18L18 6M6 6l12 12\"></path></svg></button></div><div class=\"mt-6 flow-root\"><div class=\"-my-6 divide-y divide-gray-500/10\"><div class=\"space-y-2 py-6\"><a href=\"/chat\" class=\"-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Chat</a> <a href=\"/meeting\" class=\"-mx-3 block rounded-lg px-3 py-2 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Meetings</a></div><div class=\"py-6\"><a hx-post=\"/sign-out\" type=\"button\" class=\"cursor-pointer -mx-3 block rounded-lg px-3 py-2.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50\">Log out</a></div></div></div></div></div></header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
