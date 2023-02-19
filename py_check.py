#!/usr/bin/env python3

from typing import List, Tuple, Dict
from argparse import ArgumentParser
import numpy as np
import math

Vector = Tuple[float]

def main() -> None:
    # parser = ArgumentParser()
    # parser.add_argument('rolls')
    # args = parser.parse_args()
    # dice_count, dice_sides = map(int, args.rolls.split('d'))

    dice_count = 2
    dice_sides = 6
    min_val = dice_count
    max_val = dice_count * dice_sides
    mid_val = (min_val + max_val) // 2

    normal = get_normal(dice_count)
    points = prepare_points(dice_count, dice_sides)
    thresholds = prepare_thresholds(dice_count, dice_sides, max_val - min_val + 1, normal)

    total = dice_sides ** dice_count
    diag_len = np.linalg.norm(np.ones(dice_count) * dice_sides)
    eps = diag_len / dice_sides / 1000
    for i, threshold in enumerate(thresholds):
        portion = collect_points(normal, threshold + eps, points)
        to_check = measure_volume(threshold, dice_count, dice_sides)
        print(f'{i + min_val:2}: {len(portion):4} / {len(portion) / total:.3} || {to_check:.3}')

def get_normal(dice_count: int) -> Vector:
    normal = np.ones(dice_count)
    return normal / np.linalg.norm(normal)

def prepare_points(dice_count: int, dice_sides: int) -> List[Vector]:
    total = dice_sides ** dice_count
    return [generate_point(i, dice_count, dice_sides) for i in range(total)]

def get_distance(normal: Vector, point: Vector) -> float:
    return np.dot(normal, point)

def generate_point(idx: int, dice_count: int, dice_sides: int) -> Vector:
    items = []
    residue = idx
    factor = dice_sides ** (dice_count - 1)
    for _ in range(dice_count):
        part, residue = divmod(residue, factor)
        factor /= dice_sides
        items.append(part)
    return np.array(items)

def prepare_thresholds(dice_count: int, dice_sides: int, count: int, normal: Vector) -> List[float]:
    vec = np.zeros(dice_count)
    ret: List[float] = []
    for _ in range(count):
        ret.append(np.dot(vec, normal))
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
    value = np.dot(normal, point)
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

if __name__ == '__main__':
    main()
