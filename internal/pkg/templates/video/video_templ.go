// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package video

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"github.com/maxzhovtyj/live/internal/config"
	"github.com/maxzhovtyj/live/internal/models"
	"github.com/maxzhovtyj/live/internal/pkg/templates/layout"
)

func VideoRoom(c models.Context) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"\"><div class=\"bg-gray-100 flex justify-center items-center min-h-screen\"><div id=\"video-container\" class=\"flex lg:flex-none justify-center gap-4\"><div class=\"flex-1 p-2\"><video playsinline muted class=\"w-full h-full rounded\" style=\"transform: scaleX(-1);\" autoplay id=\"localClientVideo\"></video></div></div></div><div class=\"justify-between fixed bottom-0 left-0 z-50 flex w-full h-16 px-8 bg-white border-t border-gray-200\"><div class=\"items-center text-dark-500 flex\"><svg class=\"w-3 h-3 me-2\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 20 20\"><path d=\"M10 0a10 10 0 1 0 10 10A10.011 10.011 0 0 0 10 0Zm3.982 13.982a1 1 0 0 1-1.414 0l-3.274-3.274A1.012 1.012 0 0 1 9 10V6a1 1 0 0 1 2 0v3.586l2.982 2.982a1 1 0 0 1 0 1.414Z\"></path></svg> <span id=\"time-now-text\">12:43 PM</span></div><div class=\"flex items-center justify-center gap-3\"><button id=\"mute-microphone-btn\" data-tooltip-target=\"toggle-microphone-tooltip\" type=\"button\" class=\"p-2.5 group bg-gray-100 rounded-full hover:bg-gray-200 focus:outline-none focus:ring-4 focus:ring-gray-200\"><svg focusable=\"false\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" class=\"Hdh4hc cIGbvc NMm5M\"><path d=\"M12 14c1.66 0 3-1.34 3-3V5c0-1.66-1.34-3-3-3S9 3.34 9 5v6c0 1.66 1.34 3 3 3z\"></path> <path d=\"M17 11c0 2.76-2.24 5-5 5s-5-2.24-5-5H5c0 3.53 2.61 6.43 6 6.92V21h2v-3.08c3.39-.49 6-3.39 6-6.92h-2z\"></path></svg></button> <button id=\"turn-off-camera-btn\" type=\"button\" class=\"p-2.5 bg-gray-100 group rounded-full hover:bg-gray-200 focus:outline-none focus:ring-4 focus:ring-gray-200\"><svg focusable=\"false\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" class=\"Hdh4hc cIGbvc NMm5M\"><path d=\"M18 10.48V6c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2v-4.48l4 3.98v-11l-4 3.98zm-2-.79V18H4V6h12v3.69z\"></path></svg></button></div><form action=\"/\" class=\"mb-0 items-center justify-center flex\"><button type=\"submit\" class=\"inline p-2.5 group bg-gray-100 rounded-full hover:bg-gray-200 focus:outline-none focus:ring-4 focus:ring-gray-200\"><svg class=\"w-6 h-6 text-gray-800\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" fill=\"none\" viewBox=\"0 0 24 24\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5.693 16.013H7.31a1.685 1.685 0 0 0 1.685-1.684v-.645A1.684 1.684 0 0 1 10.679 12h2.647a1.686 1.686 0 0 1 1.686 1.686v.646c0 .446.178.875.494 1.19.316.317.693.495 1.14.495h1.685a1.556 1.556 0 0 0 1.597-1.016c.078-.214.107-.776.088-1.002.014-4.415-3.571-6.003-8-6.004-4.427 0-8.014 1.585-8.01 5.996-.02.227.009.79.087 1.003a1.558 1.558 0 0 0 1.6 1.02Z\"></path></svg></button></form></div></div><script src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%s/static/index.js", config.Get().Hostname))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/pkg/templates/video/video.templ`, Line: 58, Col: 70}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Main(c).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
