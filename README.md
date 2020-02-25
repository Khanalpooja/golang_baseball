Programming Languages (CS524) 
Programming Assignment I
Feb 20, 2020 




Prepared By : Pooja Khanal , A25288740




Includes : 
* Software Description Document
* Software Test Plan Document 

























Table of contents 


Software Description Document        4
System Overview        4
Referenced Documents        4
Concepts of Execution        4
Operations/Computations        4
User input and Prompts        4
Abstract Data type (ADT)        5
Player structure        5
Code Outline        5
List of Functions used        5
Software Test Plan Document        7
    Test Procedures        7
Procedure followed for Testing and Testing Strategies        7
Black box testing        7
White box testing        7
Test cases :        7
Tested Environments:        7
       Sample Runs
       Assumptions        9
________________
Software Description Document 
1. System Overview 
A program is developed in GO programming language to read Baseball Player objects from an input file into a list/collection. The players are stored in collection in such a way that when they are displayed to the screen, they are in sorted order by lastname( then firstname to break name ordering tiles) . Each line of the data file will contain one player's name and stats.
{ firstname lastname plateappearances atbats singles doubles triples homeruns walks hitbypitch} 
Also, It is possible there are errors in the data lines. The errors are checked for that each numeric value is valid, and that all 10 values are on the line. The program reads until the end of the file to read for players.
2. Referenced Documents 
[1] Programming Assignment I , CS-424/524 Spring 2020 
[2] https://tour.golang.org/
3.  Concepts of Execution 
   1. Operations/Computations
* batting average: this is the sum of the hits (singles, doubles, triples, home runs) divided by the number of at-bats.
* slugging percentage The slugging percentage is a weighted average. It is computed by:
𝑠𝑙𝑢𝑔𝑔𝑖𝑛𝑔 = (𝑠𝑖𝑛𝑔𝑙𝑒𝑠+2∗𝑑𝑜𝑢𝑏𝑙𝑒𝑠+3∗𝑡𝑟𝑖𝑝𝑙𝑒𝑠+4∗h𝑜𝑚𝑒𝑟𝑢𝑛𝑠)/𝑛𝑢𝑚𝑏𝑒𝑟 𝑜𝑓 𝑎𝑡 𝑏𝑎𝑡𝑠
* on base percentage OBP is the sum of all hits, walks and hit-by-pitch divided by the number of plate appearances.
* Team’s overall batting average is also computed.

   1. User input and Prompts
   * The user is asked for inputfilename, Y/N to declare whether to write the output to file or not, and outputfilename if to be written. 
                                                        
   4. Abstract Data type (ADT)
   1. Player structure
The player structure has the following attributes. 
-{firstname lastname plateappearances atbats singles doubles triples homeruns walks hitbypitch} 
Each player is added to the list of players after reading from file, one by one 

   5. Code Outline
List of Functions used
   1. func (p Player) getBattingAverage() float64
   1. Purpose: Calculate the batting average of each player.Calculation mentioned above 
   2. Argument : Player object 
   3. Return Type : float64

      2. func (p Player) getSluggingPercentage() float64
      1. Purpose: Calculate the slugging percentage of each player. Calculation mentioned above
      2. Argument : Player object 
      3. Return Type : float64

         3. func (p Player) getOnBasePercentage() float64
         1. Purpose: Calculate the Base percentage of each player. Calculation mentioned above
         2. Argument : Player object 
         3. Return Type : float64

            4. func getPlayer(s []string, lineCount int) (Player, error)
            1. Purpose: //get the details of the player and create an instance of each Player Structure
            2. Argument : splitter and number of lines 
            3. Return Type : Player, error

               5. func checkErrors(s []string, lineCount int) error
               1. check for errors in the line,
               2. Line contains not enough data
               3. line %d: Line contains invalid numeric data
               4. Argument: splitter and number of lines 
               5. Return type : error

                  6. func getPlayers(fileName string) ([]Player, []error)
                  1. Read from the file and append each player to the Player Structure
                  2. Argument: filename
                  3. Player structure and Error Structure Array

                     7. func printErrorReport(errors []error)
                     1. Purpose: print the overall error report
                     2. Argument: Error Structure 
                     3. Return type: void

                        8. func printPlayerReport(players []Player)
                        1. Purpose: print the overall player report with the tabular structure
                        2. Argument: Player Structure
                        3. Return type: void

                           9. func handlePrintToFile(players []Player, errors []error)
                           1. Purpose: /Output stream handler to write to file
                           2. Argument: Player and Error Structure
                           3. Return Type : void

                              10. func main() {
                              1. Purpose: Main function to run the program. This is implemented in source.cpp
                              2. Arguments: Not Applicable 
                              3. Return type : Integer after successful completion of program  
                              4. Contains the test abstraction and calls the required test function for running the application. 










Software Test Plan Document
                              1. Test Procedures

                                 1. Purpose: To test the validity of software and to test if the software is acceptable to the market or not. For this program, various test cases are created 

                                 2. Procedure followed for Testing and Testing Strategies 

                                    1. Black box testing
                                    1. The input file was processed and seen if the output was written correctly to the output file or not. 
                                    2. Below mentioned test cases were run.

                                       2. White box testing 
                                       1. Modular approach was taken and each function was tested to see if it works or not. For such, the return type and return values were also compared. 
                                       2. After all the modules (functions) were tested, all the modules were integrated into one program and tested.

                                          2. Test cases : 
                                          1. File without errors 
                                          2. File with all type of errors 
                                          3. File with one type of error 
                                          4. File with line having both type of error
                                          5. Integers with long int ( instead of string) 
                                          6. Alphabetically sorted or not
3. Tested Environments:
                                          1. macbook pro 2015, go version go1.13.8 darwin/amd64
                                          2. go version go1.11.5 windows/amd64

4. Sample Runs
 Command to compile and build: go run *.go 
                                             3. Sample Input I 


Hank Aaron 13941 12364 2294 624 98 755 1402 32
Chipper Jones 10614 8984 1671 2020 38 468 567 897
Ty Cobb 111 11434 3053 724 295 117 1249 94 889
Jonny Bench 8674 7658 1254 212 24 389 891 19
Tony Gwynn 2147483647 9288 2378 543 85 135 434 24
John Smoltz 1167 948 118 26 2 5 79 67 89
Output : 
  

                                             4. Sample Input II


Hank Aaron 13941 12364 2294 624 98 755 1402 32
Chipper Jones 10614 8984 1671 -2.345 38 468 1512
Ty Cobb 111 11434 3053 724 295 117 1249 94 889
Jonny Bench 8674 7658 1254 string 24 389 891 19
Tony Gwynn 2147483647 9288 2378 543 85 135 434 24
John Smoltz 11677777777777777777777 948 118 26 2 5 79 67 89
Output: 
  

5.Assumptions
                                             1. The file format (input) is always consistent. 
                                             2. There is no new line at the end of the input file.
                                             3. Round of to 3 decimal places for all float values.
