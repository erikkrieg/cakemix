# Cakemix

Cakemix is a simple yet scrumptious tool for templating files using [Go templating](http://golang.org/pkg/text/template/) and [Sprig](https://github.com/Masterminds/sprig). It allows you to template file paths and contents by using a `cakemix.yaml` file that contains key-value pairs. These pairs can be interpolated in the names and contents of files and directories.

## Usage

To use Cakemix, create a directory with a `cakemix.yaml` file and any other files you want to use as templates. The `cakemix.yaml` file will contain the keys and values to be passed to the templated files. You can then use these values in the contents and names of files and directories.

### Basic Example

#### Creating a Template

Create a template directory containing a file called `cakemix.yaml` and another file called `{{.title}}.md`. The directory structure should look like this:

```
template/
  cakemix.yaml
  {{.title}}.md
```

The `cakemix.yaml` contains:

```yaml
title: cakemix
purpose: demonstrate how to use cakemix
```

The `{{.title}}.md` file has the following content:

```md
# {{ title .title }}

The purpose of this document is to: {{ .purpose }}
```

#### Using the Template

To use the template, run the following command:

```sh
cakemix ./template --output-dir .
```

This will create a file called `cakemix.md` in the current directory, with the following content:

```md
# Cakemix

The purpose of this document is to: demonstrate how to use cakemix
```

### More Advanced Usage

For more advanced usage, you can explore the examples provided in the [examples](examples) folder.

#### Command Line Options

Cakemix also offers various command line options. Here are some of them:

```
Usage:
  cakemix [template_dir] [flags]

Flags:
  -h, --help                 help for cakemix
  -i, --ignore-prompts       Do not prompt for values
  -o, --output-dir string    Write templates to this dir
  -f, --values-file string   YAML file containing prompts and default data (default "cakemix.yaml")`
```

#### Values File

The values file contains key-value pairs that are passed to the template. It's important to be familiar with YAML types to avoid unexpected behavior in the template logic. For example, some YAML values that might appear as strings could be parsed as booleans. The values file is named `cakemix.yaml` by default, but you can specify a different file using the `-f` flag.

##### Prompts

If a value in the values file is left empty (`nil`), Cakemix will prompt you to enter a value when you run it. For example, these two values will result in prompts:

```yaml
prompt_a:
prompt_b: nil
```

You can disable prompts by using the `-i` flag, like this: `cakemix ./template -i -o target`.

#### Template Functions

Cakemix provides a variety of template functions that can be used in your templates. You can find a complete list of these functions in the following documentation:

- [Go Template Functions](https://pkg.go.dev/text/template#hdr-Functions)
- [Sprig Functions](https://masterminds.github.io/sprig/)
- Cakemix Functions:
  - `padLeft`: Use `{{ padLeft .example " " 10}}` to ensure that `.example` has a minimum length of 10 characters and is padded with empty spaces.
