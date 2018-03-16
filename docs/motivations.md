# Motivation

The primary motivation around the development of the package check tooling is to provide a system to run checks against cohorts
of packages to assess the compatability across packages. As development needs extend beyond CRAN, or across package versions
not tied to the most recent version of all packages (which CRAN checks), tooling to mimic the CRAN checks can provide the 
same degree of confidence. Furthermore, CRAN is currently going through a number of "growing-pains" associated with the exponential
expansion of packages they are supporting. Certain foundational packages, such as `processx` have been stuck for months in 
limbo, without dialog with the CRAN maintainers, causing hoops such as embedding into other packages https://github.com/r-lib/processx/issues/94. 

To address this, a flexible system to run similar checks to those CRAN provides, while allowing customization as to the 
scrutiny of the checks will be important to confidently construct package cohorts. Likewise, to support environments with
rigorous change control processes and/or continuous integration activities, minimizing runtime dependencies is valuable.

## Prior/ongoing efforts

R CMD check execution

* `devtools::check()`
* `rcmdcheck` package

## Differences

`devtools::check()` is built around checking an in-development package in a particular folder on a developers computer. 
This tooling is expected to check package cohorts provided in a runtime or in a project folder, hence is different in scope.

The rcmdcheck package is designed around providing optimal feedback for a developer working on their specific package.
The package has not had much development activity, and still has low (~37%) test coverage in comparison to most of the 
r-lib organization packages (generally > 80%). Feature-wise, the core features do align with the objectives of 
checking a specific package, however no functionality towards checking multiple packages in a cohort is provided. 
Likewise, new feature development is not prioritized; for example, https://github.com/r-lib/rcmdcheck/issues/12,
discussed in 2016, shows that features that are important to managing scenarios other than submission to CRAN,
have not had progress. Most importantly, `rcmdcheck` is not actively following the ongoing development progress within r-lib.
For example, pointing to a old remote version of callr 1 major version behind.

rcmdcheck could potentially be 'adopted' by amgen or others willing to take on maintanence, 
however still brings up the issue of minimizing runtime dependencies, as well as implementing features around
parallel checks and needs specific to amgen (eg passing all requirements to submit to CRAN is not a problem). 