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
```
$ cp-template --folder `FolderName` --number `Number of Problems` (--style "StyleName")
```

**Example**
```
$ cp-template ABC336 7
```

## To be
- style setting
