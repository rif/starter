package main

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"path/filepath"
	"thegoods.biz/tmplmgr"
)

const (
	appname   = "app"
	store_key = "foobar"
	base_path_env = "OPENSHIFT_REPO_DIR" // "APPROOT" for heroku
	ip = "OPENSHIFT_INTERNAL_IP"
)

var (
	mode          = tmplmgr.Development
	assets_dir    = filepath.Join(env(base_path_env, ""), "assets")
	template_dir  = filepath.Join(env(base_path_env, ""), "templates")
	base_template = tmplmgr.Parse(tmpl_root("base.tmpl"))
	store         = sessions.NewCookieStore([]byte(store_key))
	base_meta     = &Meta{
		CSS: list{
			"bootstrap.min.css",
			"bootstrap-responsive.min.css",
			"screen.css",
		},
		JS: list{
			"jquery.min.js",
			"bootstrap.js",
		},
		BaseTitle: appname,
	}
)

func init() {
	//set our compiler mode
	tmplmgr.CompileMode(mode)

	//add blocks to base template
	base_template.Blocks(tmpl_root("*.block"))
}

func main() {
	handle("/", handle_index)
	handle("/status/", handle_status)
	serve_static("/assets", asset_root(""))
	if err := http.ListenAndServe(env(ip, "")+":"+env("PORT", "8080"), nil); err != nil {
		log.Fatal(err)
	}
}
