tail docker/data/log | while read line; do
  if jq -e . >/dev/null 2>&1 <<<"$line"; then
    echo $line | jq;
  else
    echo $line;
  fi
done