Decision: BenchmarkSplitOnWords-12 wins. Same result, better perf.


goos: darwin
goarch: amd64
BenchmarkFindWords-12                     	    5000	    306912 ns/op
BenchmarkFindWords2-12                    	    5000	    300958 ns/op
BenchmarkSplitOnNonLetters-12             	   30000	     49128 ns/op
BenchmarkSplitOnNonLettersOrNumbers-12    	   30000	     50589 ns/op
BenchmarkSplitOnWords-12                  	   20000	     59833 ns/op
PASS
ok  	                               	8.923s
