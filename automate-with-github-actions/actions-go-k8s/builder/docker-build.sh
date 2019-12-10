while read p; do cd $GITHUB_WORKSPACE/$p && docker build -t r00tsh3ll/actionsgo-$p:$(echo $GITHUB_SHA | head -c7) . ; done < $GITHUB_WORKSPACE/microservices.txt
