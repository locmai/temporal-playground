from datetime import timedelta
from temporalio import workflow
from shared import SortDetails
# Import our activity, passing it through the sandbox
with workflow.unsafe.imports_passed_through():
    from activities.sort import run_sort

@workflow.defn
class SortWorkflow:
    @workflow.run
    async def run(self, detail: SortDetails) -> str:
        return await workflow.execute_activity(
            run_sort, detail, schedule_to_close_timeout=timedelta(seconds=10)
        )
