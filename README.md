# pluralsightdistributedgo
Code from the PluralSight course on distributed go: https://app.pluralsight.com/library/courses/go-build-distributed-applications/table-of-contents

## setup
For my mac, I did:

`brew update`

`brew install rabbitmq`

Register `/usr/local/Cellar/rabbitmq/3.8.1/sbin` in the `PATH`

### starting rabbitmq
`rabbitmq-server`

Check the status using: `rabbitmqctl list_queues`