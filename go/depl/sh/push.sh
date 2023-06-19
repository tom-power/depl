. sh/.env &&

branch="master"
if [[ ! -z "$defaultBranch" ]]; then
    branch=$defaultBranch
fi

git push origin $branch -f
