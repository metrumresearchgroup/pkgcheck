# Library Paths

Description

.libPaths gets/sets the library trees within which packages are looked for.

Usage

.libPaths(new)

.Library
.Library.site
Arguments

new a character vector with the locations of R library trees. Tilde expansion (path.expand) is done, 
and if any element contains one of *?[, globbing is done where supported by the platform: see Sys.glob.

Details

.Library is a character string giving the location of the default library, the ‘library’ subdirectory of R_HOME.

.Library.site is a (possibly empty) character vector giving the locations of the site libraries, by default the ‘site-library’ subdirectory of R_HOME (which may not exist).

.libPaths is used for getting or setting the library trees that R knows about (and hence uses when looking for packages). If called with argument new, the library search path is set to the existing directories in unique(c(new, .Library.site, .Library)) and this is returned. If given no argument, a character vector with the currently active library trees is returned.

How paths new with a trailing slash are treated is OS-dependent. On a POSIX filesystem existing directories can usually be specified with a trailing slash: on Windows filepaths with a trailing slash (or backslash) are invalid and so will never be added to the library search path.

The library search path is initialized at startup from the environment variable R_LIBS (which should be a colon-separated list of directories at which R library trees are rooted) followed by those in environment variable R_LIBS_USER. Only directories which exist at the time will be included.

By default R_LIBS is unset, and R_LIBS_USER is set to directory ‘R/R.version$platform-library/x.y’ of the home directory (or ‘Library/R/x.y/library’ for CRAN macOS builds), for R x.y.z.

.Library.site can be set via the environment variable R_LIBS_SITE (as a non-empty colon-separated list of library trees).

Both R_LIBS_USER and R_LIBS_SITE feature possible expansion of specifiers for R version specific information as part of the startup process. The possible conversion specifiers all start with a % and are followed by a single letter (use %% to obtain %), with currently available conversion specifications as follows:

%V
R version number including the patchlevel (e.g., 2.5.0).

%v
R version number excluding the patchlevel (e.g., 2.5).

%p
the platform for which R was built, the value of R.version$platform.

%o
the underlying operating system, the value of R.version$os.

%a
the architecture (CPU) R was built on/for, the value of R.version$arch.

(See version for details on R version information.)

Function .libPaths always uses the values of .Library and .Library.site in the base namespace. .Library.site can be set by the site in ‘Rprofile.site’, which should be followed by a call to .libPaths(.libPaths()) to make use of the updated value.

For consistency, the paths are always normalized by normalizePath(winslash = "/").

Value

A character vector of file paths.

References

Becker, R. A., Chambers, J. M. and Wilks, A. R. (1988) The New S Language. Wadsworth & Brooks/Cole.

See Also

library

Examples

.libPaths()                 # all library trees R knows about