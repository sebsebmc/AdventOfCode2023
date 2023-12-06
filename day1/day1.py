import re
test = """two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
"""
with open('input.txt', 'r') as f:
    lines = f.readlines() # test.splitlines() 
    sum = 0
    lookup= {"0":0,"1":1,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,
            "one":1,"two":2,"three":3,"four":4,"five":5,"six":6,"seven":7,"eight":8,"nine":9}
    for line in lines:
        if line != "\n":
            print(line.strip())
            m = re.findall(r"(?=(\d|one|two|three|four|five|six|seven|eight|nine))", line.strip())
            print(m)
            print(lookup[m[0]]*10+lookup[m[-1]])
            sum += (lookup[m[0]]*10)+(lookup[m[-1]])

    print(sum)