import go

from CallExpr call
where call.getFunc().getName() = "os.Open"
select call, "Using 'os.Open' without proper error handling can lead to resource leaks."
