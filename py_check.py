#!/usr/bin/env python3

from typing import List, Tuple, Dict
from argparse import ArgumentParser
import math

Vector = Tuple[float]

def main() -> None:
    parser = ArgumentParser()
    parser.add_argument('rolls')
    args = parser.parse_args()
    dice_count, dice_sides = map(int, args.rolls.split('d'))
    # dice_count = 3
    # dice_sides = 6

    # print('### Brute')
    # do_brute_variant(dice_count, dice_sides)
    # print('### Smart')
    do_smart_variant(dice_count, dice_sides)

def do_brute_variant(dice_count: int, dice_sides: int) -> None:
    min_val = dice_count
    max_val = dice_count * dice_sides
    mid_val = (min_val + max_val) // 2

    normal = get_normal(dice_count)
    points = prepare_points(dice_count, dice_sides)
    thresholds = prepare_thresholds(dice_count, dice_sides, mid_val - min_val + 1, normal)

    total = dice_sides ** dice_count
    diag_len = norm((dice_sides, ) * dice_count)
    eps = diag_len / dice_sides / 1000
    for i, threshold in enumerate(thresholds):
        portion = collect_points(normal, threshold + eps, points)
        fraction = len(portion) / total
        print(f'{i + min_val:2}: {len(portion):4} / {fraction:.4}')

def get_normal(dice_count: int) -> Vector:
    normal = (1, ) * dice_count
    return (1 / norm(normal),) * dice_count

def prepare_points(dice_count: int, dice_sides: int) -> List[Vector]:
    total = dice_sides ** dice_count
    return [generate_point(i, dice_count, dice_sides) for i in range(total)]

def norm(v: Vector) -> float:
    return math.sqrt(dot(v, v))

def dot(a: Vector, b: Vector) -> float:
    sum = 0
    for a_i, b_i in zip(a, b):
        sum += a_i * b_i
    return sum

def get_distance(normal: Vector, point: Vector) -> float:
    return dot(normal, point)

def generate_point(idx: int, dice_count: int, dice_sides: int) -> Vector:
    items = []
    residue = idx
    factor = dice_sides ** (dice_count - 1)
    for _ in range(dice_count):
        part, residue = divmod(residue, factor)
        factor /= dice_sides
        items.append(part)
    return tuple(items)

def prepare_thresholds(dice_count: int, dice_sides: int, count: int, normal: Vector) -> List[float]:
    vec = [0] * dice_count
    ret: List[float] = []
    for _ in range(count):
        ret.append(dot(vec, normal))
        for idx in range(dice_count - 1, -1, -1):
            if vec[idx] + 1 < dice_sides:
                vec[idx] += 1
                break
    return ret

def point_to_str(point: Vector) -> str:
    return ''.join(map(lambda t: str(int(t) + 1), point))

def print_points(points: List[Vector]) -> str:
    keys = map(point_to_str, points)
    index: Dict[str, int] = {}
    for key in keys:
        k = ''.join(sorted(key))
        index[k] = index.get(k, 0) + 1
    items = (f'{k}({v})' for k, v in index.items())
    return ' '.join(items)

def check_point(point: Vector, normal: Vector, threshold: float) -> bool:
    value = dot(normal, point)
    return value <= threshold

def collect_points(normal: Vector, threshold: float, points: List[Vector]) -> List[Vector]:
    return [*filter(lambda t: check_point(t, normal, threshold), points)]

def get_volume(t: float, n: int) -> float:
    sum = 0.0
    factorials = [1] * (n + 1)
    for i in range(2, n + 1):
        factorials[i] = factorials[i - 1] * i
    sign = 1
    for i in range(math.floor(t) + 1):
        coeff = factorials[n] / (factorials[i] * factorials[n - i])
        part = (t - i) ** n
        sum += sign * coeff * part
        sign = -sign
    return sum / factorials[n]

def measure_volume(value: float, dice_count: int, dice_sides: int) -> float:
    t = value * math.sqrt(dice_count) / dice_sides
    vol = get_volume(t, dice_count)
    return vol

def do_smart_variant(dice_count: int, dice_sides: int) -> None:
    min_val = dice_count
    max_val = dice_count * dice_sides
    mid_val = (min_val + max_val) // 2

    grand: float = dice_sides ** dice_count
    prev_volume: int = 0
    curr_dist = 1

    vol_diffs = prepare_diffs(dice_count)
    vol_counts = [0] * len(vol_diffs)
    item_count = 0

    for val in range(min_val, mid_val + 1):
        curr_volume = get_volume(curr_dist / dice_sides, dice_count) * grand
        diff_volume = curr_volume - prev_volume

        new_items, filled_items = distribute_volume(diff_volume, vol_diffs, vol_counts)
        item_count += filled_items

        prev_volume = curr_volume
        curr_dist += 1

        print(f'{val:2}: {new_items:4}')


def prepare_diffs(dice_count: int) -> List[float]:
    volumes = [get_volume(t, dice_count) for t in range(1, dice_count + 1)]
    for i in range(dice_count - 1, 0, -1):
        volumes[i] = volumes[i] - volumes[i - 1]
    return volumes

def distribute_volume(volume: float, diffs: List[float], counts: List[int]) -> Tuple[int, int]:
    residue = volume
    new_items = 0
    filled_items = 0
    for i in range(len(diffs) - 1, 0, -1):
        cc = counts[i]
        if cc > 0:
            diff = cc * diffs[i]
            residue -= diff
            counts[i] = 0
            assert residue > 0
            if i < len(diffs) - 1:
                counts[i + 1] = cc
            else:
                filled_items = cc
    assert residue > 0
    t = residue / diffs[0]
    new_items = round(t)
    counts[1] = new_items
    return new_items, filled_items

def prepare_pickers(dice_count: int, dice_sides: int, count: int) -> List[Vector]:
    ret: List[Vector] = []
    vec = (0,) * dice_count
    for _ in range(count):
        for idx in range(dice_count - 1, -1, -1):
            if vec[idx] + 1 <= dice_sides:
                vec[idx] += 1
                break
        ret.append(tuple(vec))
    return ret

if __name__ == '__main__':
    main()
