# Configuration

## general settings

* threads
* debug
* libpaths
* rpath
* output
* loglevel

## whitelist/blacklist

A `whitelist` or `blacklist` can be maintained for the dependency check
to only check a subset of packages in a folder of tarballs. Only one can be present
at a given time. These can be specified on the `pkgcheck.toml` 
or directly through the CLI via the package name. 

```toml
whitelist = [
  "dplyr",
  "ggplot2"
]
```

multiple packages can be specified on the CLI via multiple calls:

```bash
--whitelist=dplyr --whitelist=ggplot2
```

The CLI also supports a `whiteliststr` argument that allows a comma separated
list to be specified instead.

```bash
whiteliststr="dplyr,ggplot2"
```
