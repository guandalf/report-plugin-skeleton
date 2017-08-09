# report-plugin-skeleton

This is the seed project to author a new [reporting plugin](https://docs.getgauge.io/plugins.html#reporting-plugins) for [Gauge](https://getgauge.io).

Note: If you are looking to get a non-binary text format of the execution result, please see https://github.com/apoorvam/json-report

## Getting Started

TLDR: Follow the `TODO`s in the code. 

- Change `plugin.json`
  - Set the placeholder values with plugin specific information
  - `gaugeVersionSupport` field indicates the range of `gauge` version that the plugin works with. This is usually driven by the `gauge-proto` API. If there are breaking changes, the plugin should upgrade itself to be compatible with `gauge`. `gaugeVersionSupport` enables plugin authors to handle this at their pace, by enabling a lock down of gauge version.
- Change Imports
  - all the internal package dependency uses this skeleton repository's namespace. You should change this to be the repository of your plugin.
- Properties/flags - if you need to setup flags/configurations' default values, you can do so by adding to `env.AddDefaultPropertiesToProject()`
- Implement your logic
  - At this point, your plugin is ready to receive messages from Gauge
  - At the end of execution, your plugin will receive a `gauge_messages.SuiteExecutionResult` instance, that contains meta as well as detailed information.
  - Navigate through the values and choose relevant fields. `ProtoSuiteResult` message defined in `spec.proto` should be a good place to start to understand the schema.
## Build

### Requirements
* [Golang](http://golang.org/)

### Compiling
Download dependencies
```
go get -t ./...
```
Compilation
```
go run build/make.go
```

For cross-platform compilation

```
go run build/make.go --all-platforms
```


## Install
After compilation

```
go run build/make.go --install
```

### Installing to a CUSTOM_LOCATION

```
go run build/make.go --install --plugin-prefix CUSTOM_LOCATION
```

## Run

Note: Run after install.

- Create a new Gauge project

```
gauge init java
```

- Add this plugin to project

```
gauge install <plugin name>
```

- Execute gauge specs

```
gauge run specs
```

- Inspect the reports directory

```
tree ./reports # only on bash with tree command installed
```

## Package

Note: Run after build

```
go run build/make.go --distro
```

For distributable across platforms: Windows and Linux for both x86 and x86_64

```
go run build/make.go --distro --all-platforms
```

### Installing a package

```
gauge --install <plugin name> --file <plugin artifact zip>
```

## Deploy

New distribution details need to be updated in the `<plugin name>-install.json` file in the [gauge plugin repository](https://github.com/getgauge/gauge-repository). 

Note: This is required if you wish to install the plugin using `gauge install` command. 

Alternatively, if you wish to privately host a plugin repository, you may change the [`GAUGE_REPOSITORY_URL`](https://docs.getgauge.io/configuration.html#gauge-properties-2) in `gauge.properties`


