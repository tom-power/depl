cat docker/data/event | jq -R 'try fromjson catch .'
