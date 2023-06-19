branch="master"
if [[ ! -z "$defaultBranch" ]]; then
    branch=$defaultBranch
fi

git fetch --all && git reset --hard origin/$branch
