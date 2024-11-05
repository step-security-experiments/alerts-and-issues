import go

from CallExpr call
where call.getFunc().getName() = "unsafeFunctionName"  
select call, "This function call may be unsafe or require additional checks."
