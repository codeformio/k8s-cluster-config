package k8svalidgitopssource

violation[{"msg": msg}] {
  branch := input.parameters.branchBlocklist[_]
  obj := input.review.object
  obj.spec.source.targetRevision == branch
  msg := sprintf("application <%v> has an invalid source branch <%v>", [obj.metadata.name, branch])
}

