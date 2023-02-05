#/usr/bin/env python3

from typing import Final, List, Dict
import math

DiceRoll = List[int]

DICE_COUNT: Final[int] = 3
DICE_SIDES: Final[int] = 6

total_count = DICE_SIDES ** DICE_COUNT

def init_roll() -> DiceRoll:
    return [1 for _ in range(DICE_COUNT)]

def advance_roll(roll: DiceRoll) -> None:
    for idx in range(DICE_COUNT - 1, -1, -1):
        roll[idx] = roll[idx] + 1
        if roll[idx] > DICE_SIDES:
            roll[idx] = 1
        else:
            break

def measure_roll(roll: DiceRoll) -> int:
    return sum(roll)

def get_roll_key(roll: DiceRoll) -> str:
    return ''.join(map(str, sorted(roll)))

def get_roll_idx(roll: DiceRoll) -> int:
    idx = 0
    k = 1
    for i in range(DICE_COUNT - 1, -1, -1):
        idx += k * (roll[i] - 1)
        k *= DICE_SIDES
    return idx

def get_roll_from_idx(idx: int) -> DiceRoll:
    roll = init_roll()
    residue = idx
    factor = total_count
    for i in range(DICE_COUNT):
        factor //= DICE_SIDES
        k, residue = divmod(residue, factor)
        roll[i] = k + 1
    return roll

def test_1() -> None:
    roll = init_roll()
    index: Dict[int, Dict[str, int]] = {}
    for _ in range(total_count):
        key = get_roll_key(roll)
        val = measure_roll(roll)
        index_item = index.get(val)
        if not index_item:
            index_item = {}
            index[val] = index_item
        index_item[key] = index_item.get(key, 0) + 1
        advance_roll(roll)

    check_sum = 0
    for val in index.keys():
        index_item = index[val]
        val_cnt = sum(index_item.values())
        items = [f'{key}: {cnt}' for key, cnt in index_item.items()]
        text = '  '.join(items)
        check_sum = check_sum + val_cnt
        print(f'{val:2}: {val_cnt:4} #  {text}')
    if check_sum != total_count:
        raise RuntimeError('mismatch')

def test_2() -> None:
    roll = init_roll()
    index: Dict[int, List[int]] = {}
    for _ in range(total_count):
        idx = get_roll_idx(roll)
        val = measure_roll(roll)
        index_item = index.get(val)
        if not index_item:
            index_item = []
            index[val] = index_item
        index_item.append(idx)
        advance_roll(roll)

    width = math.ceil(math.log10(total_count))
    for val in index.keys():
        index_item = index[val]
        data = ' '.join(map(lambda t: str(t).rjust(width), index_item))
        diffs = [index_item[i] - index_item[i - 1] for i in range(1, len(index_item))]
        diffs_data = ' '.join(map(lambda t: str(t).rjust(width - 1), diffs))
        print(f'{val:2}: {data}')
        print(f'  {diffs_data}')

if __name__ == '__main__':
    test_2()
