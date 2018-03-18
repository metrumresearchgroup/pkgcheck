# R CMD check notes

From the CLI R CMD check should point to a package tarball. From the simplest perspective
the command can be run as

```bash
R CMD check <pkg.tar.gz>
```

When run, this will automatically create a folder called `<pkgname>.Rcheck` to which the 
following are dumped:

* source extracted from tarball
* installed pkg directory
* extracted tests directory
* 00check.log - log file from the R CMD check
* 00install.out - installation

## example package test1

A simple package with a single function `my_median` and a test is created to test
the simplest set of checks when all is correct

The following show an example of what the logs look like:

### 00check.log

```text
* using log directory ‘/Users/devin/clients/amgen/genpkgs/test1.Rcheck’
* using R version 3.4.3 (2017-11-30)
* using platform: x86_64-apple-darwin17.4.0 (64-bit)
* using session charset: UTF-8
* checking for file ‘test1/DESCRIPTION’ ... OK
* checking extension type ... Package
* this is package ‘test1’ version ‘0.0.1’
* package encoding: UTF-8
* checking package namespace information ... OK
* checking package dependencies ... OK
* checking if this is a source package ... OK
* checking if there is a namespace ... OK
* checking for executable files ... OK
* checking for hidden files and directories ... OK
* checking for portable file names ... OK
* checking for sufficient/correct file permissions ... OK
* checking whether package ‘test1’ can be installed ... OK
* checking installed package size ... OK
* checking package directory ... OK
* checking DESCRIPTION meta-information ... OK
* checking top-level files ... OK
* checking for left-over files ... OK
* checking index information ... OK
* checking package subdirectories ... OK
* checking R files for non-ASCII characters ... OK
* checking R files for syntax errors ... OK
* checking whether the package can be loaded ... OK
* checking whether the package can be loaded with stated dependencies ... OK
* checking whether the package can be unloaded cleanly ... OK
* checking whether the namespace can be loaded with stated dependencies ... OK
* checking whether the namespace can be unloaded cleanly ... OK
* checking loading without being on the library search path ... OK
* checking dependencies in R code ... OK
* checking S3 generic/method consistency ... OK
* checking replacement functions ... OK
* checking foreign function calls ... OK
* checking R code for possible problems ... OK
* checking Rd files ... OK
* checking Rd metadata ... OK
* checking Rd cross-references ... OK
* checking for missing documentation entries ... OK
* checking for code/documentation mismatches ... OK
* checking Rd \usage sections ... OK
* checking Rd contents ... OK
* checking for unstated dependencies in examples ... OK
* checking examples ... NONE
* checking for unstated dependencies in ‘tests’ ... OK
* checking tests ... OK
  Running ‘testthat.R’
* checking PDF version of manual ... OK
* DONE
Status: OK
```

### 00install.out

```text
* installing *source* package ‘test1’ ...
** R
** preparing package for lazy loading
** help
*** installing help indices
** building package indices
** testing if installed package can be loaded
* DONE (test1)
```

### tests/testthat.Rout

```text
R version 3.4.3 (2017-11-30) -- "Kite-Eating Tree"
Copyright (C) 2017 The R Foundation for Statistical Computing
Platform: x86_64-apple-darwin17.4.0 (64-bit)

R is free software and comes with ABSOLUTELY NO WARRANTY.
You are welcome to redistribute it under certain conditions.
Type 'license()' or 'licence()' for distribution details.

R is a collaborative project with many contributors.
Type 'contributors()' for more information and
'citation()' on how to cite R or R packages in publications.

Type 'demo()' for some demos, 'help()' for on-line help, or
'help.start()' for an HTML browser interface to help.
Type 'q()' to quit R.

> library(testthat)
> library(test1)
> 
> test_check("test1")
══ testthat results  ═════════════════════════════════════════════════════════════
OK: 1 SKIPPED: 0 FAILED: 0
> 
> proc.time()
   user  system elapsed 
  0.937   0.062   0.865 
```


## Settings

* `--as-cran` - sets the check scenario to run as it will on CRAN

### Environment Variables

* `R_LIBS` - library paths R will look for packages
  * The library search path is initialized at startup from the environment variable 'R_LIBS' (which should be a colon-separated list of directories at which R library trees are rooted) followed by those in environment variable 'R_LIBS_USER'. Only directories which exist at the time will be included.
* `R_TESTS` - ?
* `NOT_CRAN` - allows tests to take arbitrary amounts of time

https://stackoverflow.com/questions/24387660/how-to-change-libpaths-permanently-in-r