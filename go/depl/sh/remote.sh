. sh/.env &&
ssh -t $remoteUser@$remoteHost "bash --login -c 'cd $deployDir && sh/depl.sh $2'"