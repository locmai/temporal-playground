import asyncio
from temporalio.client import Client

# Import the workflow from the previous code
from workflows.sort import SortWorkflow
from shared import SORT_TASK_QUEUE_NAME, SortDetails
import uuid

async def main():
    # Create client connected to server at the given address
    client = await Client.connect("localhost:7233")


    workflow_id = str(uuid.uuid4())
    # Test data here:
    detail = SortDetails("quick", [5,3,2,1,4,100])


    # Execute a workflow
    result = await client.execute_workflow(SortWorkflow.run, detail, id=workflow_id, task_queue=SORT_TASK_QUEUE_NAME)
    print(f"result: {result}")
    print(f"workflow_id: {workflow_id}")

if __name__ == "__main__":
    asyncio.run(main())
