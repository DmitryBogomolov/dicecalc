#!/usr/bin/env python3

from typing import List, Tuple, Dict
import numpy as np
import math

Vector = List[float]

def main() -> None:
    dice_count = 3
    dice_sides = 6
    normal = get_normal(dice_count)
    points = prepare_points(dice_count, dice_sides)

    thresholds = [
        get_distance(normal, (0, 0, 0)) + 0.0001,
        get_distance(normal, (0, 0, 1)) + 0.0001,
        get_distance(normal, (0, 0, 2)) + 0.0001,
        get_distance(normal, (0, 0, 3)) + 0.0001,
        get_distance(normal, (0, 0, 4)) + 0.0001,
        get_distance(normal, (0, 0, 5)) + 0.0001,
        get_distance(normal, (0, 1, 5)) + 0.0001,
        get_distance(normal, (0, 2, 5)) + 0.0001,
    ]

    portion = collect_points(normal, (thresholds[6], thresholds[7]), points)
    print(print_points(portion))

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

def check_point(point: Vector, normal: Vector, range: Tuple[float, float]) -> bool:
    value = np.dot(normal, point)
    return range[0] < value and value <= range[1]

def collect_points(normal: Vector, range: Tuple[float, float], points: List[Vector]) -> List[Vector]:
    return [*filter(lambda t: check_point(t, normal, range), points)]

if __name__ == '__main__':
    main()
