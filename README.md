# songcword
## Motivation 
*Curious how many unique words pop songs use*

## Installation
```bash
  git clone https://github.com/aluminick/songcword.git
  go build songcword.go
```
## Usage
Flags:
  * -input-file=/path/to/lyrics.txt. **Required**
  * -format=txt|json. (Optional, Default=txt)

Examples:
```bash
  ./songcword -input-file=lyrics.txt -format=txt
```
  Output:
```bash
    ...
    believe - 2
    first - 2
    hit - 1
    night - 1
    made - 1
    father - 1
    believer - 1
    get - 2
    steps - 1
    track - 5
    fake - 1
    baby - 1
    players - 2
    formalize - 1

    Original Word Count:  504
    Unique Word Count: 166
    less 67.06349%
```
```bash
    ./songcword -input-file=lyrics.txt -format=json
```
  Output:
```bash
    {"a":12,"about":1,"ain't":2,"altar":5,"always":1,"and":9,"another":5,"baby":1,"be":3,"being":1,"believe":2,"believer":1,"believes":1,"best":1,"bridegroom":1,"but":6,"can":1,"can't":7,"catchy":1,"cause":5,"chance":1,"child":1,"clean":1,"clouds":1,"come":2,"communion":1,"cross":1,"crushing":2,"desi":1,"do":1,"don't":8,"down":1,"drama":1,"explain":1,"fake":1,"father":1,"father's":2,"feels":6,"first":2,"fleshy":1,"fools":2,"formalize":1,"gave":1,"get":2,"go":4,"god":7,"goin":1,"got":1,"gotta":1,"hardest":1,"he":3,"he'll":1,"hear":1,"heart":1,"hesi":1,"hit":1,"hold":29,"holy":17,"honor":1,"how":1,"i":19,"i'll":1,"i'm":2,"if":1,"in":7,"is":3,"it":4,"it's":1,"jet":1,"jetski":1,"joe":1,"know":9,"leaving":2,"lefty":1,"let's":1,"life":2,"like":9,"lionel":1,"lot":1,"love":1,"made":1,"make":1,"making":1,"me":31,"men":2,"messi":1,"messy":1,"might":2,"my":2,"name":1,"next":1,"night":1,"nirvana":1,"no":9,"now":1,"of":1,"on":5,"opens":1,"or":1,"oscar":1,"out":1,"parlay":1,"part":1,"pesci":1,"pimps":2,"players":2,"pleases":1,"praises":1,"proud":2,"rapper":1,"rent":1,"river":1,"runnin":5,"rush":2,"saint":1,"say":7,"second":5,"see":2,"short":1,"sing":1,"sinners":1,"sky":1,"snack":1,"so":7,"son":1,"speed":1,"spots":1,"stand":1,"star":5,"step":1,"steps":1,"suffer":1,"take":2,"takes":1,"temper":1,"that":6,"the":36,"they":5,"think":1,"to":9,"too":2,"touch":1,"track":5,"trip":1,"trust":1,"tween":2,"union":1,"up":2,"us":2,"vespas":1,"wait":5,"wanna":2,"water":2,"way":8,"we":4,"we're":2,"weed":1,"week":1,"well":1,"when":4,"wise":2,"with":2,"yeah":1,"you":13,"young":2,"your":2}

Original Word Count:  504
Unique Word Count: 166
less 67.06349%
```

You can also output it to file
```bash
  ./songcword -input-file=lyrics.txt -format=json > words.json
```
TODO:
1. testing