---
title: Preparing a Release
date: 2018-05-28T16:27:37-07:00
draft: false
---

## Semantic Versioning

Releases are managed via annotated `git` tags.
Tags must follow the [Semantic Versioning](https://semver.org/) scheme.

## Performing a Release

```bash
git checkout ${COMMIT}
git tag -a v${MAJOR}.${MINOR}.${PATCH} -m "Release v${MAJOR}.${MINOR}.${PATCH}"
git push --tags
make push
```

e.g.

```bash
git tag -a v0.1.0 -m "Release v0.1.0"
```

## Performing a Prerelease

```bash
git checkout ${COMMIT}
git tag -a v${MAJOR}.${MINOR}.${PATCH}-${PRERELEASE} -m "Prerelease v${MAJOR}.${MINOR}.${PATCH}-${PRERELEASE}"
git push --tags
make push
```

e.g.

```bash
git tag -a v0.1.0-alpha.0 -m "Prerelease v0.1.0-alpha.0"
```
