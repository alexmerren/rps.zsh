#!/usr/bin/env zsh

GITHUB_TOKEN=""

function rps() {
    # Create a temporary file that we can send the list of repositories to.
    QUEUE=(mktemp -ut rps_queue)
    mkfifo $QUEUE

    # Hash the github token to version the cache of list_repositories, and
    # insert the contents of cache into the queue.
    SHA=(echo $GITHUB_TOKEN | shasum -a 512 | awk '{print $1}')
    CACHE="rps_cache_$SHA.txt"
    cat $CACHE 2>/dev/null > $QUEUE &

    # Call list repositories and send the output to the cache and the queue.
    # De-duplicate the list coming from the queue and use fzf to select.
    list_repositories $GITHUB_TOKEN | tee $CACHE > $QUEUE &
    cat $QUEUE | unique_repositories | fzf | read -l repo

    git clone $repo

    rm $QUEUE
}
