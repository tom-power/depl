. sh/.env &&
ssh -t $remoteUser@$remoteHost "bash --login -c 'cd $deployDir && depl.sh $2'"