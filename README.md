# Go Math

CLI utility to learn multiplication.

```
$ ./go_math --help
Usage of ./go_math:
  -f int
    	range left border (shorthand) (default 1)
  -from int
    	range left border (default 1)
  -t int
    	range right border (shorthand) (default 9)
  -to int
    	range right border (default 9)
  -l int
    	number of questions (shorthand) (default 20)
  -limit int
    	number of questions (default 20)
```

Example:
```
$ ./go_math --from=5 --to=8 --limit=8 
Learn range from 5 to 8 with limit 8
8 x 7 = 56
Correct
5 x 7 = 35
Correct
6 x 6 = 36
Correct
7 x 7 = ?
Correct answer is 49
7 x 6 = 43
Correct answer is 42
6 x 7 = 42
Correct
5 x 5 = 25
Correct
5 x 6 = 30
Correct
There are 2 incorrect answers, repeat? y/n: y
7 x 7 = 49
Correct
7 x 6 = 42
Correct
Done
```
