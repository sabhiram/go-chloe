# chloe
[![Build Status](https://travis-ci.org/sabhiram/chloe.svg?branch=master)](https://travis-ci.org/sabhiram/chloe) [![Coverage Status](https://coveralls.io/repos/sabhiram/chloe/badge.png)](https://coveralls.io/r/sabhiram/chloe)

Chloe is a command line utility written in Go to simplify deletion of un-needed files. 

Files to be deleted are inferred from a json file (`--input`) which is expected to contain an array with key `chloe`. The lines in this array will be interpreted the same way a `.gitignore` file is parsed. 

The rules for parsing the files specified under the `chloe` key obey the rules found at the [gitignore docs](http://git-scm.com/docs/gitignore). The parsing rules and interface methods can be found at the [go-git-ignore](https://github.com/sabhiram/go-git-ignore) library.

## Sample json file

*test.json*
```json
{
    "other_key": "some value",

    "chloe": [
        "**/bower/**/*.md",
        "**/*.out"
    ]
}
```

The above file will instruct `chloe` to tag any files with a `.md` extension two folders deep from a folder called "bower". It also tags all `.out` files within a folder (relative to the root dir where `chloe` is run)

## Sample Usage:

Assuming the above `test.json` file (this is a make-believe environment...)
```sh
$chloe list -i test.json
Found 5 extra files:
 - app/public/bower/angular/README.md
 - app/public/bower/angular-route/README.md
 - app/public/bower/bootstrap/README.md
 - app/public/bower/jquery-ui/README.md
 - app/public/bower/underscore/README.md

$chloe dispatch -i test.json
Found 5 extra files:
 - app/public/bower/angular/README.md
 - app/public/bower/angular-route/README.md
 - app/public/bower/bootstrap/README.md
 - app/public/bower/jquery-ui/README.md
 - app/public/bower/underscore/README.md
Purge 5 files? [ Yes | No ]: 
y
Deleted 5 files!

$chloe list -i test.json
Found no files to cleanup
```

## Command line usage:

```
    Usage:

        chloe <command> [<options>]

    Commands:

        list             lists all files which are deletable
        dispatch         deletes any files which are deemed redundant


    Options:

        -i --input       sets the input JSON file, default is "bower.json"
        -f --force       force delete without prompting user
        -v --version     prints the application version
        -h --help        prints this help menu


    Version:

        1.0.0

```
