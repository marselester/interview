# Probability theory

Table of content:

- [Rules](#rules)
- [Conditional probability](#conditional-probability)
- [Bayes' theorem](#bayes-theorem)
  - [Medical test problem](#medical-test-problem)
  - [Taxi problem](#taxi-problem)
- [Counting](#counting)
- [Binomial distribution](#binomial-distribution)
- [Normal distribution](#normal-distribution)

References:

- [Answering questions with data](https://crumplab.github.io/statistics/probability-sampling-and-estimation.html) by Matthew J. C. Crump
- [Discrete math](https://www.youtube.com/watch?v=4wV9xGJXFjg&list=PLHXZ9OQGMqxersk8fUxiUMSIx0DBqsKZS&index=60) by Dr. Trefor Bazett

## Rules

Rules of probability:

- `P(not A) = 1 - P(A)` all elementary events that don't belong to A
- `P(A or B) = P(A) + P(B) - P(A and B)` at least one of A or B will happen
- `P(A and B) = P(A | B) * P(B)` A and B both will happen
- `P(A | B) = P(A and B) / P(B)` A will happen given that B already happened

Sample space is a set of all possible elementary events, e.g., {x1, x2, x3, x4, x5}.
Every time we make an observation, the outcome will be one and only one of these elementary events.
Each elementary event x is assigned a probability P(x) — a number between 0 and 1.
The probabilities of the elementary events need to add up to 1 (law of total probability).

There are two red, two blue, and one green marbles.
Here is a probability distribution.

| marble   | probability
| ---      | ---
| x1 red   | P(x1) = 0.2
| x2 red   | P(x2) = 0.2
| x3 blue  | P(x3) = 0.2
| x4 blue  | P(x4) = 0.2
| x5 green | P(x5) = 0.2

Non-elementary event A "red marble" is a subset {x1, x2} of a sample space.
If any of elementary events x1, x2 occurs, then A also occurred.

```
P(A) = P(x1) + P(x2) = 0.2 + 0.2 = 0.4
```

Probability of getting a red marble is 2/5 because there are two desired marbles (red) and five in total.
Odds of getting a red marble is 2/3 because there are two favorable marbles (red) and three remaining (two blue and one green).

```
// Probability is a ratio of favorable over all possible outcomes.
p = P(A) = 2 / 5

// Odds is a ratio of favorable over unfavorable outcomes.
q = Q(A) = 2 / 3

p = q / (1 + q)
q = p / (1 - p)
```

`P(A or B)` — at least one of A or B occurred (union without double counting):

- `P(A) + P(B)` when events are non-overlapping (disjoint union), e.g., you can't be in two places at once (mutually exclusive events).
  `P(A or B)` is a probability of being in place A or in place B.
- `P(A) + P(B) - P(A and B)` when events are overlapping, e.g., you can know English and French.
  Subtraction of `P(A and B)` removes double counting.

What is a probablity of picking red marbles two times `P(red, red)`?
These two events might be:

- independent `P(red, red) = 2/5 * 2/5` (marbles were returned back)
- dependent `P(red, red) = 2/5 * 1/4` (marbles were not returned back).
  Note, since a red marble wasn't returned back after first event, there are 4 marbles left and only 1 is red.

## Conditional probability

`P(A | B)` — probability of event A occuring given that event B already happened (note the past tense).
Division narrows down the sample space from `1` (all events) to `P(B)` because non-B events are impossible (have zero probability).

```
P(A | B) = P(A and B) / P(B)
```

<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" preserveAspectRatio="xMidYMid meet" viewBox="147.4615384615384 144.38461538461542 178.6153846153847 120.15384615384619" width="174.62" height="116.15"><defs><path d="M258.46 203.46C258.46 229.78 237.09 251.15 210.77 251.15C184.45 251.15 163.08 229.78 163.08 203.46C163.08 177.14 184.45 155.77 210.77 155.77C237.09 155.77 258.46 177.14 258.46 203.46Z" id="b1JbNvCZND"></path><path d="M306.15 203.46C306.15 229.78 284.78 251.15 258.46 251.15C232.14 251.15 210.77 229.78 210.77 203.46C210.77 177.14 232.14 155.77 258.46 155.77C284.78 155.77 306.15 177.14 306.15 203.46Z" id="b4C8nuLWx3"></path><path d="M148.46 145.38L323.08 145.38L323.08 261.54L148.46 261.54L148.46 145.38Z" id="e2JEaWhHna"></path><text id="a1f2N2XTff" x="251.15" y="203.46" font-size="19" font-family="Open Sans" font-weight="700" font-style="normal" letter-spacing="0" alignment-baseline="before-edge" transform="matrix(1 0 0 1 -17.692307692307452 -23.07692307692301)" style="line-height:100%" xml:space="preserve" dominant-baseline="text-before-edge"><tspan x="251.15" dy="0em" alignment-baseline="before-edge" dominant-baseline="text-before-edge" text-anchor="middle">A A∩B B</tspan></text></defs><g><g><use xlink:href="#b1JbNvCZND" opacity="1" fill="#c5d528" fill-opacity="1"></use></g><g><use xlink:href="#b4C8nuLWx3" opacity="1" fill="#204eb5" fill-opacity="0.4"></use></g><g><g><use xlink:href="#e2JEaWhHna" opacity="1" fill-opacity="0" stroke="#b2b3cf" stroke-width="2" stroke-opacity="1"></use></g></g><g id="arevtAUEm"><use xlink:href="#a1f2N2XTff" opacity="1" fill="#000000" fill-opacity="0.5"></use></g></g></svg>

What is the probability of being an alcoholic (A) given that a patient is a man (B)?
There are 2.25% of adults men who are alcoholics.

```
// Probability of being a man.
P(B) = 0.5

// Probability of being an alcoholic and being a man.
P(A and B) = 0.0225

// Probability of being an alcoholic given that a patient is a man.
P(A | B) = P(A and B) / P(B) = 0.0225 / 0.5 = 0.045
```

## Bayes' theorem

Bayes' theorem is a way to relate different conditional probabilities.

Since intersections `P(A and B)`, `P(B and A)` are equal, one of conditional probabilities can be substituted.
Depending on a problem, sometimes `P(A | B)` computation is easier than `P(B | A)`.

```
P(A | B) = P(A and B) / P(B)
P(B | A) = P(B and A) / P(A)

P(B and A) = P(B | A) * P(A)
P(A | B) = (P(B | A) * P(A)) / P(B)
```

What is the probability of two children being girls if at least one is a girl?
There are four permutations: girl boy, girl girl, boy girl, boy boy.

```
// Probability of two girls.
P(A) = 1 / 4

// Probability of at least one girl.
P(B) = 3 / 4

// Probability of at least one girl given that two children are girls.
P(B | A) = 1

// Probability of two girls given that at least one is a girl.
P(A | B) = (P(B | A) * P(A)) / P(B) = (1 * 1/4) / 3/4 = 1/3
```

### Medical test problem

Probability of having a disease is 1%.
Patient is tested positive.

Medical tests might be inaccurate:

- 5% false positive rate — tested that a patient has a disease, but actually doesn't have it.
  In all situations when a patient doesn't have a disease,
  in 5% situations the test tells that a patient has a disease.
- 10% false negative rate — tested that a patient doesn't have a disease, but actually has it.
  In all situations when a patient has a disease,
  in 10% situations the test tells that a patient doesn't have a disease.

How confident are you in the test result?

```
// Probability of having a disease.
P(A) = 0.01

// Probability of being tested positive.
P(B) = ?

// Probability of having a disease given that a test is positive.
P(A | B) = (P(B | A) * P(A)) / P(B)

// Probability of testing positive given that a patient has a disease (true positive).
// 90% of times when test is positive a patient has a disease
// since false negative rate is 10%.
P(B | A) = 0.9
```

`P(B)` probability of being tested positive is a disjoint union of two cases:

- `P(B | A) * P(A)` — tested positive and having a disease
- `P(B | not A) * P(not A)` — tested positive and not having a disease

```
// Probability of not having a disease.
P(not A) = 1 - P(A) = 1 - 0.01 = 0.99

// Probability of testing positive given that a patient doesn't have a disease (false positive).
P(B | not A) = 0.05

// Probability of testing positive given that a patient has a disease.
P(B) = P(B | A) * P(A) + P(B | not A) * P(not A) = 0.9 * 0.01 + 0.05 * 0.99 = 0.0585

// Probability of having a disease given that a test is positive.
P(A | B) = (P(B | A) * P(A)) / P(B) = (0.9 * 0.01) / 0.0585 = 0.15

                P(true positive) * P(A)
------------------------------------------------------
P(true positive) * P(A) + P(false positive) * P(not A)
```

15% confidence of having a disease after a positive medical test result.
The second test should improve confidence.
Meaning of `P(B)` is changing to "tested positive twice in a row".

```
// Probability of testing positive twice given that a patient has a disease (true positive).
P(B | A) = 0.9 * 0.9 = 0.81

// Probability of testing positive twice given that a patient doesn't have a disease (false positive).
P(B | not A) = 0.05 * 0.05 = 0.0025

// Probability of testing positive twice given that a patient has a disease.
P(B) = P(B | A) * P(A) + P(B | not A) * P(not A) = 0.81 * 0.01 + 0.0025 * 0.99 = 0.010575

// Probability of having a disease given that two tests are positive.
P(A | B) = (P(B | A) * P(A)) / P(B) = (0.81 * 0.01) / 0.010575 = 0.76
```

76% confidence of having a disease after two positive medical tests in a row.

### Taxi problem

The following problem is from "Thinking, Fast and Slow" by Daniel Kahneman.

> A cab was involved in a hit-and-run accident at night.
> Two cab companies, the Green and the Blue, operate in the city.
>
> 85% of the cabs in the city are Green and 15% are Blue.
> A witness identified the cab as Blue.
> The court tested the reliability of the witness under the circumstances that
> existed on the night of the accident and concluded that
> the witness correctly identified each one of the two colors 80% of the time and
> failed 20% of the time.
>
> What is the probability that the cab involved in the accident was Blue rather than Green?

```
// Probability of a cab being Blue (15% of the cabs in the city).
P(A) = 0.15

// Probability of a cab being identified as Blue by a witness.
P(B) = ?

// Probability of a Blue cab being involved in the accident
// given that a witness said a cab was Blue.
P(A | B) = (P(B | A) * P(A)) / P(B)

// Probability of a cab being identified as Blue
// given that a cab was indeed Blue (true positive).
// 80% of times when a witness said it was Blue, a cab was actually Blue.
P(B | A) = 0.8
```

`P(B)` probability of a cab being identified as Blue is a disjoint union of two cases:

- `P(B | A) * P(A)` — identified as Blue and actually being Blue
- `P(B | not A) * P(not A)` — identified as Blue and actually being Green

```
// Probability of a cab being Green (85% of the cabs in the city).
P(not A) = 0.85

// Probability of a cab being identified as Blue
// given that a cab was Green (false positive).
// The witness failed to identify each one of the two colors 20% of the time.
P(B | not A) = 0.2

// Probability of a cab being identified as Blue by a witness.
P(B) = P(B | A) * P(A) + P(B | not A) * P(not A) = 0.8 * 0.15 + 0.2 * 0.85 = 0.29

// Probability of a Blue cab being involved in the accident
// given that a witness said a cab was Blue.
P(A | B) = (P(B | A) * P(A)) / P(B) = (0.8 * 0.15) / 0.29 = 0.41
```

41% confidence of a Blue cab being involved in the accident.

## Counting

There are 1000 (10 * 10 * 10) three-digit PIN-codes:

- 10 options (0..9) for first digit
- 10 options (0..9) for second digit
- 10 options (0..9) for third digit

There are six ways (3! = 3 * 2 * 1) to arrange three letters C, A, T:

- 3 options for first letter
- 2 options for second letter because first was already used
- 1 option because only one letter left

Counting branches from the root also gives 3! (six branches):

```
      _________________
     /        |        \
    C         A         T
   / \       / \       / \
  A   T     C   T     A   C
 /     \   /     \   /     \
T       A T       C C       A
```

There are six ways (permutations) of picking two letters out of three letters C, A, T
(order is important): CA, CT, AC, AT, TC, TA.

```
// Out of n=3 letters pick r=2.
n P r = n! / (n - r)! = 3! / (3 - 2)! = 6
```

There are three ways (combinations) of choosing two letters out of three letters C, A, T
(order doesn't matter, remove duplicates with r!): CA, CT, AT.

```
// Out of n=3 letters choose r=2.
n C r = n! / ((n - r)! * r!) = 3! / (3 - 2)! * 2! = 3
```

There are 34650 ways (330 * 35 * 3 * 1) to reorder the word MISSISSIPPI (four S, four I, two P, one M).
The order of repeating letters doesn't matter `11! / (4! * 4! * 2!)`.

```
// Out of n=11 letters choose r=4 to place letter S.
n C r = 11! / (7! * 4!) = 330

// There are 11-4=7 letters left after placing four S letters.
// Out of n=7 letters choose r=4 to place letter I.
n C r = 7! / (3! * 4!) = 35

// There are 7-4=3 letters left after placing four I letters.
// Out of n=3 letters choose r=2 to place letter P.
n C r = 3! / (1! * 2!) = 3

// There is 3-2=1 letter left after placing two P letters.
// Out of n=1 letters choose r=1 to place letter M.
n C r = 1! / 1! = 1
```

## Binomial distribution

The binomial distribution is discrete (histogram-like): probability of a specific value.

> On 1 face of each die there's a picture of a skull; the other 5 faces are all blank.
> If I roll all 20 dice, what's the probability that I'll get exactly 4 skulls?

Probability of rolling r=4 skulls out of n=20 times is 0.2 (20%),
assuming probability of a single die is p=1/6.

Formula for binomial distribution.

```
// 4 out of 20 combinations multiplied by
// probability of 4 skulls multiplied by
// probability of 16 non-skulls.
P(r | p,n) = n! / ((n - r)! * r!) * p**r * (1-p)**(n-r)
P(r | p,n) = 20! / (16! * 4!) * 1/6**4 * (1-1/6)**16 = 0.2022
```

On a histogram the binomial distribution with size n=20 (horizontal axis),
each bar depicts the probability of one specific outcome (vertical axis).
The heights of the bars sum to 1.

## Normal distribution

The normal distribution is continuous (smooth curve):
probability that the value lies within a particular range of values
(calculate the area under the curve).

The horizontal axis corresponds to the value of some variable (e.g., temperature),
and the vertical axis is probability density (how likely we are to observe that temperature).

A normal distribution is described with two parameters,
the mean (index of central tendency) and the standard deviation.
The standard deviation is a measure of variance.
When it is large, the values are not clustered around the mean (great variance from the mean).

The area under the curve for the normal distribution must equal 1.
68.3% of the area falls within one standard deviation of the mean,
95.4% — two standard deviations.

If mean is changed from 0 to 5, the bell curve shifts from left to right.
If standard deviation is increased, the peak of the distribution stays in the same place,
but the distribution gets wider.
