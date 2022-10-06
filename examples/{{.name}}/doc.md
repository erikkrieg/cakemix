# {{ .nested.foo }}

This is a message from values: "{{ .str_quoted }}"

Here is list of points:
{{- range .list }}
{{ . }}. Point {{ . }}
{{- end -}}

