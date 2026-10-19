package main
import("os";"github.com/oneworkspacestudio/one-workspace/pkg/apphost";"github.com/oneworkspacestudio/my-search/internal/app";"github.com/oneworkspacestudio/my-search/internal/webui")
var version="0.1.0-preview.1"
func main(){if e:=apphost.Main(apphost.Config{ID:"search",Name:"My Search",Version:version,Listen:"127.0.0.1:9011"},webui.Files(),func(h *apphost.Host)(apphost.Application,error){return app.New(h)});e!=nil{apphost.Report(e);os.Exit(1)}}
