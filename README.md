## Tuenti challenge 8

These are my solutions to some of the problems from the 8th tuenti challenge (2018). All of them are in Go.

* Problem 5


### Problem 5 DNA Slicer

In this problem we received some slices of DNA and we had to find the ones that formed a double helix. For example,
there was TAC, TA, CGAT, GATCG, GAT and ATG. In that case the following pieces formed a double helix:

    TAC GAT
    TA CGAT

while GATCG and ATG were just noise. So the solution was 1, 2, 3, 5.

A way to work this one out was to think of the DNA as growing. Say that we had already matched TAC and CGAT

    TAC
      CGAT

If this couple belongs to a solution, then we have to find a DNA slice finishing in TA and a DNA slice starting
in GAT. We can choose either GATCG or GAT. If we choose GATCG, we would have to keep looking for a DNA slice
starting in CG.

    TAC GATCG
      CGAT

Only question is were to start. We can start with all the possible options, that is for slice TAC we look for
slices starting in 'TAC', for slices ending in 'T' and slices staring in 'AC', 'TA' and 'C', and 'TAC'.

The struct defined in go would be:

    type growingMatching struct {
    	startString, endString string
    	remainingParts map[int]bool
    }



