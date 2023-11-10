package {{.PkgName}}

{{- range .Imports}}
import {{.}}
{{- end }}

default allow = false

{{- if (eq $.LogicType "and") }}
allow {
    {{- range $.RuleItems }}
        {{- if (eq .Effect "deny") }}
    not {{.FnName}}
        {{- else }}
    {{.FnName}}
        {{- end}}
    {{- end }}
}

{{- range $.RuleItems }}
{{.FnName}} {
    {{- range .Subs }}
    {{.FnName}}
    {{- end }}

    {{- range .Rules }}
    {{.FnName}}
    {{- end }}
}
    {{- range .Subs }}
{{.FnName}} {
	{{.FnBody}}
}
    {{- end }}
    {{- range .Rules }}
{{.FnName}} {
	{{.FnBody}}
}
    {{- end }}
{{- end }}

{{- else }}
    {{- range $.RuleItems }}
allow {
        {{- if (eq .Effect "deny") }}
    not {{.FnName}}
        {{- else }}
    {{.FnName}}
        {{- end}}
}
    {{- end }}


{{- range $.RuleItems }}
{{- $rootFn := .FnName }}
{{- $cur_subs := .Subs }}

{{- if .Rules}}
    {{- range .Rules }}
{{$rootFn}} {
        {{- range $cur_subs }}
    {{.FnName}}
        {{- end }}
    {{.FnName}}
}
    {{- end }}
{{- else}}
    {{- range $cur_subs }}
{{$rootFn}} {
    {{.FnName}}
}
    {{- end }}
{{- end }}

    {{- range .Subs }}
{{.FnName}} {
	{{.FnBody}}
}
    {{- end }}

    {{- range .Rules }}
{{.FnName}} {
	{{.FnBody}}
}
    {{- end }}

{{- end }}

{{- end }}