---
title: Documentation
date: 2018-05-28T16:27:37-07:00
draft: false
---

## Generating the Documentation

The documentation is generated using [Hugo](https://github.com/gohugoio/hugo).
To generate the documentation, simply run:

```bash
make docs
```

## Developing the Documentation

First, ensure that hugo is installed. Once you have done that, run:

```bash
cd ./hugo && hugo server --watch --verbose --disableFastRender
```

Now you can modify the source code and see changes in real time.
