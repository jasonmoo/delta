#delta

a delta encoder/decoder for integer ranges.  Reads writes from/to STDIN/STDOUT

####Usage:
  -d=false: compress the range supplied


###Example
    jason@mbp ~: go get github.com/jasonmoo/delta

	jason@mbp ~: seq 1 10
	1
	2
	3
	4
	5
	6
	7
	8
	9
	10

	jason@mbp ~: seq 1 10 | delta
	1-10

	jason@mbp ~: seq 1 10 | delta | delta -d
	1
	2
	3
	4
	5
	6
	7
	8
	9
	10

[LICENSE](https://raw.githubusercontent.com/jasonmoo/delta/master/LICENSE)

