#/usr/bin/env python3

from typing import Final, List, Dict

DICE_COUNT: Final[int] = 3
DICE_SIDES: Final[int] = 6

total_count = DICE_SIDES ** DICE_COUNT

def init_roll() -> List[int]:
    return [1 for _ in range(DICE_COUNT)]

def advance_roll(roll: List[int]) -> None:
    for idx in range(len(roll) - 1, -1, -1):
        roll[idx] = roll[idx] + 1
        if roll[idx] > DICE_SIDES:
            roll[idx] = 1
        else:
            break

def measure_roll(roll: List[int]) -> int:
    return sum(roll)

def get_roll_key(roll: List[int]) -> str:
    return ''.join(map(str, sorted(roll)))

def main() -> None:
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

if __name__ == '__main__':
    main()
