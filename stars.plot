set terminal png size 1200,1000
set xdata time
set timefmt "%s"
set output "out/stars.png"
set grid
plot "out/stars.dat" using 1:2 index 0 title "GitHub stars" with lines
