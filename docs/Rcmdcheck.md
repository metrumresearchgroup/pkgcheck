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


### Help Docs

```text
Usage: R CMD check [options] pkgs

Check R packages from package sources, which can be directories or
package 'tar' archives with extension '.tar.gz', '.tar.bz2',
'.tar.xz' or '.tgz'.

A variety of diagnostic checks on directory structure, index and
control files are performed.  The package is installed into the log
directory and production of the package PDF manual is tested.
All examples and tests provided by the package are tested to see if
they run successfully.  By default code in the vignettes is tested,
as is re-building the vignette PDFs.

Options:
  -h, --help            print short help message and exit
  -v, --version         print version info and exit
  -l, --library=LIB     library directory used for test installation
                        of packages (default is outdir)
  -o, --output=DIR      directory for output, default is current directory.
                        Logfiles, R output, etc. will be placed in 'pkg.Rcheck'
                        in this directory, where 'pkg' is the name of the
                        checked package
      --no-clean        do not clean 'outdir' before using it
      --no-codoc        do not check for code/documentation mismatches
      --no-examples     do not run the examples in the Rd files
      --no-install      skip installation and associated tests
      --no-tests        do not run code in 'tests' subdirectory
      --no-manual       do not produce the PDF manual
      --no-vignettes    do not run R code in vignettes nor build outputs
      --no-build-vignettes    do not build vignette outputs
      --ignore-vignettes    skip all tests on vignettes
      --run-dontrun     do run \dontrun sections in the Rd files
      --run-donttest    do run \donttest sections in the Rd files
      --use-gct         use 'gctorture(TRUE)' when running examples/tests
      --use-valgrind    use 'valgrind' when running examples/tests/vignettes
      --timings         record timings for examples
      --install-args=   command-line args to be passed to INSTALL
      --test-dir=       look in this subdirectory for test scripts (default tests)
      --no-stop-on-test-error   do not stop running tests after first error
      --check-subdirs=default|yes|no
                        run checks on the package subdirectories
                        (default is yes for a tarball, no otherwise)
      --as-cran         select customizations similar to those used
                        for CRAN incoming checking

The following options apply where sub-architectures are in use:
      --extra-arch      do only runtime tests needed for an additional
                        sub-architecture.
      --multiarch       do runtime tests on all installed sub-archs
      --no-multiarch    do runtime tests only on the main sub-architecture
      --force-multiarch run tests on all sub-archs even for packages
                        with no compiled code

By default, all test sections are turned on.
```


Note: R CMD check and R CMD build run R processes with --vanilla in which none of the user’s startup files are read. 
If you need R_LIBS set (to find packages in a non-standard library) you can set it in the environment: 
also you can use the check and build environment files (as specified by the environment variables 
R_CHECK_ENVIRON and R_BUILD_ENVIRON; if unset, files ~/.R/check.Renviron and ~/.R/build.Renviron are used) 
to set environment variables when using these utilities.


## development version of R

per: https://twitter.com/henrikbengtsson/status/982102597178814464

To reproduce #rstats-devel CRAN errors, add below to .travis.yml file:

env:
- R_KEEP_PKG_SOURCE=yes
- _R_S3_METHOD_LOOKUP_BASEENV_AFTER_GLOBALENV_=true
- _R_S3_METHOD_LOOKUP_USE_TOPENV_AS_DEFENV_=true