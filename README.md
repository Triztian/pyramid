# Pyramid

In this challenge you have to find the quickest path down the pyramid
A pyramid is represented as such:
	1
	2 3
	4 5 6
Where you always start from the top and have to find your way to the bottom
You can only slide downwards and you can only slide to the 2 adjecent field downward
fx 1 -> [2, 3], 2 -> [4, 5] and 3 -> [5, 6] as their only possible moves

Usage:
  pyramid [command]

Available Commands:
  generate    Generate a pyramid and print it
  help        Help about any command
  slide       Calulates the quickest path down the pyramid.

Flags:
  -h, --help   help for pyramid

Use "pyramid [command] --help" for more information about a command.