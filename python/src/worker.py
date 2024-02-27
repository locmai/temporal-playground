import asyncio
import concurrent.futures
from temporalio.client import Client
from temporalio.worker import Worker

# Import the activity and workflow from our other files
from activities.sort import run_sort
from workflows.sort import SortWorkflow
from shared import SORT_TASK_QUEUE_NAME

async def main():
    # Create client connected to server at the given address
    client = await Client.connect("localhost:7233")

    # Run the worker
    with concurrent.futures.ThreadPoolExecutor(max_workers=100) as activity_executor:
        worker = Worker(
          client,
          task_queue=SORT_TASK_QUEUE_NAME,
          workflows=[SortWorkflow],
          activities=[run_sort],
          activity_executor=activity_executor,
        )
        await worker.run()

if __name__ == "__main__":
    asyncio.run(main())
