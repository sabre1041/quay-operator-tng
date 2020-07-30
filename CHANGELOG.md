# Changelog

### v0.0.1

Initial release. Operator is functional and will deploy/upgrade Quay and all managed components.

Bugfixes:
* [#10][]: Add name prefix to created objects to prevent name collisions. 
* [#17][]: Fix create/update logic when reconciling `QuayRegistry` multiple times.
* [#22][]: Stop checking upgrade deployment status once upgrade has completed.

Enhancements:
* [#1][]: Greenfield.
* [#2][]: Initial Kustomize definitions.
* [#3][]: Scaffold project using Kubebuilder.
* [#4][]: `QuayRegistry` API definition.
* [#5][]: Implement `Inflate()` function to "inflate" a `QuayRegistry` into k8s resources using Kustomize.
* [#6][]: Implement reconcile loop for controller.
* [#8][]: Add CI to project using GitHub workflows.
* [#11][]: Add single namespace mode.
* [#12][]: Auto-generate secret keys.
* [#13][]: Generate component secrets in Go versus Kustomize.
* [#14][]: Improve test reliability by using `resourceVersion` of pods.
* [#15][]: Add initial `ClusterServiceVersion` for OLM installation.
* [#16][]: Switch to in-memory filesystem for Kustomize.
* [#19][]: Add `spec.components` API.
* [#20][]: Add `spec.desiredVersion` API.
* [#21][]: Add `status.currentVersion` API.
