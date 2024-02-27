from dataclasses import dataclass

SORT_TASK_QUEUE_NAME = "SORT_TASK_QUEUE"

@dataclass
class SortDetails:
    name: str
    numbers: list[int]
