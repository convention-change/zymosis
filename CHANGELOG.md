# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.1.3](https://github.com/convention-change/zymosis/compare/1.1.2...v1.1.3) (2024-09-11)

### BREAKING CHANGE:

* This commit changes the indirect dependencies' versions and

### 🐛 Bug Fixes

* update golang.org/x packages to newer versions ([5b28f88b](https://github.com/convention-change/zymosis/commit/5b28f88be5bdc1b6dd4216ffa57c03c208e27ba1))

* downgrade Masterminds/sprig/v3 to 3.2.1 for compatibility ([d3f06896](https://github.com/convention-change/zymosis/commit/d3f06896857635f56e0ebf7e1359eef6795a15b6))

## [1.1.2](https://github.com/convention-change/zymosis/compare/1.1.1...v1.1.2) (2024-09-11)

### 👷‍ Build System

* bump github.com/sinlov-go/unittest-kit from 1.1.1 to 1.2.1 ([3cc83e33](https://github.com/convention-change/zymosis/commit/3cc83e338a23c7c4b162cd2f60ea605f1c019f3b))

* bump github.com/bar-counter/slog from 1.4.0 to 1.4.1 ([193a21ab](https://github.com/convention-change/zymosis/commit/193a21ababf3c13a2386e51531e3cca3acc791f3))

* bump github.com/Masterminds/sprig/v3 from 3.2.1 to 3.3.0 ([ea7e8813](https://github.com/convention-change/zymosis/commit/ea7e8813fa9ec047ce69ff77bf5de4e61c45a9a1))

## [1.1.1](https://github.com/convention-change/zymosis/compare/1.1.0...v1.1.1) (2024-09-11)

### 🐛 Bug Fixes

* streamline git hash handling and file permission assignments ([50ee7426](https://github.com/convention-change/zymosis/commit/50ee7426b5145a1603a6c0ed0d185ada42f5884d)), fix [#28](https://github.com/convention-change/zymosis/issues/28)

### 👷‍ Build System

* bump go module and update gomod configuration ([17a05614](https://github.com/convention-change/zymosis/commit/17a056147501504a8e371e5f27f020bd2f31957d))

## [1.1.0](https://github.com/convention-change/zymosis/compare/1.0.0...v1.1.0) (2024-05-20)

### ✨ Features

* update full of update build pipline ([dd8b06ab](https://github.com/convention-change/zymosis/commit/dd8b06ab4b9dea733dc5d18cd032aa1ab233625b)), feat [#20](https://github.com/convention-change/zymosis/issues/20)

### 👷‍ Build System

* change golang build version and update full doc ([00d2c9e7](https://github.com/convention-change/zymosis/commit/00d2c9e7074ce84cf87556e0d6f17be9d52797d6))

* change golangci/golangci-lint-action use version latest ([85760fe1](https://github.com/convention-change/zymosis/commit/85760fe12912fcbad94d066ef1104c2a43928adf))

* update golangci-lint version to v1.55.2 to fix build at go 1.21+ ([b429032b](https://github.com/convention-change/zymosis/commit/b429032b806b44190579b4d6139888eb7e13a986))

## 1.0.0 (2023-08-23)

### ✨ Features

* generate go project git_rev_parse ([11a060aa](https://github.com/convention-change/zymosis/commit/11a060aad16b971b37c790aac3b113dee43eee17)), feat [#1](https://github.com/convention-change/zymosis/issues/1)

* add git_command to get cwd path git info as branch hash ([8322eee3](https://github.com/convention-change/zymosis/commit/8322eee38f47dbfc023b93684a758f13e845487c))
