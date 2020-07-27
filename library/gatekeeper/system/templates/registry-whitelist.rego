package k8sallowedregistries

violation[{"msg": msg}] {
  c := containers[_]
  invalid_image(c.image)
  msg := sprintf("container <%v> has an invalid image <%v>, allowed registries are %v", [c.name, c.image, input.parameters.registries])
}

spec := input.review.object.spec

# Pods
containers[c] { c := spec.containers[_] }
containers[c] { c := spec.initContainers[_] }
containers[c] { c := spec.ephemeralContainers[_] }

# Deployments / StatefulSets
containers[c] { c := spec.template.spec.initContainers[_] }
containers[c] { c := spec.template.spec.containers[_] }
containers[c] { c := spec.template.spec.ephemeralContainers[_] }

invalid_image(img) {
  matches := [match | reg = input.parameters.registries[_] ; match = startswith(img, reg)]
  not any(matches)
}
