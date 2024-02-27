from temporalio import activity
from sorts.sorts import *
from shared import SortDetails

@activity.defn
def run_sort(detail: SortDetails) -> str:
    print(f"Start sorting {detail.numbers} with {detail.name} sort")
    match detail.name:
        case "bubble":
            bubble(detail.numbers)
        case "quick":
            quick(detail.numbers, 0, len(detail.numbers) - 1)
        case _:
            print(f"{detail.name} sort has not been implemented")
    print(f"End sorting {detail.numbers} with {detail.name} sort")
    return f"{detail.numbers}"
