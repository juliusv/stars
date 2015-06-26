default: plot

.PHONY: fetch
fetch:
	mkdir -p out
	go run ./stars.go -owner=prometheus -repo=prometheus -github-token=$(GITHUB_TOKEN) > out/stars.dat

.PHONY: plot
plot: fetch
	gnuplot stars.plot

.PHONY: show
show: plot
	eog out/stars.png
