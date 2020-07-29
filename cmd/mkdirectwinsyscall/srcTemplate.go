package main

const srcTemplate = `
{{define "main"}}// Code generated by 'go generate'; DO NOT EDIT.

package {{packagename}}
import (
	{{range .StdLibImports}}"{{.}}"
	{{end}}
	{{range .ExternalImports}}"{{.}}"
	{{end}}
	{{.BananaImport}}
)

var _ unsafe.Pointer

{{.VarBlock}}

{{range .Funcs}}{{if .HasStringParam}}{{template "helperbody" .}}{{end}}{{template "funcbody" .}}{{end}}
{{end}}

{{define "funcbody"}}
func {{.HelperName}}({{.HelperParamList}}) {{template "results" .}}{
	{{.BananaLoader}}
{{template "tmpvars" .}}	{{template "syscall" .}}	{{template "tmpvarsreadback" .}}
{{template "seterror" .}}{{template "printtrace" .}}	return
}
{{end}}


{{define "results"}}{{if .Rets.List}}{{.Rets.List}} {{end}}{{end}}

{{define "bananaloader"}} {{end}}

{{define "tmpvars"}}{{range .Params}}{{if .TmpVarCode}}	{{.TmpVarCode}}
{{end}}{{end}}{{end}}

{{define "syscall"}}{{.Rets.SetReturnValuesCode}}{{.BananaphoneSyscall}}(sysid, {{.SyscallParamList}}){{end}}

{{define "tmpvarsreadback"}}{{range .Params}}{{if .TmpVarReadbackCode}}
{{.TmpVarReadbackCode}}{{end}}{{end}}{{end}}

{{define "seterror"}}{{if .Rets.SetErrorCode}}	{{.Rets.SetErrorCode}}
{{end}}{{end}}

{{define "printtrace"}}{{if .PrintTrace}}	print("SYSCALL: {{.Name}}(", {{.ParamPrintList}}") (", {{.Rets.PrintList}}")\n")
{{end}}{{end}}

`