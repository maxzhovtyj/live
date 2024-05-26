// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"github.com/maxzhovtyj/live/internal/models"
	"github.com/maxzhovtyj/live/internal/pkg/templates/layout"
)

func Chat(c models.Context, chatID int32) templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"h-full mx-auto flex max-w-7xl justify-between p-6 lg:px-8 gap-2 flex-wrap\"><div class=\"flex-1 rounded-2xl py-4 lg:p-2 bg-gray-100 flex flex-col gap-2 align-items-center\"><div hx-trigger=\"load\" hx-get=\"/conversations\"></div><button hx-get=\"/modal\" hx-target=\"#modals-here\" hx-trigger=\"click\" data-bs-toggle=\"modal\" data-bs-target=\"#modals-here\" class=\"btn btn-primary\">New Chat</button><div id=\"modals-here\" class=\"modal modal-blur fade\" style=\"display: none\" aria-hidden=\"false\" tabindex=\"-1\"><div class=\"modal-dialog modal-lg modal-dialog-centered\" role=\"document\"><div class=\"modal-content\"></div></div></div></div><div class=\"justify-between flex flex-grow flex-col rounded-2xl bg-gray-100 h-full p-4 lg:p-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if chatID == -1 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 class=\"text-center m-4 text-4xl break-words\">Please select chat</h1>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"chat-messages\" class=\"overflow-y-auto h-full\"></div><form hx-on:submit=\"this.reset()\" hx-ext=\"ws\" ws-connect=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(
					"/ws/chat?id=%d", chatID))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/pkg/templates/chat.templ`, Line: 38, Col: 33}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" ws-send id=\"form\" class=\"mt-2 bg-white w-full pl-3 pr-1 py-1 rounded-3xl border border-gray-200 items-center gap-2 inline-flex\n        justify-between\"><div class=\"flex items-center gap-2 grow\"><input name=\"chat-message\" class=\"w-full grow shrink basis-0 text-black text-xs font-medium leading-4 focus:outline-none\" placeholder=\"Type here...\"></div><div class=\"flex items-center gap-2\"><button type=\"submit\" class=\"btn btn-primary rounded-3xl\"><span class=\"text-white text-xs font-semibold leading-4 px-2\">Send</span></button></div></form>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><script>\n    document.addEventListener(\"htmx:wsAfterMessage\", e => {\n        const messagesDiv = document.getElementById(\"chat-messages\");\n        messagesDiv.scrollTop = messagesDiv.scrollHeight;\n    })\n</script>")
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
