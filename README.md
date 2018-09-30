# Introduction

This is some sample code that demonstrates using Prometheus Kubernetes discovery and scraping packages in a standalone binary.

There are other pieces not implemented, but this is useful for getting started.

The example is setup to run as a pod in Kubernetes and will discover all nodes, endpoints, services, pods, etc.. and try to scrape them using some hard-coded target options (/metrics).

# Demo

You need a Kubernetes cluster available and setup so that kubectl works.

```bash
$ make run
```

You should see it discover various targets, scrape them and output the scraped values to the console.

# Building

```$bash

$ make
$ make run

```

# License
MIT

