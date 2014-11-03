#delta

a simple delta encoder/decoder for simple integer ranges

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
