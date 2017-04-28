# Changelog - plumber

### 0.0.1-beta4

__Changes__

- `plumber` command line

    - add support for package import into its arguments.

      Before to select a type it was: '[]byte', 'semver.Version', '*my.PointerType'

      Now, you can use: 'github.com/mh-cbon/semver/*my.PointerType'

    - go:generate: fixed the template!
    - go:generate: When the output is a file, it is automatically go fmted.

- Other fixes
  - updated cli help to add the new import capability
  - updated README to add the new import capability
  - improved byteStream error message when a wrong pipe is connected on it.
  - updated the demo so it does work.



__Contributors__

- mh-cbon

Released by mh-cbon, Fri 28 Apr 2017 -
[see the diff](https://github.com/mh-cbon/plumber/compare/0.0.1-beta3...0.0.1-beta4#diff)
______________

### 0.0.1-beta3

__Changes__

- README: fix

__Contributors__

- mh-cbon

Released by mh-cbon, Thu 27 Apr 2017 -
[see the diff](https://github.com/mh-cbon/plumber/compare/0.0.1-beta2...0.0.1-beta3#diff)
______________

### 0.0.1-beta2

__Changes__

- README: improved demo

__Contributors__

- mh-cbon

Released by mh-cbon, Thu 27 Apr 2017 -
[see the diff](https://github.com/mh-cbon/plumber/compare/0.0.1-beta1...0.0.1-beta2#diff)
______________

### 0.0.1-beta1

__Changes__

- Fix and improvements in the README

__Contributors__

- mh-cbon

Released by mh-cbon, Thu 27 Apr 2017 -
[see the diff](https://github.com/mh-cbon/plumber/compare/0.0.1-beta...0.0.1-beta1#diff)
______________

### 0.0.1-beta

__Changes__

- Initialize the project.

__Contributors__

- mh-cbon

Released by mh-cbon, Thu 27 Apr 2017 -
[see the diff](https://github.com/mh-cbon/plumber/compare/78db9e516e383d770dfeac358221809f0d4e1528...0.0.1-beta#diff)
______________


