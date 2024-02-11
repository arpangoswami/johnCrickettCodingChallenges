Follow along at https://codingchallenges.fyi/challenges/challenge-wc to build your own solution

arpangoswami@Arpans-MacBook-Air wc_CliTool % ./ccwc test.txt         
7145    58164   342190  test.txt
arpangoswami@Arpans-MacBook-Air wc_CliTool % go build ccwc.go        
arpangoswami@Arpans-MacBook-Air wc_CliTool % cat test.txt | ./ccwc -c
342190
arpangoswami@Arpans-MacBook-Air wc_CliTool % cat test.txt | ./ccwc -m
339292
arpangoswami@Arpans-MacBook-Air wc_CliTool % cat test.txt | ./ccwc -w
58164
arpangoswami@Arpans-MacBook-Air wc_CliTool % cat test.txt | ./ccwc -l
7145
arpangoswami@Arpans-MacBook-Air wc_CliTool % ./ccwc -c test.txt     
342190  test.txt
arpangoswami@Arpans-MacBook-Air wc_CliTool % ./ccwc -m test.txt 
339292  test.txt
arpangoswami@Arpans-MacBook-Air wc_CliTool % ./ccwc -w test.txt 
58164   test.txt
arpangoswami@Arpans-MacBook-Air wc_CliTool % ./ccwc -l test.txt 
7145    test.txt
arpangoswami@Arpans-MacBook-Air wc_CliTool % ./ccwc test.txt 
7145    58164   342190  test.txt