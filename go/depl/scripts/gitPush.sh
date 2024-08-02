. sh/.env &&

branch="main"
if [[ ! -z "$defaultBranch" ]]; then
    branch=$defaultBranch
fi

git push origin $branch -f
