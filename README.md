# prospace
Prospace Backend Code Challenge - GALAXY MERCHANT TRADING GUIDE 

## Running program
1. Clone this repo
2. go build
3. ./Prospace

## System design solution
The program consist of 6 parts:
1. Reading input from cli
2. Categorizing each input line 
    - Whether it is a intergalactic to roman numeral translation, e.g "glob is I"
    - Or if it is a common metal value, e.g "glob glob Silver is 34 Credits"
    - Or it is a question, e.g "how many Credits is glob prok Silver ?"
3. Building intergalactic to roman numeral map
4. Building common metal value map
5. Building answer for each question
6. Print results