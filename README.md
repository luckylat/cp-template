# CP-Template

## What is this


## How to Install
TBD

## How to Use

### 1. Set environment
Set environment file `$XDG_CONFIG_HOME/cp-template/settings.toml` or `$HOME/.config/cp-template/settings.toml`

```toml
[templates]
    default = "filepathA"        # required
    somestyle = "filepathB"      # option
    somestyle2 = "filepathC"     
```

### 2. Execute Command
```sh
$ cp-template --folder `FolderName` --number `Number of Problems` (--style `StyleName`)
```

#### Example
```sh
$ cp-template --folder ABC336 --number 7
```

##### Example /w StyleName

```toml
[templates]
    default = "A filepath"
    query = "Other filepath"
```

```sh
$ cp-template --folder CF1928 --number 6 --style query
```

## To be
- style setting
