package k8svalidgitopssource

test_input_allowed {
    input := { "review": prod, "parameters": { "branchBlocklist": ["non-prod"] }}
    results := violation with input as input
    count(results) == 0
}

test_input_deny {
    input := { "review": nonProd, "parameters": { "branchBlocklist": ["non-prod"] }}
    results := violation with input as input
    count(results) == 1
}

nonProd = {
  "object": {
    "metadata": {
      "name": "abc"
    },
    "spec": {
      "source": {
        "targetRevision": "non-prod",
      },
    },
  }
}

prod = {
  "object": {
    "metadata": {
      "name": "xyz"
    },
    "spec": {
      "source": {
        "targetRevision": "prod",
      },
    },
  }
}
