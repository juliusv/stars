# Stars

GitHub star stats grapher. Uses the GitHub API to query all stars over the
history of a repository and graph the cumulative counts over time.

Just a quick hack so far.

## Prerequisites

You will need `make`, `gnuplot`, `go`, and optionally `eog` installed (for
viewing the resulting PNG).

## Using

Edit the `Makefile` to configure the GitHub repo you're interested in.

To only fetch data and generate a graph PNG:

```bash
GITHUB_TOKEN=XXXXXXXX make plot
```

To fetch data, generate a graph PNG, and open the image with `eog`:

```bash
GITHUB_TOKEN=XXXXXXXX make show
```
