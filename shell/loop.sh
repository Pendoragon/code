#!/bin/bash

function test {
  local count=10
  local fail=0

  for (( i = 0; i < $((count)); i++ ));do
    host "caicloud-prod-2016-cluster.caicloudapp.com"
    if [[ "$?" != "0" ]]; then
      return_val=1
      break
    fi
  done

  return $((fail))
}

test
