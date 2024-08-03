branch="main"
if [ -f "sh/.env" ]; then
    . sh/.env
    if [ ! -z "$defaultBranch" ]; then
        branch=$defaultBranch
    fi
fi

git fetch --all && git reset --hard origin/$branch
