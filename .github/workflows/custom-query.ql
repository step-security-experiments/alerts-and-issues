import go

from CallExpr queryCall, FunctionCallArg userArg
where
  queryCall.getTarget().getName() = "Query" and
  queryCall.getArgument(0) = userArg and
  userArg instanceof StringLiteral and
  userArg.getValue().matches(".*OR.*")
select queryCall, "Potential SQL injection vulnerability detected: User input used directly in query."
