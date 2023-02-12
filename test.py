#/usr/bin/env python3

from typing import List, Dict
from argparse import ArgumentParser

DiceRoll = List[int]

def measure_roll(roll: DiceRoll) -> int:
    return sum(roll)

def get_roll_idx(roll: DiceRoll, dice_count: int, dice_sides: int) -> int:
    idx = 0
    k = 1
    for i in range(dice_count - 1, -1, -1):
        idx += k * (roll[i] - 1)
        k *= dice_sides
    return idx

def get_roll_key(roll: DiceRoll) -> str:
    return ''.join(map(str, sorted(roll)))

def get_roll_from_idx(idx: int, dice_count: int, dice_sides: int) -> DiceRoll:
    roll = [0 for _ in range(dice_count)]
    residue = idx
    factor: int = dice_sides ** dice_count
    for i in range(dice_count):
        factor //= dice_sides
        k, residue = divmod(residue, factor)
        roll[i] = k + 1
    return roll

def generate_test_data(dice_sides: int, dice_count: int) -> None:
    index: Dict[int, Dict[str, int]] = {}
    total_count = dice_sides ** dice_count
    for i in range(total_count):
        roll = get_roll_from_idx(i, dice_count, dice_sides)
        key = get_roll_key(roll)
        val = measure_roll(roll)
        index_item = index.get(val)
        if not index_item:
            index_item = {}
            index[val] = index_item
        index_item[key] = index_item.get(key, 0) + 1

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

if __name__ == '__main__':
    parser = ArgumentParser()
    parser.add_argument('schema')
    args = parser.parse_args()
    dice_count, dice_sides = map(int, args.schema.split('d'))
    print(dice_count, dice_sides)
    generate_test_data(dice_sides=dice_sides, dice_count=dice_count)
