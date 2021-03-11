# pongo2-runner

A small utility to render [pongo2](https://github.com/flosch/pongo2)
templates. Intended to be used as a smarter replacement to `envsubst` 
to create smarter configuration files in the context
of containers.

## Example

Given the following template file saved in `./examples/test-2.cfg`
```sh
{% if env.SHELL == "/bin/zsh" %}
is_using_awesome_shell="true"
{% else %}
shell="{{ env.SHELL }}"
home="{{ env.HOME }}"
{% endif %}
```

The output of
```sh
SHELL=/bin/zsh pongo2-runner ./examples/test-2.cfg
```

is:

```sh
is_using_awesome_shell="true"
```

while the output of
```sh
SHELL=/bin/bash pongo2-runner ./examples/test-2.cfg
```

is:
```sh
shell="/bin/bash"
home="/home/username/"
```

## Compiling

### Requirements

- Go (1.16+)

### Compiling statically

```bash
git clone https://github.com/swisscom/pongo2-runner
cd pongo2-runner
CGO=0 go build -o ./pongo2-runner ./cmd
```