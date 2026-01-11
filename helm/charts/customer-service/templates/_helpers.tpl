{{/*
Expand the name of the chart.
*/}}
<<<<<<<< HEAD:helm/charts/communication-service/templates/_helpers.tpl
{{- define "communication-service.name" -}}
========
{{- define "customer-service.name" -}}
>>>>>>>> dev:helm/charts/customer-service/templates/_helpers.tpl
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
<<<<<<<< HEAD:helm/charts/communication-service/templates/_helpers.tpl
{{- define "communication-service.fullname" -}}
========
{{- define "customer-service.fullname" -}}
>>>>>>>> dev:helm/charts/customer-service/templates/_helpers.tpl
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
<<<<<<<< HEAD:helm/charts/communication-service/templates/_helpers.tpl
{{- define "communication-service.chart" -}}
========
{{- define "customer-service.chart" -}}
>>>>>>>> dev:helm/charts/customer-service/templates/_helpers.tpl
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
<<<<<<<< HEAD:helm/charts/communication-service/templates/_helpers.tpl
{{- define "communication-service.labels" -}}
helm.sh/chart: {{ include "communication-service.chart" . }}
{{ include "communication-service.selectorLabels" . }}
========
{{- define "customer-service.labels" -}}
helm.sh/chart: {{ include "customer-service.chart" . }}
{{ include "customer-service.selectorLabels" . }}
>>>>>>>> dev:helm/charts/customer-service/templates/_helpers.tpl
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
<<<<<<<< HEAD:helm/charts/communication-service/templates/_helpers.tpl
{{- define "communication-service.selectorLabels" -}}
app.kubernetes.io/name: {{ include "communication-service.name" . }}
========
{{- define "customer-service.selectorLabels" -}}
app.kubernetes.io/name: {{ include "customer-service.name" . }}
>>>>>>>> dev:helm/charts/customer-service/templates/_helpers.tpl
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
<<<<<<<< HEAD:helm/charts/communication-service/templates/_helpers.tpl
{{- define "communication-service.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "communication-service.fullname" .) .Values.serviceAccount.name }}
========
{{- define "customer-service.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "customer-service.fullname" .) .Values.serviceAccount.name }}
>>>>>>>> dev:helm/charts/customer-service/templates/_helpers.tpl
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
