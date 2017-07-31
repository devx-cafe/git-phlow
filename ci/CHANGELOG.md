## Changelog
The new release of git-phlow adds support for windows in _beta_. Overall improvements have been made to increase stability and remove
various bugs and odd behaviours.

#### Features
- Windows is now support in a pre-release version
- Add Less and More browsing #150 @groenborg
- Add no-color option for windows #148 @kryptag

#### Improvements
- Use Scanner in stead of buffered Reader #149 @groenborg
- Add windows binary to GitHub release #151 @groenborg
- Added windows worker to concourse ci #144

#### Bug fixes
- Colors are buggy when terminal width is changed #131 @groenborg
- username and password are no longer printed #158 @groenborg
- version was stuck at 1.0.0 #167 @groenborg
- Windows release is now zip and not tar.gz #169 @groenborg
