# Brief Summary of K-mer Model v.s. Markov Model

## On Training Data

How well does each model correctly categorize the training data?

It appears that **Markov Model performs better than the equivalent K-mer model** (data shown below)

### 235.fa (positive strand)
| Model | Pos | Neg | Gnm |
| ----- | --- | --- | --- |
| 1st order Markov| 160 | 54 | 21 |
| 2-mer | 164 | 50 | 21 |

| Model | Pos | Neg | Gnm |
| ----- | --- | --- | --- |
| 2nd order Markov| 171 | 49 | 15 |
| 3-mer | 167 | 49 | 19 |

| Model | Pos | Neg | Gnm |
| ----- | --- | --- | --- |
| 3rd order Markov| 173 | 54 | 8 |
| 4-mer | 170 | 50 | 15 |

| Model | Pos | Neg | Gnm |
| ----- | --- | --- | --- |
| 4th order Markov| 188 | 43 | 4 |
| 5-mer | 176 | 51 | 8 |

| Model | Pos | Neg | Gnm |
| ----- | --- | --- | --- |
| 5th order Markov| 210 | 21 | 4 |
| 6-mer | 192 | 37 | 6 |

### 300.fa (negative strand)

| Model | Pos | Neg | Gnm |
| ----- | --- | --- | --- |
| 1st order Markov| 90 | 181 | 29 |
| 2-mer | 84 | 185 | 31 |

| Model | Pos | Neg | Gnm |
| ----- | --- | --- | --- |
| 2nd order Markov| 85 | 200 | 15 |
| 3-mer | 84 | 192 | 24 |

| Model | Pos | Neg | Gnm |
| ----- | --- | --- | --- |
| 3rd order Markov| 81 | 212 | 7 |
| 4-mer | 82 | 203 | 15 |

| Model | Pos | Neg | Gnm |
| ----- | --- | --- | --- |
| 4th order Markov| 66 | 229 | 5 |
| 5-mer | 78 | 214 | 8 |

| Model | Pos | Neg | Gnm |
| ----- | --- | --- | --- |
| 5th order Markov| 31 | 266 | 3 |
| 6-mer | 51 | 244 | 5 |

## On Test Sequence

How accurate does each model predict the states on testseq.fa?

 - 0th order: same performance (as expected... basically the same thing)
 - 1st and 2nd order: Markov model performs better than K-mer counter part
  - Higher orders: K-mer model performs better

  One thing to note is that the testseq in generated using k-mer model.


| Order | Accuracy |  K | Accuracy |
| ----- | -------- |  - | -------- |
| 0 | 0.313 | 1 | 0.313 |
| 1 | 0.444 | 2 | 0.366 |
| 2 | 0.478 | 3 | 0.443 |
| 3 | 0.501 | 4 | 0.535 |
| 4 | 0.521 | 5 | 0.580 |
| 5 | 0.526 | 6 | 0.585 |
