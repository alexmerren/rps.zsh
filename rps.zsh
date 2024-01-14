#!/usr/bin/env zsh

GITHUB_TOKEN=""

function rps() {
    # Create a temporary file that we can send the list of repositories to.
    QUEUE=$(mktemp -ut rps_queue)
    mkfifo $QUEUE

    # Hash the github token to version the cache of list_repositories, and
    # insert the contents of cache into the queue.
    SHA=$(echo $GITHUB_TOKEN | shasum -a 512 | awk '{print $1}')
    CACHE="/tmp/rps_cache_$SHA.txt"
    cat $CACHE 2>/dev/null > $QUEUE &!

    # Call list repositories and send the output to the cache and the queue.
    # De-duplicate the list coming from the queue and use fzf to select.
    _rps_list $GITHUB_TOKEN | tee $CACHE > $QUEUE &!
    cat $QUEUE | _rps_unique | fzf | read repo

    git clone git@github.com:$repo.git

    rm $QUEUE
}

rps
