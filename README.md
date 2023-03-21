# dicecalc

Probabilities calculation for dice rolls.

Roll a set of dices and measure (there are different ways) the result. What is the probability to get that result?

----

### Sum

For a set of dices take the sum of all values.

For a **2d6** rolls possible values are `2..12`. There are `6^2 = 36` total dice combinations. Let's show some probabilities.

- `2` is given by 1 set - `(1,1)`; probability is `1/36`
- similar for `12` - `(6,6)`
- `3` is given by 2 sets - `(1,2)`, `(2,1)`; probability is `2/36`
- similar for `11` - `(5,6)`, `(6,5)`
- `4` is given by 3 sets - `(1,3)`, `(2,2)`, `(3,1)`; probability is `3/36`
- similar for `10` - `(4,6)`, `(5,5)`, `(6,4)`
- `5` and `9` - `4/36`
- `6` and `8` - `5/36`
- `7` - `6/36`

The task is to calculate probabilities for an arbitrary **MdN** rolls.

### Max

For a set of dices take maximum of all values.

For a **2d6** rolls possible values are `1..6`. There are `6^2 = 36` total dice combinations. Let's show some probabilities.

- `1` is given by 1 set - `(1,1)`; probability is `1/36`
- `2` is given by 3 sets - `(1,2)`, `(2,1)`, `(2,2)`; probability is `3/36`
- `3` is given by sets - `(1,3)`, `(2,3)`, `(3,1)`, `(3,2)`, `(3,3)`; probability is `5/36`
- `4` - `7/36`
- `5` - `9/36`
- `6` - `11/36`

The task is to calculate probabilities for an arbitrary **MdN** rolls.

### Min

Just an opposite to **max**.

----

There are [console](./app/) and [server](./server/) wrappers over the library.

----

Yandex Cloud [function](./yc_function/) is deployed.

https://functions.yandexcloud.net/d4eppt029egh1r56fsq2
