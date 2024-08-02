. sh/.env &&
ssh -t $remoteUser@$remoteHost \
"bash --login -c 'export defaultBranch=$defaultBranch; cd $deployDir && depl.sh $2'"
