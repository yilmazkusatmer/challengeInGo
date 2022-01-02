Write a programm that check if a space-separated string (second parameter) against 
the structure of a pattern passed in the form of 
individual characters as the first parameter.

| Input pattern | input text      | Result |
|---------------|-----------------|--------|
| xyxy          | tom tim tom tim | true   |
| xyyy          | foo fan fan foo | false  |
| xyyx          | tom tim tim tom | true   |
| xxxx          | tom tom tom tim | false  |



